FROM golang:1.17-alpine

LABEL maintainer="Joel Courtney <euphemize@gmail.com>"
LABEL application="go-aemo"

ARG GOLANGCI_VERSION="v1.62.0"
ENV GOLANGCI_VERSION $GOLANGCI_VERSION

RUN apk add bash build-base curl gcc git inotify-tools make openssh postgresql-client protobuf-dev tzdata zip

# We need to copy the entire directory to allow `go mod tidy` to work properly
COPY . /app

WORKDIR /app

RUN make dependencies && make dependencies-download
RUN make tools

# Get the golangci-lint binary
RUN echo "Installing golangci-lint $GOLANGCI_VERSION" && \
    wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s ${GOLANGCI_VERSION}

COPY . /app

WORKDIR /app
