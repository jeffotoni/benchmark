use hyper::client::Client;
use hyper::client::HttpConnector;
use rocket::State;
use rocket::Config;

#[macro_use]
extern crate rocket;

#[get("/v1/user")]
async fn index(client: &State<Client<HttpConnector>>) -> String {
    // change here to the microservice 2 url.
    let url = "http://127.0.0.1:3000/v1/avatar";
    let req = hyper::Request::builder()
        .method(hyper::Method::GET)
        .uri(url)
        .body(hyper::Body::from(""))
        .expect("request builder build failed");
    let resp = client.request(req).await.expect("request failed");
    let body_bytes = hyper::body::to_bytes(resp.into_body())
        .await
        .expect("body failed");
    return String::from_utf8(body_bytes.to_vec()).unwrap();
}

#[launch]
fn rocket() -> _ {
    let client = hyper::Client::new();
    let config = Config {
        address: "127.0.0.1".parse().unwrap(),
        port: 8080, // Change the port number here
        ..Config::default()
    };
    rocket::custom(config).manage(client).mount("/", routes![index])
    //rocket::build().manage(client).mount("/", routes![index])
}
