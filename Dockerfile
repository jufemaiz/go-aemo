FROM golang:1.17-alpine as builder

ARG SSH_KEY

RUN apk add bash build-base curl gcc git inotify-tools make openssh postgresql-client protobuf-dev tzdata zip

# Needed for private GIT repository access
RUN mkdir -p /root/.ssh && \
    echo "$SSH_KEY" > /root/.ssh/id_rsa && \
    chmod 600 /root/.ssh/id_rsa && \
    ssh-keyscan github.com > /root/.ssh/known_hosts

# Check that we can connect to github
RUN ssh -T git@github.com; LOGIN=$?; \
    if [ "$LOGIN" = "255" ] ; \
    then echo "----" && echo "SSH KEY ERROR (CODE: $LOGIN): YOU NEED TO USE BUILD ARGS! README.md FOR MORE INFO" && echo "----" && exit 1; \
    else exit 0; \
    fi

# Set GOPRIVATE
ENV GOPRIVATE github.com/enosi/*

RUN touch ~/.gitconfig && \
    (echo "[url \"ssh://git@github.com/\"]" > ~/.gitconfig) && \
    (echo "  insteadOf = https://github.com/" >> ~/.gitconfig)

# We need to copy the entire directory to allow `go mod tidy` to work properly
COPY . /app

WORKDIR /app

RUN make dependencies && make dependencies-download
RUN make tools

FROM golang:1.17-alpine

LABEL maintainer="Joel Courtney <euphemize@gmail.com>"
LABEL application="go-aemo"

ARG GOLANGCI_VERSION="v1.39.0"
ENV GOLANGCI_VERSION $GOLANGCI_VERSION

RUN apk add bash build-base curl gcc git inotify-tools make openssh postgresql-client protobuf-dev tzdata zip

# Set GOPRIVATE
ENV GOPRIVATE github.com/enosi/*

RUN touch ~/.gitconfig && \
    (echo "[url \"ssh://git@github.com/\"]" > ~/.gitconfig) && \
    (echo "  insteadOf = https://github.com/" >> ~/.gitconfig)

COPY --from=builder /go /go

# Get the golangci-lint binary
RUN echo "Installing golangci-lint $GOLANGCI_VERSION" && \
    wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s ${GOLANGCI_VERSION}

COPY . /app

WORKDIR /app
