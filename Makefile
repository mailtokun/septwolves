GO_DOCKER_ENV = CGO_ENABLED=0 GOOS=linux GOARCH=amd64
IMAGE_REGISTRY = mailtokun/yutu
IMAGE_TAG = $(shell git log -1 --pretty=%h)
build2:
	go build -o main main.go
	$(GO_DOCKER_ENV) go build -ldflags "-s -w" -gcflags=-trimpath=${GOPATH} -asmflags=-trimpath=${GOPATH} -o ./main ./main.go
	docker build -t $(IMAGE_REGISTRY):$(IMAGE_TAG) .
	docker tag $(IMAGE_REGISTRY):$(IMAGE_TAG) $(IMAGE_REGISTRY):latest
	docker push $(IMAGE_REGISTRY):latest
	docker push $(IMAGE_REGISTRY):$(IMAGE_TAG)
	@echo "done"
run-rainbow:
	docker stop yutu || true
	docker rm -f yutu || true
	docker run -d --network="host" --name=yutu -v /var/run/docker.sock:/var/run/docker.sock --env GITHUB_REPO=$(GITHUB_REPO) --env GITHUB_BRANCH=$(GITHUB_BRANCH) --env GITHUB_TOKEN=$(GITHUB_TOKEN) $(IMAGE_REGISTRY):$(IMAGE_TAG) /yutu/main
