install:
	set CGO_ENABLED=0 
	go install .

push:
	git push -u origin main

add:
	pass-manager add -a Leetcode -e prasad28sumit@gmail.com -u sum28it -p Something secret

get:
	pass-manager get secret
	
test:
	pass-manager init "secret"
	pass-manager init "secret"
	pass-manager get -a Leetcode secret
	pass-manager add -a Leetcode -e prasad28sumit@gmail.com -u sum28it -p Something secret
	pass-manager add -a Github -e prasad28sumit@gmail.com -u sum28it -p Whatever secret
	pass-manager get -a Leetcode secret
	pass-manager delete -a Leetcode secret
	pass-manager get -a Leetcode secret

reset: 
	pass-manager reset secret

build-win-amd64:
	set CGO_ENABLED=0 
	set GOOS=windows
	set GOARCH=amd64
	go build -o bin/passm-win-amd64.exe

build-linux-amd64:
	set CGO_ENABLED=0
	set GOOS=linux
	set GOARCH=amd64
	go build -o bin/passm-linux-amd64.exe

build-win-arm:
	set CGO_ENABLED=0 
	set GOOS=windows
	set GOARCH=arm
	go build -o bin/passm-win-arm.exe

build-linux-arm:
	set CGO_ENABLED=0 
	set GOOS=linux
	set GOARCH=arm
	go build -o bin/passm-linux-arm.exe

build-all:
	make build-win-amd64
	make build-win-arm
	make build-linux-amd64
	make build-linux-arm
