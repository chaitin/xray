FROM tutum/lamp:latest

MAINTAINER c0ny1 <root@gv7.me>

RUN cd /tmp/ &&\
    git clone https://github.com/Audi-1/sqli-labs.git && \
    cd sqli-labs && \
    git checkout e96f21776372c8613a7e565106e62bc01a59355e && \
    rm -rf /app/* && \
    mv -f /tmp/sqli-labs/* /app/ && \
    chown www-data:www-data -R /app && \
    rm -rf /tmp/sqli-labs

EXPOSE 80 3306
