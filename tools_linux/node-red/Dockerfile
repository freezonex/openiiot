FROM nodered/node-red

COPY node-red.svg /usr/src/node-red/node_modules/@node-red/editor-client/public/red/images/node-red.svg
COPY theme.js /usr/src/node-red/node_modules/@node-red/editor-api/lib/editor/theme.js

CMD ["npm", "start", "--", "--userDir", "/data"]
