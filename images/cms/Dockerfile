FROM node:10.10-alpine as build-stage
WORKDIR /app

COPY cms .

RUN npm i -g yarn
RUN apk add --no-cache git

CMD yarn --pure-lockfile && yarn cache clean && yarn install && yarn run dev