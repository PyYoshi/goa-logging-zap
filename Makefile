.PHONY: all
all: test

.PHONY: deps
deps: glide
	./glide install

.PHONY: glide
glide:
ifeq ($(shell uname),Darwin)
	curl -L https://github.com/Masterminds/glide/releases/download/0.11.1/glide-darwin-amd64.zip -o glide.zip
	unzip glide.zip
	mv ./darwin-amd64/glide ./glide
	rm -fr ./darwin-amd64
	rm ./glide.zip
else
	curl -L https://github.com/Masterminds/glide/releases/download/v0.11.1/glide-v0.11.1-linux-amd64.zip -o glide.zip
	unzip glide.zip
	mv ./linux-amd64/glide ./glide
	rm -fr ./linux-amd64
	rm ./glide.zip
endif

.PHONY: test
test:
	go test
