FROM golang:1.19-buster AS buildenv

LABEL maintainer="Chidozie Okafor <Chidosiky2015@gmail.com>"

# ARG BIN
# ARG VERSION
# ARG GIT_COMMIT

# ENV BIN=${BIN}
# ENV VERSION=${VERSION}
# ENV GIT_COMMIT=${GIT_COMMIT}

ENV GO111MODULE=on
ENV CGO_ENABLED=0

# SHELL ["/bin/bash", "-c"]

# Create a location in the container for the source code.
RUN mkdir -p /app
WORKDIR /app

# Copy the module files first and then download the dependencies. If this
# doesn't change, we won't need to do this again in future builds.
COPY go.* /app/
# COPY go.sum .
# COPY go.mod .


# COPY . .

RUN echo "Downloading dependencies..."
RUN pwd
RUN go mod download
RUN go mod verify

# Copy the source code into the container.
COPY pkg pkg
COPY cmd/auth cmd/auth
# COPY . .

# RUN go mod download
# RUN go mod verify

RUN go build \
    -mod=readonly \
    -tags=persistence_mysql \
    -ldflags "-X 'placio-pkg/buildinfo.BuildTime=$(date -u '+%Y-%m-%d %H:%M:%S')'" \
    -a -o /go/bin/app ./cmd/auth

FROM scratch
COPY --from=buildenv /go/bin/app /go/bin/app
COPY --from=buildenv /etc/ssl/certs /etc/ssl/certs
ENTRYPOINT ["/go/bin/app"]
