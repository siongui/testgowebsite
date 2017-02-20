# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../go1.7.5)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)

GO_VERSION=1.7.5
DEV_DIR=../
STATIC_WEBSITE_DIR=website

devserver: check_dir test
	@echo "\033[92mDevelopment Server Running ...\033[0m"
	@go run devserver/server.go -dir=$(STATIC_WEBSITE_DIR) -port=8000

test: fmt
	@go test -v

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt *.go
	@go fmt devserver/*.go

check_dir:
	@echo "\033[92mCreate website directory if not exists ...\033[0m"
	@[ -d $(STATIC_WEBSITE_DIR) ] || mkdir -p $(STATIC_WEBSITE_DIR)

update_ubuntu:
	@echo "\033[92mUpdating Ubuntu ...\033[0m"
	@sudo apt-get update && sudo apt-get upgrade && sudo apt-get dist-upgrade

download_go:
	@echo "\033[92mDownloading and Installing Go ...\033[0m"
	@cd $(DEV_DIR) ; wget https://storage.googleapis.com/golang/go$(GO_VERSION).linux-amd64.tar.gz
	@cd $(DEV_DIR) ; tar xvzf go$(GO_VERSION).linux-amd64.tar.gz

install:
	@echo "\033[92mInstalling Go template utility ...\033[0m"
	go get -u github.com/siongui/gotemplateutil

clean:
	@rm -rf pkg/ src/ website/
