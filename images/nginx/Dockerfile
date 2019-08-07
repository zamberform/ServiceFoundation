FROM node:12-alpine as build-stage
WORKDIR /app

COPY cms .
RUN apk update && \
    npm install -g npm
RUN apk add --no-cache git
RUN npm i -g yarn

RUN yarn cache clean && yarn && yarn run build:prod

FROM nginx:alpine as production-stage

ARG /var/log/nginx

RUN echo http://dl-cdn.alpinelinux.org/alpine/edge/community/ >> /etc/apk/repositories \
 && echo http://dl-cdn.alpinelinux.org/alpine/edge/main/ >> /etc/apk/repositories \
 && apk update \
 && apk upgrade

COPY images/nginx/nginx.conf /etc/nginx/
COPY --from=build-stage /app/dist /var/www/html
RUN ln -s /var/www/html/images /etc/nginx/html

RUN ln -sf /dev/stdout /var/log/nginx/access.log \
 && ln -sf /dev/stderr /var/log/nginx/error.log \
 && rm -rf /tmp/*