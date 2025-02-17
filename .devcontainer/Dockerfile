FROM golang:1.23.6-bookworm

# Create a non-root user
RUN useradd -m -s /bin/bash appuser

SHELL ["/bin/bash", "-o", "pipefail", "-c"]

# Set up the environment including C++ development tools
# Note (stitle): Use "apt-cache policy [deps]" on unpinned deps to find names
RUN apt-get update && \
    apt-get install -y \
        cmake=3.25.1-1 \
        clang=1:14.0-55.7~deb12u1 \
        clangd=1:14.0-55.7~deb12u1 \
        curl=7.88.1-10+deb12u8 \
        unzip=6.0-28 --no-install-recommends && \
    rm -rf /var/lib/apt/lists/*

# Set up UV a modern package manager for Python
COPY --from=ghcr.io/astral-sh/uv:latest /uv /uvx /bin/

# Switch to the non-root user
USER appuser

# Set up TypeScript development with bun for appuser
RUN curl -fsSL https://bun.sh/install | bash

# Set up Rust
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y

# Add a health check for the container
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 CMD curl -f http://localhost/ || exit 1

ENTRYPOINT ["/bin/bash"]
