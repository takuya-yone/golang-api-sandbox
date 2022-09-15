windows:
	GOOS=windows go build -o bin/go-api.exe
linux:
	go build -o bin/go-api