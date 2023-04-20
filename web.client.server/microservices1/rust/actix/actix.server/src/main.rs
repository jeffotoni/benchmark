// @jeffotoni
use actix_web::{get, web, App, HttpResponse, HttpServer, Responder};
use reqwest::Client;
use serde_json::Value;
use chrono::Utc;

#[get("/v1/client/get")]
async fn get_client(client: web::Data<Client>) -> impl Responder {
    let response = client.get("http://localhost:3000/v1/customer/get")
        .send()
        .await;

    match response {
        Ok(res) => {
            if let Ok(json) = res.json::<Value>().await {
                HttpResponse::Ok()
                    .append_header(("Content-Type", "application/json"))
                    .append_header(("Engine", "Rust"))
                    .append_header(("Location", "/v1/client/post"))
                    .append_header(("Date", Utc::now().format("%Y-%m-%dT%H:%M:%S.%fZ").to_string()))
                    .json(json)
            } else {
                HttpResponse::InternalServerError().finish()
            }
        },
        Err(_) => HttpResponse::InternalServerError().finish(),
    }
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    let client = Client::new();

    HttpServer::new(move || {
        App::new()
            .app_data(web::Data::new(client.clone()))
            .service(get_client)
    })
    .bind("127.0.0.1:8080")?
    .run()
    .await
}

