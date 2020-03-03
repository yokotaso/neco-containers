FROM quay.io/cybozu/ubuntu:18.04

ARG CORRETTO_VERSION=11.0.3.7.1
ARG CORRETTO_DEB=java-11-amazon-corretto-jdk_11.0.3.7-1_amd64.deb

RUN apt-get update && apt-get install -y --no-install-recommends \
      java-common \
    && rm -rf /var/lib/apt/lists/*
RUN curl -o ${CORRETTO_DEB} https://d3pxv6yz143wms.cloudfront.net/${CORRETTO_VERSION}/${CORRETTO_DEB} \
    && dpkg --install ${CORRETTO_DEB} \
    && rm ${CORRETTO_DEB}

ENV LANG C.UTF-8
ENV JAVA_HOME /usr/lib/jvm/java-11-amazon-corretto