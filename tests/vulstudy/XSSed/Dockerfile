FROM php:5.5-apache

MAINTAINER c0ny1 <root@gv7.me>

# set DirectoryIndex:index.htm
COPY docker-php.conf /etc/apache2/conf-enabled/

RUN apt-get update && \
    apt-get install -y git && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* && \
    cd /tmp/ && \
    git clone https://github.com/aj00200/xssed.git && \
    cd xssed && \
    mv * /var/www/html/ && \
    rm -rf /tmp/xssed