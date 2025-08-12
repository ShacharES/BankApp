#!/bin/sh

set -e

echo "run db migration"
. /app/app.env   # or: . ./app.env  (since WORKDIR is /app)

# fail fast if missing
: "${DB_SOURCE:?DB_SOURCE is not set}"

/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"