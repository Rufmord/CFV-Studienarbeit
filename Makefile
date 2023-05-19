all: windows-32 windows-64 linux-32 linux-64 linux-32-arm linux-64-arm
windows-64:
	env GOOS=windows GOARCH=amd64 go build -o ./bin/CFV-Studienarbeit-64.exe
windows-32:
	env GOOS=windows GOARCH=386 go build -o ./bin/CFV-Studienarbeit-32.exe
linux-32:
	env GOOS=linux GOARCH=386 go build -o ./bin/CFV-Studienarbeit-32 main.go
linux-64:
	env GOOS=linux GOARCH=amd64 go build -o ./bin/CFV-Studienarbeit-64 main.go
linux-32-arm:
	env GOOS=linux GOARCH=arm go build -o ./bin/CFV-Studienarbeit-32-arm main.go
linux-64-arm:
	env GOOS=linux GOARCH=arm64 go build -o ./bin/CFV-Studienarbeit-64-arm main.go
clear:
	rm -rf ./bin/CFV-Studienarbeit*