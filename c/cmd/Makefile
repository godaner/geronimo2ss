all:clean build
clean:
	-rm c-linux64
	-rm c-mac64
	-rm c-win64.exe
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o c-linux64
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o c-mac64
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o c-win64.exe