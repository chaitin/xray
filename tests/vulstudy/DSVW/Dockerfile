FROM python:2.7-jessie

MAINTAINER c0ny1 <root@gv7.me>

RUN pip install lxml && \
    mkdir /app && \
    cd /app && \
    wget https://github.com/stamparm/DSVW/raw/master/dsvw.py
WORKDIR /app/

EXPOSE 65412
CMD python dsvw.py
