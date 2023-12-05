FROM grafana/grafana-enterprise

# Update Logo
ARG SRC_DIR=deployment/grafana/img
ARG DEST_DIR=/usr/share/grafana/public/img
COPY ${SRC_DIR}/grafana_icon.svg ${DEST_DIR}/grafana_icon.svg

# Copy grafana.ini configuration file
COPY deployment/grafana/conf/grafana.ini /etc/grafana/grafana.ini

# Update Login Title
# grep -rnwl './' -e 'Welcome to Grafana'
USER root
RUN find /usr/share/grafana/public/build/ -name "*.js" -exec sed -i 's|AppTitle="Grafana")|AppTitle="FreezoneX")|g' {} \;
RUN find /usr/share/grafana/public/build/ -name "*.js" -exec sed -i 's|LoginTitle="Welcome to Grafana")|LoginTitle="Welcome to FreezoneX")|g' {} \;