fn main() -> Result<(), Box<dyn std::error::Error>> {
	// Build healthcheck service
    tonic_build::configure()
        .out_dir("src")
        .compile(
            &["../../proto/botapi/healthcheck/v1/heathcheck_service.proto"],
            &["../../proto/botapi/healthcheck/v1"],
        )
        .unwrap();
    Ok(())
}
