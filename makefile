.PHONY: all rall fmt tags test testv lc doc turnin

all:
	go install ./...

rall:
	go build -a ./...

fmt:
	gofmt -s -w -l .

tags:
	gotags `find . -name "*.go"` > tags

test:
	go test ./...

testv:
	go test -v ./...

lc:
	wc -l `find . -name "*.go"`

doc:
	godoc -http=:8000

turnin:
	git archive -o turnin.zip HEAD
	chmod 600 turnin.zip
	cp turnin.zip /classes/cse223b/sp14/labs/turnin/lab1/`whoami`.zip
