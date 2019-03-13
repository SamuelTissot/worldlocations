#!/usr/bin/env sh

set -e;
printf "booting worldlocations.io ENVIRONMENT: %s, DB_INSTANCE: %s\n" $GO_ENV $SQL_INSTANCES_NAMES;
/app_src/cloud_sql_proxy -instances=$SQL_INSTANCES_NAMES -credential_file=/app_src/wlio-sqlproxy-sa.json | >&1 \
                 exec /app_src/app;