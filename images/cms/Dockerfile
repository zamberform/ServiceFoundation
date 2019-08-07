FROM node:12-alpine as build-stage
WORKDIR /app

COPY cms .
RUN apk update && \
    npm install -g npm
RUN apk add --no-cache git
RUN npm i -g yarn

RUN yarn cache clean && yarn && yarn run build:prod

FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]

