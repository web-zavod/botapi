fn main() -> Result<(), Box<dyn std::error::Error>> {
	// Build command service
	tonic_build::configure()
		.out_dir("src")
		.compile(
            &["command_service.proto"],
            &["../../proto/botapi/command/v1"],
			)
		.unwrap();

	// Build expence service
	tonic_build::configure()
		.out_dir("src")
		.compile(
            &["expense_service.proto"],
            &["../../proto/botapi/expense/v1"],
			)
		.unwrap();

	// Build healthcheck service
    tonic_build::configure()
        .out_dir("src")
        .compile(
            &["heathcheck_service.proto"],
            &["../../proto/botapi/healthcheck/v1"],
        )
        .unwrap();

	// Build media service
    tonic_build::configure()
        .out_dir("src")
        .compile(
            &["media_service.proto"],
            &["../../proto/botapi/media/v1"],
        )
        .unwrap();

    Ok(())
}
