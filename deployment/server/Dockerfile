# Build stage
FROM golang:alpine AS builder
WORKDIR /openiiot
COPY biz biz
COPY *.go .
COPY go.mod go.mod
COPY go.sum go.sum
RUN go build -o server .

# Final stage
FROM alpine
WORKDIR /openiiot
COPY --from=builder /openiiot/server .
COPY biz/config/conf/*.yml /openiiot/conf/
# COPY public/* /openiiot/public/

EXPOSE 8085 

CMD ["./server", "-conf-dir=conf/config.yml"]
