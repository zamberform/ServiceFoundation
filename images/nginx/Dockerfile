FROM nginx:alpine

ARG /var/log/nginx

COPY nginx.conf /etc/nginx/
COPY default.conf /etc/nginx/conf.d/

RUN echo http://dl-cdn.alpinelinux.org/alpine/edge/community/ >> /etc/apk/repositories \
 && echo http://dl-cdn.alpinelinux.org/alpine/edge/main/ >> /etc/apk/repositories \
 && apk update && apk upgrade \
 \
 # forward request and error logs to docker log collector
 && ln -sf /dev/stdout /var/log/nginx/access.log \
 && ln -sf /dev/stderr /var/log/nginx/error.log \
 && rm -rf /tmp/*