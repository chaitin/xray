FROM tutum/lamp:latest

MAINTAINER c0ny1 <root@gv7.me>

#COPY . /tmp/
RUN cd /tmp/ && \
    git clone -b master https://github.com/Hackademic/hackademic.git

RUN rm -rf /app/*  && \
    cp -r /tmp/hackademic/* /app/ && \
    chown www-data:www-data -R /app/ && \
    rm -rf /tmp/hackademic