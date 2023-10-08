# Use the official Node-RED image as the base image
FROM nodered/node-red

# Copy your custom logo to the appropriate directory in the container
COPY logo_title.png /usr/src/node-red/node_modules/node-red/static

# Set any additional configurations if needed
# ...

# Continue with the CMD instruction from the base image
CMD ["npm", "start", "--", "--userDir", "/data"]
