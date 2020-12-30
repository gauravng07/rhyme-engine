SERVICE ?= scb-rhyme-engine

# Enable this in PRD, Currently doing only locally.
#DOCKER_REGISTRY?=todo
#DOCKER_REPOSITORY_NAMESPACE?=scb-repo
#DOCKER_ID?=
#DOCKER_REPOSITORY_IMAGE=$(SERVICE)
#DOCKER_REPOSITORY=$(DOCKER_REGISTRY)/$(DOCKER_REPOSITORY_NAMESPACE)/$(DOCKER_REPOSITORY_IMAGE)

clean-test-cache:
	@ go clean -testcache

unit-test: clean-test-cache
	@ echo "Running Tests" && go test -v ./internal/...; ES=$$?; if [ $$ES -eq 1 ]; then echo "Test Failure..."; fi; exit $$ES

lint:
	@ golangci-lint run ./...

coverage: clean-test-cache
	@ go test -covermode=set -coverprofile=coverage.out ./...
	@ go tool cover -html=coverage.out
	@ goverreport -coverprofile=coverage.out; ES=$$?; if [ $$ES -eq 1 ]; then echo "Coverage failed."; fi; exit $$ES
	@ rm coverage.out

ci-docker-auth:
	@echo "Logging in to $(DOCKER_REGISTRY) as $(DOCKER_ID)"
	@docker login -u $(DOCKER_ID) -p $(DOCKER_PASSWORD) $(DOCKER_REGISTRY)

ci-docker-build:
	docker build -t $(DOCKER_REPOSITORY):$(GIT_HASH) . --build-arg SERVICE=$(SERVICE)
	docker tag $(DOCKER_REPOSITORY):$(GIT_HASH) $(DOCKER_REPOSITORY):latest

ci-docker-push: ci-docker-auth
	docker push $(DOCKER_REPOSITORY)

docker-build:
	docker build -t scb-rhyme-engine:local . --build-arg SERVICE=$(SERVICE)