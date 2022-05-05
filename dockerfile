FROM golang:1.18-alpine AS build_base

RUN apk --update add git

# Set the Current Working Directory inside the container
WORKDIR /app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out .

# Start fresh from a smaller image
FROM alpine:3.15.0
RUN apk add ca-certificates

COPY --from=build_base /app/out /app

# This container exposes port 8080 to the outside world
EXPOSE 4000

# Run the binary program produced by `go install`
CMD ["/app"]