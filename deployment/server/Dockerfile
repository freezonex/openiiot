# Build stage
FROM golang:1.22.6 AS builder
WORKDIR /openiiot
COPY biz biz
COPY *.go ./
COPY go.mod go.mod
COPY go.sum go.sum
RUN go build -o server .

# Final stage
FROM golang:1.22.6
WORKDIR /openiiot
COPY --from=builder /openiiot/server .
COPY biz/config/conf/*.yml /openiiot/conf/
# COPY public/* /openiiot/public/

# Create a directory for volume mount
RUN mkdir -p /volumes

EXPOSE 8085 

CMD ["./server", "-conf-dir=conf"]
