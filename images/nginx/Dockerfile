FROM nginx:alpine

ARG /var/log/nginx

RUN echo http://dl-cdn.alpinelinux.org/alpine/edge/community/ >> /etc/apk/repositories \
 && echo http://dl-cdn.alpinelinux.org/alpine/edge/main/ >> /etc/apk/repositories \
 && apk update \
 && apk upgrade \
 && apk add --no-cache nginx inotify-tools openssl

RUN     openssl req 	-x509 -nodes \
			-days 365 \
			-newkey rsa:2048 \
			-keyout server.key \
			-out server.crt  \
			-subj "/C=FR/ST=Aquitaine/L=Bordeaux/O=MeMyselfAndI/OU=IT Department/CN=webapp.local" \
	&& openssl dhparam -out dhparam.pem 2048 \
	&& mkdir /etc/certs \
	&& mv server.* /etc/certs \
	&& mv dhparam.pem /etc/certs \
	&& apk del openssl \
	&& rm -rf /var/cache/apk/* 


COPY nginx.conf /etc/nginx/
COPY default.conf /etc/nginx/conf.d/

RUN ln -sf /dev/stdout /var/log/nginx/access.log \
 && ln -sf /dev/stderr /var/log/nginx/error.log \
 && rm -rf /tmp/*

COPY reload.sh /
RUN chmod +x reload.sh
RUN mkdir -p /run/nginx
EXPOSE 80 443
CMD ["/reload.sh"]