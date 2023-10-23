
ARG VARIANT=1-bullseye
FROM mcr.microsoft.com/devcontainers/go:0-${VARIANT}

ARG NODE_VERSION="lts/*"
RUN if [ "${NODE_VERSION}" != "none" ]; then su vscode -c ". /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi
WORKDIR /workspaces/products-go

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./

EXPOSE 9000