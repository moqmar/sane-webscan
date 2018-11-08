build:
	go-bindata web
	go build -o sane-webscan *.go
build-dev:
	go-bindata --debug web
	go build -o sane-webscan *.go
run: build-dev
	./sane-webscan
