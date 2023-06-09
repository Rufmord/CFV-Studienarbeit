package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func main() {
	if runtime.GOOS == "windows" {
		StartProcess("timeout", "120") // start a process in windows that sleeps for 120 seconds
	} else {
		StartProcess("sleep", "120") // start a process in linux that sleeps for 120 seconds
	}
	go OpenPort(0)            // open an incoming port; Port 0 uses the next free port that is available
	CallPort()                // call thi.de:443 and never close the connection
	LoadImage()               // load an image from thi.de
	OpenFile("./open_me.txt") // open a file and never close it
	<-make(chan bool)         // run endless
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

func LoadImage() {
	image, err := http.Get("https://www.thi.de/typo3conf/ext/in2template/Resources/Public/Images/Site/thi_logo_wb_RGB.svg")
	// save image in variable to write it in RAM
	if image == nil || err != nil {
		panic("Could not load image of thi.de because of: " + err.Error())
	}
}

func OpenFile(path string) {
	f, err := os.Open(path)
	data := make([]byte, 6)
	_, err = f.Read(data)
	if err != nil {
		panic("Could not open file because of: " + err.Error())
	}
}
