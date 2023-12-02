FROM golang:1.20 AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o cardea /app/main.go

FROM alpine

WORKDIR /app
COPY --from=builder /app/wincoi /app/wincoi
COPY --from=builder /app/configs /app/configs

EXPOSE 8080

CMD ["/app/wincoi"]