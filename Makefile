exe:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build  -ldflags -H=windowsgui -o fantasy_tool.exe cmd/fantasy-tool/main.go

linux:
	go run cmd/fantasy-tool/main.go