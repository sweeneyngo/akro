FROM golang:1.22-bookworm as builder

WORKDIR /usr/src/app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /run-app .

FROM debian:bookworm
RUN apt-get update && apt-get install -y ca-certificates && apt-get clean

COPY --from=builder /run-app /usr/local/bin/
COPY --from=builder /usr/src/app/model.json /app/model.json

CMD ["/usr/local/bin/run-app", "-data", "/app/model.json"]
