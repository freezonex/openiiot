FROM nginx:latest

ARG SRC_DIR=deployment/web/html
ARG DEST_DIR=/usr/share/nginx/html

# Update html
COPY ${SRC_DIR}/ ${DEST_DIR}/

# html image
ARG SRC_IMAGE=deployment/web/html/images
ARG DEST_IMAGE=/usr/share/nginx/html/images
COPY ${SRC_IMAGE}/ ${DEST_IMAGE}/