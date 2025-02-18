.PHONY: run-go run-cpp run-python format lint run-typescript run-rust generate-proto


######### Run commands
run-go:
	@go run go/main.go

run-cpp:
	@cmake -S . -B build
	@cmake --build build
	@./build/hello-world

run-python:
	@uv run python/main.py

run-rust:
	@cargo run

run-typescript:
	@bun run typescript/main.ts


### Lifecycle commands
format:
	@bun run fmt

lint:
	@bun run lint

generate-proto:
	@bun run buf generate
