#!/usr/bin/env bash
docker run -it --rm -v `pwd`:/web -p 8000:8080 halverneus/static-file-server