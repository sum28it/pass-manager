install:
	set CGO_ENABLED=0
	go install -ldflags "-X github.com/sum28it/pass-manager/cmd.CommitId=${shell git rev-parse HEAD}" .

tidy:
	go mod tidy	
	go mod vendor

push:
	git push -u origin main

add:
	pass-manager add -a Leetcode -e prasad28sumit@gmail.com -u sum28it

get:
	pass-manager get

test:
	pass-manager init 
	pass-manager init 
	pass-manager get -a Leetcode 
	pass-manager add -a Leetcode -e prasad28sumit@gmail.com -u sum28it
	pass-manager add -a Github -e prasad28sumit@gmail.com -u sum28it
	pass-manager get -a Leetcode 
	pass-manager delete -a Leetcode 
	pass-manager get -a Leetcode 

reset: 
	pass-manager reset

build-win-amd64:
	set CGO_ENABLED=0 
	set GOOS=windows
	set GOARCH=amd64
	go build -o bin/passm-win-amd64.exe -ldflags "-X github.com/sum28it/pass-manager/cmd.CommitId=$(shell git rev-parse HEAD)"

build-linux-amd64:
	set CGO_ENABLED=0
	set GOOS=linux
	set GOARCH=amd64
	go build -o bin/passm-linux-amd64.exe -ldflags "-X github.com/sum28it/pass-manager/cmd.CommitId=$(shell git rev-parse HEAD)"

build-win-arm:
	set CGO_ENABLED=0 
	set GOOS=windows
	set GOARCH=arm
	go build -o bin/passm-win-arm.exe -ldflags "-X github.com/sum28it/pass-manager/cmd.CommitId=$(shell git rev-parse HEAD)"

build-linux-arm:
	set CGO_ENABLED=0 
	set GOOS=linux
	set GOARCH=arm
	go build -o bin/passm-linux-arm.exe -ldflags "-X github.com/sum28it/pass-manager/cmd.CommitId=$(shell git rev-parse HEAD)"

build-all:
	make build-win-amd64
	make build-win-arm
	make build-linux-amd64
	make build-linux-arm

docker-build: 
	docker build -t sum28it/pass-manager .

docker-run:
	docker run --name pass-manager -it sum28it/pass-manager
