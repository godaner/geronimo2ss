all: build
build:build_linux64 build_win64 build_mac64
build_linux64:
	-rm s-linux64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o s-linux64
	upx -9 s-linux64
build_win64:
	-rm s-win64.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o s-win64.exe
	upx -9 s-win64.exe
build_mac64:
	-rm s-mac64
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o s-mac64
	upx -9 s-mac64