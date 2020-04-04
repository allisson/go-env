.PHONY: lint test

lint:
	if [ ! -f ${GOPATH}/bin/revive ] ; \
	then \
		go get -u github.com/mgechev/revive; \
	fi;
	${GOPATH}/bin/revive -config revive.toml -formatter friendly .

test:
	go test -covermode=count -coverprofile=count.out -v ./...
