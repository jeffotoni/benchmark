//@jeffotoni
use hyper::{
    service::{make_service_fn, service_fn},
    Body, Request, Response,
};
use reqwest::Client;
use std::convert::Infallible;
use std::net::SocketAddr;
use std::sync::Arc;
use std::time::Duration;
use hyper::header::HeaderValue;

async fn get_client(client: Arc<Client>, _req: Request<Body>) -> Result<Response<Body>, Infallible> {
    if _req.uri().path() != "/v1/client/get" {
        let not_found = Response::builder()
            .status(hyper::StatusCode::NOT_FOUND)
            .body(Body::from("Not found"))
            .unwrap();

        return Ok(not_found);
    }
    let response = client.get("http://localhost:3000/v1/customer/get")
        .send()
        .await;

    let json = match response {
        Ok(res) => res.text().await.unwrap(),
        Err(_) => "Error".to_string(),
    };

    let mut response = Response::new(Body::from(json));
    *response.status_mut() = hyper::StatusCode::OK;

    response.headers_mut().insert(hyper::header::CONTENT_TYPE, hyper::header::HeaderValue::from_static("application/json"));
    response.headers_mut().insert("Engine", HeaderValue::from_static("Rust"));    
    response.headers_mut().insert(hyper::header::LOCATION, hyper::header::HeaderValue::from_static("/v1/client/post"));
    response.headers_mut().insert(hyper::header::DATE, hyper::header::HeaderValue::from_str(&chrono::Utc::now().to_rfc3339()).unwrap());

    Ok(response)
}

#[tokio::main(worker_threads = 32)]
async fn main() -> Result<(), Box<dyn std::error::Error + Send + Sync>> {
    let client = Arc::new(
        reqwest::Client::builder()
            .tcp_keepalive(Duration::from_secs(30))
            .build()
            .unwrap()
    );
    let make_svc = make_service_fn(move |_conn| {
        let client = client.clone();
        async { Ok::<_, Infallible>(service_fn(move |req| get_client(client.clone(), req))) }
    });

    let addr = SocketAddr::from(([127, 0, 0, 1], 8080));
    let server = hyper::Server::bind(&addr).serve(make_svc);

    println!("Listening on http://{}", addr);

    server.await?;

    Ok(())
}
