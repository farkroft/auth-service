FROM golang:1.12 as builder
LABEL maintainer="Fajar AR <github.com/farkroft>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY application.yaml /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o auth-service

FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
COPY --from=builder /app /app
EXPOSE 8080
ENTRYPOINT ["/app/auth-service"]