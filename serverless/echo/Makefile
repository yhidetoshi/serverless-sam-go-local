PHONY: deps clean build

deps:
	go get -u ./...

clean:
	rm -rf echo # 変更

build:
	GOOS=linux GOARCH=amd64 go build -o echo ./main.go
