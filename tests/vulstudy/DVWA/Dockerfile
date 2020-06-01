FROM tutum/lamp:latest

MAINTAINER c0ny1 <root@gv7.me>

ADD . /tmp/
RUN apt-get update && \
    apt-get install -y libgd-dev && \
    apt-get install -y php5-gd && \
    rm -rf /var/lib/apt/lists/*

RUN rm /app/* && \
    cd /tmp/ && \
    cp php.ini /etc/php5/apache2/php.ini && \
    cp php.ini /etc/php5/cli/php.ini && \
    wget https://github.com/ethicalhack3r/DVWA/archive/v1.9.tar.gz && \
    tar xvf v1.9.tar.gz && \
    mv ./DVWA-1.9/* /app/ && \
    chown www-data:www-data -R /app/ && \
    chmod +x run.sh && \
    ./run.sh && \
    rm -rf /tmp/* && \

