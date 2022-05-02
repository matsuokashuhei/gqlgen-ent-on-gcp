#!/usr/bin/env bash
set -eu

if [ -f .env ]
then
  source .env
else
  echo ".env not found"
  exit 1
fi

# ID=$(docker compose ps --filter status=running -q mysql)
docker compose ps --filter status=running -q mysql
ID=$(docker compose ps --filter status=running -q mysql)

if [ -z $ID ]
then
  echo "mysql not running"
  exit 1
fi

MYSQL_DSN="mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}"
echo "MYSQL_DSN: ${MYSQL_DSN}"

DOCKER_NETWORK=landin_default
docker run -v $PWD/backend/migrations:/migrations \
           --network ${DOCKER_NETWORK} \
           migrate/migrate \
           -database ${MYSQL_DSN} \
           -path /migrations \
           up