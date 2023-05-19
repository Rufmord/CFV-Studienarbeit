# Computer forensics and incident handling
## Comparison of different virtual machines for forensics
The focus is on the analysis of the main memory and the persistence of paused virtual machines.

This program
- runs a sleep process
- runs a sleep thread
- opens a random port on the local machine
- calls a remote port
- opens a file on the local machine
### Execute
To build the executables for Windows and Linux
```shell
make all
```
Only for Windows
```shell
make windows
```
Only for Linux
```shell
make linux
```
Check opened ports, files and processes with
```shell
./check.sh
```