#!/bin/bash

mysql -h ${ENV_DB_HOST} -P ${ENV_DB_PORT} -u ${ENV_DB_USER} -p${ENV_DB_PASSWORD} < /app/schema.sql
