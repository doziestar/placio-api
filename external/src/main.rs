mod handlers;
mod models;
mod routes;

use axum::{http::{
    header::{ACCEPT, AUTHORIZATION, CONTENT_TYPE},
    HeaderValue, Method,
}, routing::get, Router};

#[tokio::main]
async fn main() {
    let app = Router::new().route("/", get(|| async { "Hello, World!" }));

    let listener = tokio::net::TcpListener::bind("0.0.0.0:7075").await.unwrap();
    print!("Listening on port 7075");
    axum::serve(listener, app).await.unwrap();
}
