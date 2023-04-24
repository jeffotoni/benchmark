//@jeffotoni
use hyper::{
    service::{make_service_fn, service_fn},
    Body, Request, Response,
};

use reqwest::Client as ReqwestClient;
use std::convert::Infallible;
use std::env;
use std::sync::Arc;
use std::time::Duration;

async fn get_client(
    client: Arc<ReqwestClient>,
    _req: Request<Body>,
    m1path: String,
    m2url: String,
) -> Result<Response<Body>, Infallible> {
    if _req.uri().path() != m1path {
        let not_found = Response::builder()
            .status(hyper::StatusCode::NOT_FOUND)
            .body(Body::from("Not found"))
            .unwrap();
        return Ok(not_found);
    }

    let response = client.get(&m2url).send().await.unwrap(); // Corrigido aqui
    if !response.status().is_success() {
        panic!("Erro na requisição: {}", response.status());
    }
    let bytes = response.bytes().await.unwrap();
    let body = Body::from(bytes);

    let response = Response::builder()
        .status(hyper::StatusCode::OK)
        .header(hyper::header::CONTENT_TYPE, "application/json")
        .body(body)
        .unwrap();

    Ok(response)
}

#[tokio::main(worker_threads = 32)]
async fn main() -> Result<(), Box<dyn std::error::Error + Send + Sync + 'static>> {
    let client = Arc::new(
        ReqwestClient::builder()
            .tcp_keepalive(Duration::from_secs(30))
            .build()
            .unwrap(),
    );

    let make_svc = make_service_fn(move |_conn| {
        let client = client.clone();
        async {
            let m1path = if let Ok(val) = env::var("M1_PATH") {
                val
            } else {
                String::from("/v1/user")
            };

            let domain = match env::var("M2_DOMAIN") {
                Ok(val) if !val.is_empty() => val,
                _ => String::from("http://127.0.0.1"),
            };

            let path = match env::var("M2_PATH") {
                Ok(val) if !val.is_empty() => val,
                _ => String::from("/v1/avatar"),
            };

            let port = match env::var("M2_PORT") {
                Ok(val) if !val.is_empty() => val.parse::<u16>().unwrap(),
                _ => 3000,
            };

            let m2url: String = format!("{}:{}{}", domain, port, path);

            Ok::<_, Infallible>(service_fn(move |req| {
                get_client(client.clone(), req, m1path.clone(), m2url.clone())
            }))
        }
    });

    let domain = match env::var("M1_DOMAIN") {
        Ok(val) if !val.is_empty() => val,
        _ => String::from("0.0.0.0"),
    };

    let port = if let Ok(val) = env::var("M1_PORT") {
        val
    } else {
        String::from("8080")
    };

    let addr = format!("{}:{}", domain, port).parse().unwrap();
    let server = hyper::Server::bind(&addr).serve(make_svc);
    println!("Listening on http://{}", addr);

    server.await?;

    Ok(())
}
