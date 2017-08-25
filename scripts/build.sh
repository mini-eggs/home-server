#!/bin/bash

var=$(pwd)

source .env
export $(cut -d= -f1 .env)

# kill
docker kill $(docker ps -a -q) | true
docker stop $(docker ps -a -q) | true
docker rm $(docker ps -a -q) | true

# mysql
docker run --name mysql -d -p 3306:3306 \
-e MYSQL_USER=$MYSQL_ROOT_PASSWORD \
-e MYSQL_PASSWORD=$MYSQL_PASSWORD \
-e MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD \
-e MYSQL_DATABASE=$MYSQL_DATABASE \
hypriot/rpi-mysql

# phpmyadmin
docker run --name apache -d -i -t -p 8080:80 \
--link mysql:mysqldb \
-v $var/html:/var/www/html \
deigray/rpi-lamp:1.4
