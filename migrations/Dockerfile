FROM golang:1.19-alpine

WORKDIR /app

RUN apk --no-cache add curl

ARG VERSION=4.15.2
# PLATFORMS: linux-amd64,linux-arm64
ARG PLATFORM=linux-amd64

ENV VERSION ${VERSION}
ENV PLATFORM ${PLATFORM}

RUN curl -L "https://github.com/golang-migrate/migrate/releases/download/v${VERSION}/migrate.${PLATFORM}.tar.gz" | tar -xvz
