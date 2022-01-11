#!/bin/sh
set -eu

source .env

MYSQL_URL="mysql://landin:landin@tcp(127.0.0.1:3306)/landin"
echo "docker run -v $PWD/migrations:/migrations --network host migrate/migrate -database ${MYSQL_URL} -path $PWD/database/migrations $1 $2"
docker run -v $PWD/migrations:/migrations --network host migrate/migrate -database ${MYSQL_URL} -path /migrations $1 $2
