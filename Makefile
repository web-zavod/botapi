go_dir=gen/go
rs_dir=gen/rs
py_dir=gen/py

version = 1.1.2

proto_files=\
	proto/botapi/command/v1/*.proto \
	proto/botapi/expense/v1/*.proto \
	proto/botapi/healthcheck/v1/*.proto \
	proto/botapi/media/v1/*.proto

lint:
	docker run --volume "$(shell pwd):/workspace" --workdir /workspace bufbuild/buf lint proto

compile: compile-py compile-go

compile-go:
	docker run --volume "$(shell pwd):/workspace" --workdir /workspace golang:1.17-stretch make compile-go-in-docker

compile-py:
	docker run --volume "$(shell pwd):/workspace" --workdir /workspace  python:3.10 make compile-py-in-docker

compile-go-in-docker:
	apt-get update && apt-get install -y unzip jq
	cd /tmp; curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.19.1/protoc-3.19.1-linux-x86_64.zip
	cd /tmp; unzip protoc-3.19.1-linux-x86_64.zip -d /usr

	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

	cd $(go_dir); rm -rf *
	protoc -Iproto \
		--go_out=$(go_dir) \
		--go_opt paths=source_relative \
		--go-grpc_out=$(go_dir) \
		--go-grpc_opt paths=source_relative \
		$(proto_files)
	cd $(go_dir); mv botapi/* .; rm -rf botapi

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

