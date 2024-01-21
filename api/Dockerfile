FROM golang:1.21.5-alpine AS base

RUN apk add --no-cache curl

FROM base AS dev

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/onsi/ginkgo/ginkgo@latest

WORKDIR /opt/api
CMD ["air"]
