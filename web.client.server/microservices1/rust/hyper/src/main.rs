//@jeffotoni
use hyper::{
    service::{make_service_fn, service_fn},
    Body, Request, Response,
};
use reqwest::Client;
use std::convert::Infallible;
use std::sync::Arc;
use std::time::Duration;
use hyper::header::HeaderValue;
use std::env;

async fn get_client(client: Arc<Client>, _req: Request<Body>, m1path: String, m2url: String) -> Result<Response<Body>, Infallible> {
    if _req.uri().path() != m1path {
        let not_found = Response::builder()
            .status(hyper::StatusCode::NOT_FOUND)
            .body(Body::from("Not found"))
            .unwrap();
        return Ok(not_found);
    }

    let response = client.get(&m2url)
        .send()
        .await;

    let json = match response {
        Ok(res) => res.text().await.unwrap(),
        Err(_) => "Error".to_string(),
    };

    let mut response = Response::new(Body::from(json));
    *response.status_mut() = hyper::StatusCode::OK;

    response.headers_mut().insert(hyper::header::CONTENT_TYPE, hyper::header::HeaderValue::from_static("application/json"));
    response.headers_mut().insert("Engine", HeaderValue::from_static("Rust/Hype"));    
    response.headers_mut().insert(hyper::header::LOCATION, hyper::header::HeaderValue::from_static("/v1/user"));
    response.headers_mut().insert(hyper::header::DATE, hyper::header::HeaderValue::from_str(&chrono::Utc::now().to_rfc3339()).unwrap());

    Ok(response)
}

#[tokio::main(worker_threads = 32)]
async fn main() -> Result<(), Box<dyn std::error::Error + Send + Sync +'static>> {
    let client = Arc::new(
        reqwest::Client::builder()
            .tcp_keepalive(Duration::from_secs(30))
            .build()
            .unwrap()
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
                Ok(val) if !val.is_empty() => val,
                _ => String::from("3000"),
            };
        
            let m2url: String = format!("{}{}{}{}", domain,":",port, path);
            println!("{}",m2url);

           Ok::<_, Infallible>(service_fn(move |req| get_client(client.clone(), req, m1path.clone(), m2url.clone())))
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
    //let addr = SocketAddr::from(([127, 0, 0, 1], 8080));
    //let server = hyper::Server::bind(&addr).serve(make_svc);
    println!("Listening on http://{}", addr);

    server.await?;

    Ok(())
}
