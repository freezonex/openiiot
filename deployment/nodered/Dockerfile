FROM nodered/node-red:4.0.2

# Copy custom Node-RED nodes package.json
COPY deployment/nodered/data/package.json /data/

# Copy package.json and the plugin installation script to /opt/nodered/
COPY deployment/nodered/data/package.json /opt/nodered/
COPY deployment/nodered/install-plugins.sh /opt/nodered/

# Change file permissions for the install script
USER root
RUN chmod +x /opt/nodered/install-plugins.sh
USER node-red

# Install custom Node-RED nodes
RUN cd /data && npm install

# Define arguments with default values for source and target directories
#ARG SRC_DIR=node-red/packages/node_modules/@node-red/editor-client/src/images
ARG SRC_DIR=deployment/nodered/images
ARG DEST_DIR=/usr/src/node-red/node_modules/@node-red/editor-client/public/red/images

# Update all logo and icons
COPY ${SRC_DIR}/node-red.svg ${DEST_DIR}/node-red.svg
COPY ${SRC_DIR}/node-red-256.svg ${DEST_DIR}/node-red-256.svg
COPY ${SRC_DIR}/node-red-256.png ${DEST_DIR}/node-red-256.png
COPY ${SRC_DIR}/node-red-icon.svg ${DEST_DIR}/node-red-icon.svg
COPY ${SRC_DIR}/node-red-icon-black.svg ${DEST_DIR}/node-red-icon-black.svg

# COPY deployment/nodered/images/deploy-full-o.svg /usr/src/node-red/node_modules/@node-red/editor-client/images/node-red-icon-black.svg

# Define argument with default value for theme.js source path
#ARG SRC_THEME=node-red/packages/node_modules/@node-red/editor-api/lib/editor
ARG SRC_THEME=deployment/nodered/editor
ARG DEST_THEME=/usr/src/node-red/node_modules/@node-red/editor-api/lib/editor
COPY ${SRC_THEME}/theme.js ${DEST_THEME}/theme.js

# CMD ["npm", "start", "--", "--userDir", "/data"]
