FROM quay.io/ysksuzuki/corretto:11

RUN apt-get update && apt-get install -y --no-install-recommends mysql-client sysbench
COPY ./build/libs/demo-0.0.1-SNAPSHOT.jar /app/demo.jar
COPY resources/ /

CMD ["/app/run.sh"]
