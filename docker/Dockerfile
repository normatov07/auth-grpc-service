#Building stage
FROM golang:1.19.0-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go get github.com/lib/pq
RUN go build -o main cmd/main.go 
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

#Running stages
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY pkg/db/migration ./migration

EXPOSE 8080
# CMD ["/app/main"]