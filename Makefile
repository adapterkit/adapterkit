GOPKG ?=	moul.io/adapterkit
DOCKER_IMAGE ?=	moul/adapterkit
GOBINS ?=	.
NPM_PACKAGES ?=	.

include rules.mk

generate: install
	GO111MODULE=off go get github.com/campoy/embedmd
	mkdir -p .tmp
	echo 'foo@bar:~$$ adapterkit' > .tmp/usage.txt
	adapterkit 2>&1 >> .tmp/usage.txt
	embedmd -w README.md
	rm -rf .tmp
.PHONY: generate
