.PHONY: all

all:get glide vendor install linter test run

.PHONY: get
get:
	curl https://glide.sh/get | sh

.PHONY: glide
glide:
	-glide update

.PHONY: vendor
vendor:
	govendor sync update

.PHONY: install
install:
	gometalinter.v2 --install

.PHONY: linter
linter:
	-gometalinter.v2 --vendor --tests --skip=mocks \
		--exclude='_gen.go' --exclude='docs.go' \
		--disable=gotype --disable=errcheck --disable=gas --disable=dupl \
		--deadline=1500s --checkstyle --sort=linter ./... > static-analysis.xml

.PHONY: test
test:
	-go test -v -cover ./src/services/... > tests.xml

.PHONY: build
build:
	go build -o gotoboox

.PHONY: run
run:
	go run main.go

.PHONY: prof
prof:
	cd src/services/ && go test -v -run=^$  -bench=^BenchmarkCommentsService_BookCommentsHandler$ -benchtime=2s -cpuprofile=prof.cpu
	cd src/services/ && go tool pprof services.test prof.cpu
