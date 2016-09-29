MAIN=main.go
IMAGENAME=malgaboy/godtdns
default: gobin dockerbuild run

gobin: 
	GOOS=linux GOARCH=amd64 go build -o $@ -ldflags "-w -s" $(MAIN) 

dockerbuild: 
	docker build -t $(IMAGENAME) .	

run:
	docker run --rm -it $(IMAGENAME)

.PHONY: clean
clean:
	rm -f gobin
