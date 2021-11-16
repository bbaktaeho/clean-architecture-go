#!/bin/sh

docker run --name leo-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=9036 -d mysql:latest