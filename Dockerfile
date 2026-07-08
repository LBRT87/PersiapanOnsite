FROM golang:1.26-alpine

WORKDIR /app
COPY contracts/auth-service contracts/auth-service
COPY web-backend/course-service/go.mod ./
COPY . .
RUN go mod tidy

EXPOSE 8080