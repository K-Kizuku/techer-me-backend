ARG GO_VERSION=1.22
ARG ALPINE_VERSION=3.18

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as go-builder

WORKDIR /go/src/c3-tutorial

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o api cmd/main.go


FROM alpine:${ALPINE_VERSION}

WORKDIR /usr/src/c3-tutorial

COPY --from=go-builder /go/src/c3-tutorial/api api
COPY --from=go-builder /go/src/c3-tutorial/api .env
RUN chmod +x "/usr/src/c3-tutorial/api"

ENTRYPOINT ["/usr/src/c3-tutorial/api"]