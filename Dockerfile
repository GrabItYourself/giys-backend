FROM golang:1.19 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./apigateway ./apigateway
COPY ./auth ./auth
COPY ./lib ./lib
COPY ./notification ./notification
COPY ./order ./order
COPY ./payment ./payment
COPY ./shop ./shop
COPY ./user ./user

ENV GOOS=linux
ENV CGO_ENABLED=0

ARG SERVICE

RUN go build -o main ./${SERVICE}/cmd/main.go

FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates tzdata

ENV TZ=Asia/Bangkok

ARG SERVICE

COPY --from=builder /app/main .

CMD ["/app/main"]
