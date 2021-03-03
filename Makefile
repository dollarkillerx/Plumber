build:
	CGO_ENABLED=0 GOOS="linux" GOARCH="amd64" go build -o cmd/plumber cmd/main.go

upx:
	upx cmd/plumber

docker_build:
	docker rmi -f plumber:latest
	docker build -f cmd/Dockerfile -t plumber:latest  .
	docker save -o plumber.tar plumber:latest

