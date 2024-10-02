.PHONY: run
run:
	go run ./cmd/server/main.go

# make pprof_heap
.PHONY: run_pprof
run_pprof:
	go run ./cmd/pprof/main.go

# make run_pprof
.PHONY: pprof_heap
pprof_heap:
	curl -o heap.out http://localhost:6060/debug/pprof/heap
	go tool pprof heap.out

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -o jwt_oauth_sso ./cmd/server/main.go

.PHONY: docker_build
docker_build:
	docker build . -t ghcr.io/fromsi/jwt_oauth_sso:latest

.PHONY: docker_run_img
docker_run_img:
	docker run --rm -p "8080:8080" ghcr.io/fromsi/jwt_oauth_sso:latest

# https://github.com/settings/tokens
# Generate New Token > Classic > write:packages
.PHONY: docker_push
docker_push:
	docker push ghcr.io/fromsi/jwt_oauth_sso:latest
