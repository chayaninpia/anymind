FROM golang:1.18 as builder
WORKDIR /app
COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum
RUN go mod download
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/app ./cmd/anymind/main.go

FROM alpine
RUN apk add tzdata && \
    cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime && \
    echo "Asia/Bangkok" >  /etc/timezone && \
    apk del tzdata
WORKDIR /root/
COPY --from=builder /app/bin .
ENV GIN_MODE release
EXPOSE 9997
CMD ./app