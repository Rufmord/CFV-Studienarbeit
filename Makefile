all: windows linux
windows:
	env GOOS=windows go build -o ./bin/
linux:
	go build -o ./bin/CFV-Studienarbeit main.go
clear:
	rm -rf ./bin