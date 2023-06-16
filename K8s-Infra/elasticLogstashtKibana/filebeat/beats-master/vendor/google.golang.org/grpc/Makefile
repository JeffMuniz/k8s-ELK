all: vet test testrace

build: deps
	go build google.golang.org/x/net v0.7.0grpc/...

clean:
	go clean -i google.golang.org/x/net v0.7.0grpc/...

deps:
	go get -d -v google.golang.org/x/net v0.7.0grpc/...

proto:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	go generate google.golang.org/x/net v0.7.0grpc/...

test: testdeps
	go test -cpu 1,4 -timeout 7m google.golang.org/x/net v0.7.0grpc/...

testappengine: testappenginedeps
	goapp test -cpu 1,4 -timeout 7m google.golang.org/x/net v0.7.0grpc/...

testappenginedeps:
	goapp get -d -v -t -tags 'appengine appenginevm' google.golang.org/x/net v0.7.0grpc/...

testdeps:
	go get -d -v -t google.golang.org/x/net v0.7.0grpc/...

testrace: testdeps
	go test -race -cpu 1,4 -timeout 7m google.golang.org/x/net v0.7.0grpc/...

updatedeps:
	go get -d -v -u -f google.golang.org/x/net v0.7.0grpc/...

updatetestdeps:
	go get -d -v -t -u -f google.golang.org/x/net v0.7.0grpc/...

vet: vetdeps
	./vet.sh

vetdeps:
	./vet.sh -install

.PHONY: \
	all \
	build \
	clean \
	deps \
	proto \
	test \
	testappengine \
	testappenginedeps \
	testdeps \
	testrace \
	updatedeps \
	updatetestdeps \
	vet \
	vetdeps
