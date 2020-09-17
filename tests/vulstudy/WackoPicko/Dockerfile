FROM tutum/lamp:latest

MAINTAINER c0ny1 <c0ny1>

RUN apt-get update && \
    apt-get install -y libgd-dev && \
    apt-get install -y php5-gd && \
    apt-get clean && \
    cd /tmp/ && \
    git clone https://github.com/adamdoupe/WackoPicko.git && \
    cd WackoPicko && \
    git checkout 065cb92aceb6f76138786e94959034014e733b99 && \
    rm -rf /app/* && \
    mv -f /tmp/WackoPicko/website/* /app/ && \
    chmod 777 /app/upload && \
    cp current.sql / && \
    cp create_mysql_admin_user.sh / && \
    cp php.ini /etc/php5/apache2/php.ini && \
    cp php.ini /etc/php5/cli/php.ini && \
    chmod 755 /*.sh && \
    rm -rf /var/lib/apt/lists/* && \
    rm -rf /tmp/WackoPicko

CMD ["/run.sh"]
