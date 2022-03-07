rs_dir=gen/rs
py_dir=gen/py

version = 1.1.0

proto_files=\
	proto/botapi/command/v1/*.proto \
	proto/botapi/expense/v1/*.proto \
	proto/botapi/healthcheck/v1/*.proto \
	proto/botapi/media/v1/*.proto

lint:
	docker run --volume "$(shell pwd):/workspace" --workdir /workspace bufbuild/buf lint proto

compile: compile-rs compile-py

compile-rs:
	docker run --volume "$(shell pwd):/workspace" --workdir /workspace rust:1.59-buster make compile-rs-in-docker

compile-py:
	docker run --volume "$(shell pwd):/workspace" --workdir /workspace  python:3.10 make compile-py-in-docker

compile-rs-in-docker:
	find $(rs_dir)/src -type f -not -name 'lib.rs' -delete
	cd $(rs_dir)
	rustup component add rustfmt
	cargo build --lib
	sed -i "3s/.*version.*/version = \"$(version)\"/" $(rs_dir)/Cargo.toml

compile-py-in-docker:
	rm -rf $(py_dir)/*
	python -m pip install --no-cache-dir grpcio-tools==1.38.1
	python \
		-m grpc_tools.protoc \
		--python_out=$(py_dir) \
		--grpc_python_out=$(py_dir) \
		-I proto \
		$(proto_files)
	sed -i "1s/.*/__version__ = '$(version)'/" setup.py

