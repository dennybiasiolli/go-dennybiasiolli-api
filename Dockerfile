ARG GO_VERSION=1.17

##
## Build
##
FROM golang:${GO_VERSION}-alpine AS build

COPY . /app

WORKDIR /app

RUN go build -o app


##
## Deploy
##
FROM alpine:latest

ENV HTTP_LISTEN=":8080"

ENV PG_HOST=""
ENV PG_PORT="5432"
ENV PG_USER=""
ENV PG_PASSWORD=""
ENV PG_SSLMODE="disable"
ENV PG_DATABASE=""

ENV SEND_EMAIL_AFTER_CITAZIONE_ADDED="0"
ENV EMAIL_HOST=""
ENV EMAIL_PORT="587"
ENV EMAIL_HOST_USER=""
ENV EMAIL_HOST_PASSWORD=""
ENV EMAIL_DEFAULT_FROM=""
ENV EMAIL_DEFAULT_TO=""

RUN addgroup -S nonroot && adduser -S -G nonroot nonroot
USER nonroot:nonroot
WORKDIR /home/nonroot

# Whenever we call the LoadLocation() function in the time package it looks
# for the “Time Zone Database” (zoneinfo.zip ) which maintains a list of time
# zone information from locations around the world. Since the alpine image
# does not have this file we will have to inject it into the image manually
# RUN apk add --no-cache tzdata

# You may need to inject CA root certs into the Alpine image as well.
# You can do this by using apk.
# RUN apk --update add ca-certificates
# Or copying the ca-certs from the builder image (first stage)
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/


COPY --from=build /app/app app
COPY .env.default .env


EXPOSE 8080

ENTRYPOINT ["./app"]
