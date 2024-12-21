golangci-lint:
	docker run --rm -v ${PWD}:/app:ro -w /app golangci/golangci-lint:v1.62.2 golangci-lint run
