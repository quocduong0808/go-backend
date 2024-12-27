FROM golang:alpine AS builder

# ENV GOOS linux
# ENV CGO_ENABLED 0

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o go-app ./cmd/server/main.go

FROM alpine AS production

COPY config/ ./config/

COPY --from=builder /build/go-app /

RUN chmod +x /go-app

ENTRYPOINT [ "/go-app" ]

