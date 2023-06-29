install:
	set CGO_ENABLED=0 
	go install -ldflags "-X github.com/sum28it/pass-manager/pkg/user.Dir=C:\\users\\user\\Downloads\\" .

push:
	git push -u origin main

add: pass-manager add -a Leetcode -e prasad28sumit@gmail.com -u sum28it -p Something secret

delete: pass-manager 

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
	set GOOS=windows
	set GOARCH=amd64
	go build -o bin/passm.exe

build-linux-amd64:
	set GOOS=linux
	set GOARCH=amd64
	go build -o bin/passm-linux-amd64.exe