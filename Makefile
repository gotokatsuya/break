API_ONLY_PKGS=$(shell go list ./... 2> /dev/null | grep -v "/vendor/")

PRJDIR=github.com/gotokatsuya/break
PRJTARGET=break
BINDIR=misc/bin

install:
	@go get github.com/kaneshin/lime
	@glide install

build:
	@go build -v -o $(BINDIR)/$(PRJTARGET) -i $(PRJDIR)

run:
	@./$(BINDIR)/$(PRJTARGET)

run-hot:
	@lime -bin=$(BINDIR)/$(PRJTARGET) -port="1234" -app-port="8888" -ignore-pattern="(\\.git)" -immediate=true
	
test:
	@go test $(API_ONLY_PKGS)

vet:
	@echo "go vet packages"
	@go tool vet -all -structtags -shadow $(shell ls -d */ | grep -v "vendor")
