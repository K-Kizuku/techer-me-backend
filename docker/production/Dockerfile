ARG GO_VERSION=1.22
ARG ALPINE_VERSION=3.18

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as go-builder

WORKDIR /go/src/techer-me-backend

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o api cmd/main.go


FROM alpine:${ALPINE_VERSION}

WORKDIR /usr/src/techer-me-backend

COPY --from=go-builder /go/src/techer-me-backend/api api
COPY --from=go-builder /go/src/techer-me-backend/api .env
RUN chmod +x "/usr/src/techer-me-backend/api"

ENTRYPOINT ["/usr/src/techer-me-backend/api"]