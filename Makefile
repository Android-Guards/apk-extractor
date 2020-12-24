build_all:
	GOOS=windows GOARCH=386 go build -o apk-extractor-windows-386.exe
	GOOS=windows GOARCH=amd64 go build -o apk-extractor-windows-amd64.exe
	GOOS=linux GOARCH=386 go build -o apk-extractor-linux-386
	GOOS=linux GOARCH=amd64 go build -o apk-extractor-linux-amd64
	GOOS=darwin GOARCH=amd64 go build -o apk-extractor-macos-amd64
