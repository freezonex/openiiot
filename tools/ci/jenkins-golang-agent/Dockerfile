# Use the official Golang image as a parent image
FROM golang:latest

# Update and install necessary packages
RUN apt-get update && apt-get install -y \
    protobuf-compiler \
    git \
    openjdk-11-jre-headless

# Install Hertz
RUN go install github.com/cloudwego/hertz/cmd/hz@latest

# Set global Git configuration (optional and not recommended for credentials)
# It's better to configure credentials at runtime in Jenkins pipeline
# RUN git config --global user.email "your-email@example.com"
# RUN git config --global user.name "Your Name"

# Other necessary setup can be added here

CMD ["bash"]
