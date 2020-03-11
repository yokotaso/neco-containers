#!/bin/bash

NEXT_WAIT_TIME=0
until mysql -h ${MYSQL_HOST} -P ${MYSQL_PORT} -u ${MYSQL_USER} -p${MYSQL_PASSWORD} < /app/schema.sql || [ $NEXT_WAIT_TIME -eq 8 ]; do
   sleep $(( NEXT_WAIT_TIME++ ))
done
pause
