# tests section
.PHONY: test-unit
test-unit: ## Unit testing
	go test -v ./internal/...

.PHONY: test-integration
test-integration: docker-build docker-run
	@echo "==> Running integration tests"
	go test -v ./tests/integration/...
	make docker-stop


# linter section
.PHONY: lint
lint:
	golangci-lint run -v


# docker section
.PHONY: docker-build
docker-build:
	docker build -t flight-tracker .

.PHONY: docker-run
docker-run:
	docker run --name flight-tracker -d --rm -p 8080:8080 flight-tracker:latest

.PHONY: docker-stop
docker-stop:
	docker stop flight-tracker
