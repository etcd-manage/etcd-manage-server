default:
	@echo 'Usage of make: [ build | linux_build | windows_build | docker_build | docker_run | clean ]'

build: 
	@go build -o ./bin/ems ./

linux_build: 
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/ems ./

windows_build: 
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/ems.exe ./

docker_build: linux_build
	docker build -t shiguanghuxian/etcd-manage .

docker_run: docker_build
	docker-compose up --force-recreate

run: build
	@./bin/ems

install: build
	@mv ./bin/ems $(GOPATH)/bin/ems

clean: 
	@rm -f ./bin/ems*
	@rm -f ./bin/logs/*

.PHONY: default build linux_build windows_build docker_build docker_run clean