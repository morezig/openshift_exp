#!/bin/sh
echo "Build Start....."
sh -c "cd ui/backend && npm install"
sh -c "cd ui/backend && npm run build"
cd api/core
govendor list
govendor fetch +m
go-bindata-assetfs -pkg common -o common/bindata.go assets/...
