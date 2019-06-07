FROM nginx:alpine

ARG /var/log/nginx

RUN echo http://dl-cdn.alpinelinux.org/alpine/edge/community/ >> /etc/apk/repositories \
 && echo http://dl-cdn.alpinelinux.org/alpine/edge/main/ >> /etc/apk/repositories \
 && apk update \
 && apk upgrade

COPY images/nginx/nginx.conf /etc/nginx/

RUN ln -sf /dev/stdout /var/log/nginx/access.log \
 && ln -sf /dev/stderr /var/log/nginx/error.log \
 && rm -rf /tmp/*