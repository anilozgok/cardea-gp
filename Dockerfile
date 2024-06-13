FROM golang:1.21 AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o cardea /app/main.go

FROM alpine

WORKDIR /app
COPY --from=builder /app/cardea /app/cardea
COPY --from=builder /app/configs /app/configs
COPY --from=builder /app/exercise.csv /app/exercise.csv
COPY --from=builder /app/diet.csv /app/diet.csv
COPY --from=builder /app/recipes_raw_nosource_epi.json /app/recipes_raw_nosource_epi.json
EXPOSE 8080

CMD ["/app/cardea"]