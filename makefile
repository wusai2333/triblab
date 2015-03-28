.PHONY: all rall fmt tags test-lab1 test-lab2 test-lab3 testv-lab1 testv-lab2 testv-lab3 lc doc turnin

all:
	go install ./... trib/...

rall:
	go build -a ./... trib/...

fmt:
	gofmt -s -w -l .

tags:
	gotags `find . -name "*.go"` > tags

test-lab1:
	TRIB_LAB=lab1 go test ./...

testv-lab1:
	TRIB_LAB=lab1 go test -v ./...

test-lab2:
	TRIB_LAB=lab2 go test ./...

testv-lab2:
	TRIB_LAB=lab2 go test -v ./...

test-lab3:
	TRIB_LAB=lab3 go test ./...

testv-lab3:
	TRIB_LAB=lab3 go test -v ./...

lc:
	wc -l `find . -name "*.go"`

doc:
	godoc -http=:8000

turnin:
	@ echo "Turning in for `whoami`"
	git archive -o turnin.zip HEAD
	chmod 600 turnin.zip
	cp turnin.zip /classes/cse223b/sp14/labs/turnin/lab2/`whoami`.zip
