FROM node:10.10-alpine as build-stage
WORKDIR /app

COPY web .

RUN npm i -g yarn

CMD yarn install --ignore-engines && yarn run build && yarn run start