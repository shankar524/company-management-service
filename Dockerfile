## Build
FROM golang:1.19 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

## Deploy
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY ./app/migrations .
COPY .env .env

COPY --from=build /app/main .

EXPOSE 8080

ENTRYPOINT ["./main"]
