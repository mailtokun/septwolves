#!/bin/bash
echo "Using GOPATH: $GOPATH"
go mod init yutu
go mod tidy
export GO111MODULE=on


go get -insecure ./...


echo "Building service ..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./main ./main.go
if [ "$?" -ne 0 ]; then
    echo "Failed to build.";
    exit 1;
fi
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o ./main_darwin ./main.go
if [ "$?" -ne 0 ]; then
    echo "Failed to build.";
    exit 1;
fi
echo "Creating Docker image ..."
docker_reg="mailtokun"
image_name="yutu"
d=`date +%Y%m%d%H%M`
echo "${docker_reg}/${image_name}:${d}"
docker build -t ${docker_reg}/${image_name}:${d} .
#docker push ${docker_reg}/${image_name}:${d}
echo "Would you like tag this image as latest and push it to production? (Y/n)"
docker tag ${docker_reg}/${image_name}:${d} ${docker_reg}/${image_name}:latest
echo "[DONE]"
