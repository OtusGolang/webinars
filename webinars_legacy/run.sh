#!/usr/bin/env bash
docker run -it --rm -v `pwd`:/var/www/ -p 8080:80 busybox httpd -f -h /var/www/
