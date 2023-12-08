.DEFAULT_GOAL=help

build: ## downloads dependencies and create Go binary
	go mod download
	go build -o cmd/server api/*.go

run: build ## runs Go binary
	./cmd/server

grpc-client: ## runs gRPC "evans" client. Install client if not present: https://github.com/ktr0731/evans
# Sample commands:
# - show package
# - package greet.v1
# - show service
# - service GreetService
# - show rpc
# - call Greet
# To exit from stream or quit the client: ctrl+D
	evans --host localhost --port 5051 --reflection repl

help:
	@grep -E '^[a-zA-Z_-]+:.*?\#\# .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?\#\# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
