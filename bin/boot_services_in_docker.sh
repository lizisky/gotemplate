#!/bin/bash

docker stop $(docker ps -aq)
docker rm $(docker ps -aq)

# command on my linux
# docker run -d -p 3306:3306  -v /work/data/data_mysql:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=lizisky001 --name pet-mysql mysql:8.0 --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci 

# command on my iMac
docker run -d -p 3306:3306  -v /Users/andy/pet/data/data_mysql:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=lizisky001 --name pet-mysql mysql:8.0 --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci 
