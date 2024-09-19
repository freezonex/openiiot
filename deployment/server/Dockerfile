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

# Install Docker CLI
RUN apt-get update && apt-get install -y \
    ca-certificates \
    curl \
    gnupg \
    lsb-release \
    software-properties-common && \
    curl -fsSL https://download.docker.com/linux/debian/gpg | apt-key add - && \
    echo "deb [arch=amd64] https://download.docker.com/linux/debian buster stable" > /etc/apt/sources.list.d/docker.list && \
    apt-get update && apt-get install -y docker-ce-cli && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /openiiot/server .
COPY biz/config/conf/*.yml /openiiot/conf/
# COPY public/* /openiiot/public/

# Create a directory for volume mount
RUN mkdir -p /volumes

EXPOSE 8085 

CMD ["./server", "-conf-dir=conf"]
