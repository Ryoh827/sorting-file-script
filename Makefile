linux:
	GOOS=linux GOARCH=amd64 go build -o linux-amd64/sorting-file-script ./main.go
win:
	GOOS=windows GOARCH=amd64 go build -o windows-amd64/sorting-file-script.exe ./main.go
