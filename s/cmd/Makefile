all: clean build
clean:
	-rm s-linux64
	-rm s-mac64
	-rm s-win64.exe
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o s-linux64
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o s-mac64
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o s-win64.exe