FROM node:10.10-alpine as build-stage
WORKDIR /app

COPY web .

RUN apk update && \
	npm install -g npm && \
    npm install && \
    npm rebuild node-sass

CMD ["/bin/sh"]