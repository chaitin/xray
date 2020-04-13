FROM php:5.6.13-apache

MAINTAINER c0ny1 <root@gv7.me>

RUN apt-get update && \
    apt-get install -y php5-xsl && \
    apt-get install -y php5-mcrypt && \
    apt-get install -y libmcrypt-dev && \
    apt-get install -y libxslt1-dev && \
    apt-get install -y git && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN docker-php-ext-install mcrypt && \
    docker-php-ext-install xsl && \
    docker-php-ext-install mysql

RUN cd /tmp/ && \
    git clone https://github.com/SpiderLabs/MCIR.git && \
    cd MCIR &&\
    git checkout 8ca70207b692ceaf72d5a60653f6d1d83cce88ef && \
    rm -rf /var/www/html/* && \
    mv * /var/www/html/ && \
    cd /var/www/html/ && \
    sed -i "s/default_mcir_db_password/mcirpass00112233/" sqlol/includes/database.config.php && \
    sed -i "s/default_mcir_db_password/mcirpass00112233/" cryptomg/includes/db.inc.php && \
    sed -i "s/localhost/mysqldb/" sqlol/includes/database.config.php && \
    sed -i "s/localhost/mysqldb/" cryptomg/includes/db.inc.php && \
    chmod 666 xssmh/pxss.html && \
    rm -rf /var/lib/apt/lists/* && \
    rm -rf /tmp/MCIR
