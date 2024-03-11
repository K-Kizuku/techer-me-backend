ARG GO_VERSION=1.22
ARG ALPINE_VERSION=3.18

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION}

WORKDIR /go/src/c3-tutorial

RUN apk update && apk add git
RUN go install github.com/cosmtrek/air@latest

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

CMD ["air", "-c", ".air.toml"]