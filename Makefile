MAIN=main.go
IMAGENAME=damdo/docker-dtdns
default: builder

builder:
	docker run --rm -e LDFLAGS='-w -s -extldflags "-static"' -e OUTPUT=gobin  -v $(PWD):/src centurylink/golang-builder
	make dockerbuild

localbuilder:
	GOOS=linux GOARCH=amd64 go build -o gobin -ldflags "-w -s" $(MAIN)
	make dockerbuild

dockerbuild:
	docker build -t $(IMAGENAME) .

.PHONY: clean
clean:
	rm -f gobin
