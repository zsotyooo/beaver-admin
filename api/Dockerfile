FROM golang:1.21.5-alpine AS base

RUN apk add --no-cache curl

FROM base AS dev

RUN go install github.com/cosmtrek/air@latest

WORKDIR /opt/api
CMD ["air"]
