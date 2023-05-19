all: windows-32 windows-64 linux
windows-64:
	env GOOS=windows GOARCH=amd64 go build -o ./bin/CFV-Studienarbeit-64.exe
windows-32:
	env GOOS=windows GOARCH=386 go build -o ./bin/CFV-Studienarbeit-32.exe
linux:
	go build -o ./bin/CFV-Studienarbeit main.go
clear:
	rm -rf ./bin/CFV-Studienarbeit*