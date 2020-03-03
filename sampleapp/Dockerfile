FROM quay.io/ikezoe/corretto:11

COPY rest-service-0.0.1-SNAPSHOT.jar /app/rest-service.jar
COPY jvm.options /app/jvm.options
COPY run.sh /app/run.sh


CMD ["/app/run.sh"]
