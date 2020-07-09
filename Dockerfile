# Dear Programmers,

# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
# *                                                 *
# *	This file belongs to Kevin Veros Hamonangan   *
# *	and	Fandi Fladimir Dachi and is a part of     *
# *	our	last project as the student of Del        *
# *	Institute of Technology, Sitoluama.           *
# *	Please contact us via Instagram:              *
# *	sleepingnext and fandi_dachi                  *
# *	before copying this file.                     *
# *	Thank you, buddy. 😊                          *
# *                                                 *
# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# Build the user-service
FROM golang:alpine AS build

RUN apk update && apk upgrade && apk add --no-cache git

RUN mkdir /go/src/app
WORKDIR /go/src/app

ENV GO111MODULE=on

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o product-service

# Copy the newly built product-service to Alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=build /go/src/app/product-service .