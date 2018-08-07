#!/bin/bash
VERSION=1.5.0
docker build . -t docker.io/cenoq/gitea:${VERSION}
docker tag docker.io/cenoq/gitea:${VERSION} docker.io/cenoq/gitea:latest
docker push docker.io/cenoq/gitea:${VERSION}
docker push docker.io/cenoq/gitea:latest
