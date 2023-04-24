use hyper::client::Client;
use hyper::client::HttpConnector;
use hyper::http::Uri;

#[macro_use]
extern crate rocket;

struct AppState {
    client: Client<HttpConnector>,
    uri: Uri,
}

#[get("/v1/user")]
async fn index(ctx: &rocket::State<AppState>) -> Vec<u8> {
    let resp = ctx
        .client
        .get(ctx.uri.clone())
        .await
        .expect("request failed");
    hyper::body::to_bytes(resp.into_body())
        .await
        .expect("body read failed")
        .to_vec()
}

#[launch]
fn rocket() -> _ {
    let url = "http://localhost:3000/v1/avatar";
    let client = hyper::Client::new();
    let uri = Uri::from_static(url);
    let state = AppState { client, uri };
    let mut config = rocket::Config::default();
    config.address = "127.0.0.1".parse().unwrap();
    config.port = 8080;
    //config.workers = 64;    
    rocket::custom(config)
        .manage(state)
        .mount("/", routes![index])
}
