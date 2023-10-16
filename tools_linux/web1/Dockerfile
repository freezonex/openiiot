# Use the official Nginx image as the base image
FROM nginx:latest

# Copy the index.html file from the host to the container
COPY ./index.html /usr/share/nginx/html/index.html

# Copy the images directory from the host to the container
COPY ./images /usr/share/nginx/html/images

# Expose port 80 to allow incoming traffic
EXPOSE 80

# The CMD instruction specifies the command to run when the container starts
CMD ["nginx", "-g", "daemon off;"]
