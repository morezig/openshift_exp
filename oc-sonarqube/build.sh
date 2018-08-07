#!/bin/bash
VERSION=7.2
docker build . -t docker.io/cenoq/sonarqube:${VERSION}
docker tag docker.io/cenoq/sonarqube:${VERSION} docker.io/cenoq/sonarqube:latest
docker push docker.io/cenoq/sonarqube:${VERSION}
docker push docker.io/cenoq/sonarqube:latest
