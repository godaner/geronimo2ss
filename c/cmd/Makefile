all: build clean_upx
build:build_linux64 build_win64 build_mac64
build_linux64:
	-rm c-linux64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o c-linux64
	upx -9 c-linux64
build_win64:
	-rm c-win64.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o c-win64.exe
	upx -9 c-win64.exe
build_mac64:
	-rm c-mac64
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o c-mac64
	upx -9 c-mac64
clean_upx:
	-rm *.upx