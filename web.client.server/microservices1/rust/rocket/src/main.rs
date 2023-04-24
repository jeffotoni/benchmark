//@jeffotoni
use hyper::client::Client;
use hyper::client::HttpConnector;
use hyper::http::Uri;
use rocket::Config;
 
#[macro_use]
extern crate rocket;

struct AppState {
    client: Client<HttpConnector>,
    uri: Uri,
}

#[get("/v1/user")]
async fn index(ctx: &rocket::State<AppState>) -> String {
    let resp = ctx
        .client
        .get(ctx.uri.clone())
        .await
        .expect("request failed");
    let body_bytes = hyper::body::to_bytes(resp.into_body())
        .await
        .expect("body read failed");
    return String::from_utf8(body_bytes.to_vec()).unwrap();
}

#[launch]
fn rocket() -> _ {
    let url = "http://localhost:3000/v1/avatar";
    let client = hyper::Client::new();
    let uri = Uri::from_static(url);
    let state = AppState { client, uri };
    
    let config = Config {
        address: "127.0.0.1".parse().unwrap(),
        port: 8080, // Change the port number here
        ..Config::default()
    };
    rocket::custom(config)
        .manage(state)
        .mount("/", routes![index])
    //rocket::build().manage(state).mount("/", routes![index])
}
