//@jeffotoni
use chrono::prelude::Utc;
use hyper::client::{Client, HttpConnector};
use hyper::http::Uri;
use hyper::service::{make_service_fn, service_fn};
use hyper::{Body, Response};
use std::convert::Infallible;
use std::sync::Arc;

async fn handler(
    client: Arc<Client<HttpConnector>>,
    uri: Uri,
) -> Result<Response<Body>, Infallible> {
    let resp = client.get(uri).await.unwrap();
    let len = resp
        .headers()
        .get(hyper::header::CONTENT_LENGTH)
        .unwrap()
        .to_str()
        .unwrap();
    let response = Response::builder()
        .status(hyper::StatusCode::OK)
        .header("Engine", "Rust/HyperV5")
        .header("Location", "/v1/user")
        .header("Date", Utc::now().to_rfc3339())
        .header(hyper::header::CONTENT_TYPE, "application/json")
        .header(hyper::header::CONTENT_LENGTH, len)
        .body(resp.into_body())
        .unwrap();
    Ok(response)
}

#[tokio::main(worker_threads = 12)]
async fn main() -> Result<(), Box<dyn std::error::Error + Send + Sync + 'static>> {
    let url = "http://localhost:3000/v1/avatar";
    let uri = Uri::from_static(url);
    let cli = Arc::new(hyper::client::Client::new());
    let make_svc = make_service_fn(move |_| {
        let cli = Arc::clone(&cli);
        let uri = uri.clone();
        async { Ok::<_, Infallible>(service_fn(move |_| handler(Arc::clone(&cli), uri.clone()))) }
    });
    let addr = "0.0.0.0:8080".parse().unwrap();
    let server = hyper::Server::bind(&addr).serve(make_svc);
    server.await?;
    Ok(())
}

