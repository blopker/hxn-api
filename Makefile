all: build

build:
	go build main.go

run:
	go run *.go

deploy:
	git push dokku master
