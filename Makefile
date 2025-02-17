.PHONY: run-go run-cpp run-python format

run-go:
	@go run go/main.go

run-cpp:
	@cmake -S . -B build
	@cmake --build build
	@./build/hello-world

run-python:
	@uv run python/main.py

format:
	@bun run fmt

lint:
	@bun run lint

run-typescript:
	@bun run typescript/main.ts

run-rust:
	@cargo run
