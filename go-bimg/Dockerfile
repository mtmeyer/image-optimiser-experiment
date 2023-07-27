FROM golang as builder

WORKDIR /app

COPY *.go .
COPY go.mod .

RUN apt update
RUN apt -y install libvips libvips-dev

RUN go get
RUN go build -o bin

FROM debian:bookworm-slim as server

WORKDIR /app

RUN apt update
RUN apt -y install libvips

COPY --from=builder /app/bin .

CMD ["./bin"]
