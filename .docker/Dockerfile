FROM golang:1.24 AS builder

WORKDIR /api

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 80

RUN go build -o /go/bin/api-app ./cmd/main.go

FROM alpine:3.20

WORKDIR /api

COPY --from=builder /api/internal/migration /migration

COPY --from=builder /go/bin/api-app /go/bin/api-app

CMD ./go/bin/api-app