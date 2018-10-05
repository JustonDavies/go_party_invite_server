build:
#	dep ensure
	rm -f bin/*

	env GOOS=linux go build -ldflags="-s -w" -o bin/fhtagn src/main.go

	chmod -R 777 bin