package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	StartProcess("sleep", "120") // start a process that sleeps for 120 seconds
	go StartProcess("sleep")     // start a thread with sleep
	go OpenPort(0)               // open an incoming port; Port 0 uses the next free port that is available
	CallPort()                   // call thi.de:443 and never close the connection
	OpenFile("./open_me.txt")    // open a file and never close it
	<-make(chan bool)            // run endless
}

func StartProcess(cmd string, args ...string) *exec.Cmd {
	command := exec.Command(cmd, args...)
	if err := command.Start(); err != nil {
		panic("Failed to start detached process because of: " + err.Error())
	}
	return command
}

func OpenPort(port int) {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		panic("Failed to serve because of: " + err.Error())
	}

	go func() {
		err = http.Serve(listener, nil)
		if err != nil {
			panic("Failed to serve because of: " + err.Error())
		}
	}()
	fmt.Println("Using port:", listener.Addr().(*net.TCPAddr).Port)
}

func CallPort() {
	_, err := http.Get("https://thi.de/")
	if err != nil {
		panic("Could not call thi.de because of: " + err.Error())
	}
}

func OpenFile(path string) {
	_, err := os.Open(path)
	if err != nil {
		panic("Could not open file because of: " + err.Error())
	}
}
