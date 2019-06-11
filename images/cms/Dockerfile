FROM node:10.10-alpine as build-stage
WORKDIR /app

COPY cms .

RUN npm i -g yarn

CMD yarn install && yarn run dev