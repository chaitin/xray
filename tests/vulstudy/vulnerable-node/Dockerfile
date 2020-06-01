FROM ubuntu:xenial

MAINTAINER c0ny1 <root@gv7.me>

ENV STAGE "DOCKER"

RUN apt-get update && apt-get -y upgrade && \
    apt-get install -y nodejs && \
    apt-get install -y npm && \
    apt-get install -y netcat && \
    apt-get install -y git && \
    apt-get clean && \
    ln -s /usr/bin/nodejs /usr/bin/node && \
    cd /tmp/ && \
    git clone https://github.com/cr0hn/vulnerable-node.git && \
    cd ./vulnerable-node && \
    git checkout 8937dfbc012b4a76b99fb41ce14e29e95862fafb && \
    mkdir /app && \
    mv package.json /app/ && \
    cd /app && \
    npm install && \
    mv /tmp/vulnerable-node/* ./ && \
    chmod +x /app/start.sh && \
    rm -rf /var/lib/apt/lists/* && \
    rm -rf /tmp/vulnerable-node

WORKDIR /app

EXPOSE 3000
CMD [ "/app/start.sh" ]
