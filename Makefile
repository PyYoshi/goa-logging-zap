.PHONY: all
all: test

.PHONY: deps
deps: glide
	./glide install

.PHONY: update-deps
update-deps:
	bash -c "glide update --all-dependencies --strip-vendor --resolve-current"

.PHONY: glide
glide:
ifeq ($(shell uname),Darwin)
	curl -L https://github.com/Masterminds/glide/releases/download/v0.12.3/glide-v0.12.3-darwin-amd64.zip -o glide.zip
	unzip glide.zip
	mv ./darwin-amd64/glide ./glide
	rm -fr ./darwin-amd64
	rm ./glide.zip
else
	curl -L https://github.com/Masterminds/glide/releases/download/v0.12.3/glide-v0.12.3-linux-amd64.zip -o glide.zip
	unzip glide.zip
	mv ./linux-amd64/glide ./glide
	rm -fr ./linux-amd64
	rm ./glide.zip
endif

.PHONY: test
test:
	go test
