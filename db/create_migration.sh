#!/bin/sh
set -eu

echo "docker run -v $PWD/migrations:/migrations migrate/migrate create -ext sql -dir migrations -seq $1"
docker run -v $PWD/migrations:/migrations migrate/migrate create -ext sql -dir migrations -seq -dir migrations -seq $1
