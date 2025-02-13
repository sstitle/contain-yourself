FROM golang:1.23.6-bookworm

# Set up the environment including C++ development tools
RUN apt-get update && apt-get install -y cmake clang clangd unzip

# Set up UV a modern package manager for Python
COPY --from=ghcr.io/astral-sh/uv:latest /uv /uvx /bin/

# Set up TypeScript development with bun
RUN curl -fsSL https://bun.sh/install | bash

ENTRYPOINT ["/bin/sh"]
