FROM golang:1.23.6-bookworm

RUN apt-get update && apt-get install -y cmake clang clangd

COPY --from=ghcr.io/astral-sh/uv:latest /uv /uvx /bin/

ENTRYPOINT ["/bin/sh"]
