# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)

GO_VERSION=1.7.5
DEV_DIR=../

update_ubuntu:
	@echo "\033[92mUpdating Ubuntu ...\033[0m"
	@sudo apt-get update && sudo apt-get upgrade && sudo apt-get dist-upgrade

download_go:
	@echo "\033[92mDownloading and Installing Go ...\033[0m"
	@cd $(DEV_DIR) ; wget https://storage.googleapis.com/golang/go$(GO_VERSION).linux-amd64.tar.gz
	@cd $(DEV_DIR) ; tar xvzf go$(GO_VERSION).linux-amd64.tar.gz
