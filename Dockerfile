FROM golang:alpine AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod verify

COPY . .

RUN go build -o main cmd/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/main /build/main

EXPOSE 8080

CMD ./main
# TODO name image
# write docker run command with name and ports