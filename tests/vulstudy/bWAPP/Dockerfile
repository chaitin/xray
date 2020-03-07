FROM tutum/lamp:latest

MAINTAINER c0ny1 <root@gv7.me>

RUN apt-get update && \
    apt-get install -y wget zip && \
    apt-get clean && \
    rm /app/* && \
    cd /tmp && \
    wget https://jaist.dl.sourceforge.net/project/bwapp/bWAPP/bWAPPv2.2/bWAPPv2.2.zip &&\
    unzip ./bWAPPv2.2.zip && \
    mv ./bWAPP/* /app/ && \
    rm -rf /tmp/* && \
    rm -rf /var/lib/apt/lists/*

CMD ["/run.sh"]