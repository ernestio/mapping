install:
	go install -v

build:
	go build -v ./...

lint:
	gometalinter --config .linter.conf

test:
	go test -v ./... --cover
	megacheck ./...
	unconvert ./...

deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

dev-deps: deps
	go get github.com/alecthomas/gometalinter
	go get github.com/mdempsky/unconvert
	go get honnef.co/go/tools/cmd/...
	gometalinter --install

clean:
	go clean

dist-clean:
	rm -rf pkg src bin
