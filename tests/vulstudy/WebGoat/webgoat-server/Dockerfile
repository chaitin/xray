FROM openjdk:8-jre-slim

MAINTAINER c0ny1 <root@gv7.me>

RUN useradd --home-dir /home/webgoat --create-home -U webgoat && \
    apt-get update && \
    apt-get install curl -y && \
    apt-get install wget && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

COPY start.sh /home/webgoat/start.sh
RUN chmod +x /home/webgoat/start.sh
USER webgoat
RUN cd /home/webgoat/ && \
    mkdir -p .webgoat && \
    wget -O webgoat.jar https://github.com/WebGoat/WebGoat/releases/download/v8.0.0.M14/webgoat-server-8.0.0.M14.jar
