# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)

GO_VERSION=1.8.1
DEV_DIR=../
STATIC_WEBSITE_DIR=website

devserver: test
	@echo "\033[92mDevelopment Server Running ...\033[0m"
	@go run devserver/server.go -dir=$(STATIC_WEBSITE_DIR) -port=8000

test: fmt check_dir
	@go test -v

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt *.go
	@go fmt devserver/*.go

check_dir:
	@echo "\033[92mCreate website directory if not exists ...\033[0m"
	@[ -d $(STATIC_WEBSITE_DIR) ] || mkdir -p $(STATIC_WEBSITE_DIR)

github_pages: clean_static_site test
	@echo "\033[92mCreate CNAME for GitHub Pages custom domain ...\033[0m"
	@echo "watpahphotiyan.online-dhamma.net" > $(STATIC_WEBSITE_DIR)/CNAME
	@echo "\033[92mPublish to GitHub Pages ...\033[0m"
	ghp-import $(STATIC_WEBSITE_DIR); git push origin gh-pages

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
	@echo "\033[92mInstalling YAML parser ...\033[0m"
	go get -u github.com/ghodss/yaml
	@echo "\033[92mInstalling goquery (HTML parser) ...\033[0m"
	go get -u github.com/PuerkitoBio/goquery

clean_static_site:
	@echo "\033[92mRemove static site ...\033[0m"
	@rm -rf $(STATIC_WEBSITE_DIR)

clean: clean_static_site
	@rm -rf pkg/ src/
