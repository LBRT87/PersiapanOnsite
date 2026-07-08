FROM golang:1.26-alpine

WORKDIR /app
COPY ./services/auth-service ./services/auth-service
RUN go mod tidy

EXPOSE 8080

