default: help

.PHONY: run
run: ## running a server through linux (default: localhost:8080)
	go run ./cmd/server/main.go

.PHONY: test
test: ## running a test through linux
	go test -v -coverpkg=./internal/...,./cmd/server/... ./...

.PHONY: test_coverage
test_coverage: ## running a test coverage via linux
	go test -coverprofile=coverage.out -coverpkg=./internal/...,./cmd/server/... ./...
	go tool cover -html=coverage.out

.PHONY: lint
lint: ## run golangci-lint (2-minute wait)
	docker run -t --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.60 golangci-lint run -v

.PHONY: run_pprof
run_pprof: ## run after 'make pprof_heap'
	go run ./cmd/pprof/main.go

.PHONY: pprof_heap
pprof_heap: ## run before 'make run_pprof'
	curl -o heap.out http://localhost:6060/debug/pprof/heap
	go tool pprof heap.out

.PHONY: mockgen
mockgen: ## install mockgen 'go install go.uber.org/mock/mockgen@latest'
	go generate ./...

.PHONY: build
build: ## build app for linux
	CGO_ENABLED=0 GOOS=linux go build -o jwt_oauth_sso ./cmd/server/main.go

.PHONY: docker_build
docker_build: ## build app for docker
	docker build . -t ghcr.io/fromsi/jwt_oauth_sso:latest

.PHONY: docker_run_img
docker_run_img: ## run server through docker (port 8080)
	docker run --rm -p "8080:8080" ghcr.io/fromsi/jwt_oauth_sso:latest

.PHONY: docker_push
docker_push: ## open link https://github.com/settings/tokens > Generate New Token > Classic > write:packages
	docker push ghcr.io/fromsi/jwt_oauth_sso:latest

.PHONY: help
help: ## display this help message
	@cat $(MAKEFILE_LIST) | grep -e "^[a-zA-Z_\-]*: *.*## *" | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
