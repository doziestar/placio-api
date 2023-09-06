ARG GOLANG_VERSION=1.21
# Stage 1: Build the application
FROM golang:${GOLANG_VERSION}-alpine3.17

WORKDIR /work

COPY go.mod .
COPY go.sum .
COPY ./pkg/go.mod ./pkg/
COPY ./pkg/go.sum ./pkg/
COPY ./realtime_server/go.mod ./realtime_server/
COPY ./realtime_server/go.sum ./realtime_server/


# Download dependencies
RUN go mod download
RUN go install github.com/cespare/reflex@latest
RUN #go install github.com/cosmtrek/air@latest

# Copy the rest of the application source code
COPY . .

ENV PORT=7071

EXPOSE 7071

# Run the binary
CMD ["reflex", "-r", "\\.go$", "-s", "--", "go", "run", "./realtime_server/main.go"]
