FROM golang:1.23.6-bookworm

# Create a non-root user
RUN useradd -m -s /bin/bash appuser

SHELL ["/bin/bash", "-o", "pipefail", "-c"]

# Set up the environment including C++ development tools
RUN apt-get update && \
    apt-get install -y cmake clang clangd unzip --no-install-recommends && \
    rm -rf /var/lib/apt/lists/*

# Set up UV a modern package manager for Python
COPY --from=ghcr.io/astral-sh/uv:latest /uv /uvx /bin/

# Switch to the non-root user
USER appuser

# Set up TypeScript development with bun for appuser
RUN curl -fsSL https://bun.sh/install | bash

# Add a health check for the container
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 CMD curl -f http://localhost/ || exit 1

ENTRYPOINT ["/bin/bash"]
