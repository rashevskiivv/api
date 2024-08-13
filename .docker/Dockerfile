FROM golang:1.22 AS builder

WORKDIR /tax_api

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 8080

RUN go build -o /go/bin/app ./cmd/main.go

FROM alpine:3.20

WORKDIR /tax-api

COPY --from=builder /tax-api/internal/migration /migration

COPY --from=builder /go/bin/app /go/bin/app

CMD ./go/bin/app