FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY *.mod *.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 9000

CMD ["./main"]