#!/bin/sh

java -jar -Djava.security.egd=file:/dev/./urandom /home/webwolf/webwolf.jar --server.address=0.0.0.0 --server.port=8081 
