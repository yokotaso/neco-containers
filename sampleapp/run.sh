#!/bin/sh
set -uex

JVM_OPTIONS="$(cat < /app/jvm.options | grep -E -v "^#.*" | tr '\n' ' ')"
ARGS="$(eval echo "${JVM_OPTIONS}")"
env | grep -iv "PASSWORD"
exec "${JAVA_HOME}/bin/java" ${ARGS} -jar /app/rest-service.jar
