extern crate log;
extern crate env_logger;
use botapi;
use botapi::healthcheck::CheckResponse;
use botapi::healthcheck::healthcheck_service_server::{HealthcheckServiceServer, HealthcheckService};
use tonic::{transport::Server, Request, Response, Status};
use std::net::SocketAddr;

pub struct Service{}

#[tonic::async_trait]
impl HealthcheckService for Service {
    async fn check(&self, _request: Request<()>) -> Result<Response<CheckResponse>, Status> {
		log::info!("call");
		Ok(tonic::Response::new(CheckResponse{message: "message 123".to_string()}))
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
	env_logger::init();
	let addr = "[::]:5000";
    log::debug!("gRPC server serving at: {}", addr);
    let svc = HealthcheckServiceServer::new(Service{});
    let socket: SocketAddr = addr.parse().unwrap();
    Server::builder()
        .add_service(svc)
        .serve(socket)
        .await?;
    Ok(())

}
