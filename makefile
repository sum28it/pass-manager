install:
	set CGO_ENABLED=0 
	go install -ldflags "-X github.com/sum28it/pass-manager/pkg/user.Dir=C:\\users\\user\\Downloads\\" .

