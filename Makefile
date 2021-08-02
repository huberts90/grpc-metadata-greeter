DOCKER_IMAGE := grpc-metadata-greeter
VERSION := 1.0.0


.PHONY: proto
proto: ## compile protobuf using local toolchain, "convencience" target not official
	scripts/compileproto.sh

.PHONY: docker-run-test
docker-run-test:
	docker build \
		-f Dockerfile \
		-t $(DOCKER_IMAGE):$(VERSION)-test \
		.