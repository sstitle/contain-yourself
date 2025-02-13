.PHONY: run-go build-cpp

run-go:
	@go run go/main.go

run-cpp:
	@cmake -S . -B build
	@cmake --build build
	@./build/hello-world

run-python:
	@uv run python/main.py
