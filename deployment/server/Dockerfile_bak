# Start from the official Golang image to build your application
FROM golang:latest as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o ${RUN_NAME} .

# Start a new stage from scratch for a smaller final image
FROM alpine:latest

# Add bash to the image for your scripts
RUN apk add --no-cache bash

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/${RUN_NAME} /app/

# Copy necessary files and directories
COPY biz/config/conf/*.yml /app/output/conf/

# Copy additional configuration files
COPY conf/ /app/output/conf/

# Set the Current Working Directory inside the container
WORKDIR /app/output

# Command to run the executable
CMD ["./${RUN_NAME}"]
