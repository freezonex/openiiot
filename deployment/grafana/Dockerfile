FROM grafana/grafana-enterprise

# Install plugins
RUN grafana cli plugins install grafana-mqtt-datasource && \
    grafana cli plugins install satellogic-3d-globe-panel && \
    grafana cli plugins install cloudspout-button-panel && \
    grafana cli plugins install grafana-iot-twinmaker-app && \
    grafana cli plugins install tdengine-datasource
    
# Update Logo
ARG SRC_DIR=deployment/grafana/img
ARG DEST_DIR=/usr/share/grafana/public/img
COPY ${SRC_DIR}/grafana_icon.svg ${DEST_DIR}/grafana_icon.svg

# Copy grafana.ini configuration file
# COPY deployment/grafana/conf/grafana.ini /etc/grafana/grafana.ini

# Update Login Title
# grep -rnwl './' -e 'Welcome to Grafana'
USER root
RUN find /usr/share/grafana/public/build/ -name "*.js" -exec sed -i 's|AppTitle="Grafana")|AppTitle="FreezoneX")|g' {} \;
RUN find /usr/share/grafana/public/build/ -name "*.js" -exec sed -i 's|LoginTitle="Welcome to Grafana")|LoginTitle="Welcome to FreezoneX")|g' {} \;