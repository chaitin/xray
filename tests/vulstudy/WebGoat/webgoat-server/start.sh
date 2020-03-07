#!/bin/sh

java -jar -Djava.security.egd=file:/dev/./urandom /home/webgoat/webgoat.jar --server.address=0.0.0.0 --server.port=8080
