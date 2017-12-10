SUBPACKAGES := $(shell go list ./...)

.PHONY: deps test vet lint

deps:
	dep ensure

test:
	go test -v $(SUBPACKAGES)

vet:
	go vet $(SUBPACKAGES)

lint:
	golint $(SUBPACKAGES)
