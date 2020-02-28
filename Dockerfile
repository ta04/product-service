FROM golang:alpine AS build

RUN apk update && apk upgrade && apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o product-service

# Run container
FROM alpine:latest

RUN apk add --no-cache git

RUN mkdir /app
WORKDIR /app
COPY --from=build /app/product-service .

CMD ["./product-service"]