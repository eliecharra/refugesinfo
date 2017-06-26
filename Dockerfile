FROM node:8

# set the loglevel for npm with environment variable
ENV NPM_CONFIG_LOGLEVEL=warn

WORKDIR /app
COPY . /app
RUN npm install && npm run build

ENTRYPOINT [ "node", "lib/app.js" ]
