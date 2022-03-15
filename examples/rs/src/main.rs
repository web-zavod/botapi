use std::net::SocketAddr;

use tonic::{transport::Server, Request, Response, Status};

use botapi::healthcheck::healthcheck_service_server::{
    HealthcheckService, HealthcheckServiceServer,
};
use botapi::healthcheck::CheckResponse;

pub struct Service {}

#[tonic::async_trait]
impl HealthcheckService for Service {
    async fn check(&self, _request: Request<()>) -> Result<Response<CheckResponse>, Status> {
        log::debug!("processing check call");
        Ok(tonic::Response::new(CheckResponse {
            message: "ok".to_string(),
        }))
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    env_logger::init();
    let addr = "[::]:5000";
    log::debug!("gRPC server serving at: {}", addr);
    let svc = HealthcheckServiceServer::new(Service {});
    let socket: SocketAddr = addr.parse().unwrap();
    Server::builder().add_service(svc).serve(socket).await?;
    Ok(())
}
