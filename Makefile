rs_dir=gen/rs
py_dir=gen/py

version = 0.1.0

lint:
	docker run --volume "$(shell pwd):/workspace" --workdir /workspace bufbuild/buf lint proto

compile: compile-rs

compile-rs:
	docker run --volume "$(shell pwd):/workspace" --workdir /workspace rust:1.59-buster make compile-rs-in-docker

compile-rs-in-docker:
	cd $(rs_dir)
	rustup component add rustfmt
	cargo build
	sed -i "3s/.*version.*/version = \"$(version)\"/" $(rs_dir)/Cargo.toml
	cd ../..
	sed -i "3s/.*version.*/version = \"$(version)\"/" $(rs_dir)/Cargo.toml


