FROM openjdk:8-jre-slim

MAINTAINER c0ny1 <root@gv7.me>

RUN useradd --home-dir /home/webwolf --create-home -U webwolf && \
    apt-get update && \
    apt-get install curl -y && \
    apt-get install wget && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

COPY start.sh /home/webwolf/start.sh
RUN chmod +x /home/webwolf/start.sh
USER webwolf
RUN cd /home/webwolf && \
    wget -O webwolf.jar https://github.com/WebGoat/WebGoat/releases/download/v8.0.0.M14/webwolf-8.0.0.M14.jar