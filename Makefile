all: build

build:
	go build

run:
	go run *.go

deploy:
	git push dokku master

clean:
	rm hxn-api
