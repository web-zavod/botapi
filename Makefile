lint:
	docker run --volume "$(shell pwd):/workspace" --workdir /workspace bufbuild/buf lint proto
