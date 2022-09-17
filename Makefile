linux:
	go build -o bin/go-api
windows:
	GOOS=windows go build -o bin/go-api.exe