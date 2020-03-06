#!/bin/bash

NEXT_WAIT_TIME=0
until mysql -h ${ENV_DB_HOST} -P ${ENV_DB_PORT} -u ${ENV_DB_USER} -p${ENV_DB_PASSWORD} < /app/schema.sql || [ $NEXT_WAIT_TIME -eq 8 ]; do
   sleep $(( NEXT_WAIT_TIME++ ))
done
pause
