package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	cpu1  string
	cpu2  string
	ready []string
	io1   []string
	io2   []string
	io3   []string
	io4   []string
)

func initx() {
	cpu1 = ""
	cpu2 = ""
	ready = make([]string, 10)
	io1 = make([]string, 10)
	io2 = make([]string, 10)
	io3 = make([]string, 10)
	io4 = make([]string, 10)
}

func showProcess() {
	fmt.Printf("\n****Dev By CHOKCHAI JAMNOI****\n")
	fmt.Printf("CPU[1]>%s \n", cpu1)
	fmt.Printf("CPU[2]>%s \n", cpu2)
	fmt.Printf("Ready>")
	for i := range ready {
		fmt.Printf("%s|", ready[i])
	}
	fmt.Println()
	fmt.Printf("I/O[1]>")
	for i := range io1 {
		fmt.Printf("%s|", io1[i])
	}
	fmt.Println()
	fmt.Printf("I/O[2]>")
	for i := range io2 {
		fmt.Printf("%s|", io2[i])
	}
	fmt.Println()
	fmt.Printf("I/O[3]>")
	for i := range io3 {
		fmt.Printf("%s|", io3[i])
	}
	fmt.Println()
	fmt.Printf("I/O[4]>")
	for i := range io4 {
		fmt.Printf("%s|", io4[i])
	}
	fmt.Println()
	fmt.Printf("command>")
}

func getCommand() string {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.Trim(data, "\n")
	return data
}

func insertQ(q []string, data string) {
	for i := range q {
		if q[i] == "" {
			q[i] = data
			break
		}
	}
}

func newProcess(p string) {
	if cpu1 == "" {
		cpu1 = p
	} else if cpu2 == "" {
		cpu2 = p
	} else {
		insertQ(ready, p)
	}
}

func terminate(cpuName string) {
	if cpuName == "cpu1" {
		cpu1 = deleteQ(ready)
	} else if cpuName == "cpu2" {
		cpu2 = deleteQ(ready)
	}
}
func deleteQ(q []string) string {
	result := q[0]
	for i := range q {
		if i == 0 {
			continue
		}
		q[i-1] = q[i]
	}
	q[9] = ""
	return result
}

func expire(cpuName string) {
	nextQ := deleteQ(ready)

	if nextQ == "" {
		return
	}

	if cpuName == "cpu1" {
		insertQ(ready, cpu1)
		cpu1 = nextQ
	} else if cpuName == "cpu2" {
		insertQ(ready, cpu2)
		cpu2 = nextQ
	}
}

func use_ioS(ioName string, cpuName string) {
	switch ioName {
	case "1":
		io_cpu(io1, cpuName)
	case "2":
		io_cpu(io2, cpuName)
	case "3":
		io_cpu(io3, cpuName)
	case "4":
		io_cpu(io4, cpuName)
	default:
		return
	}
}

func io_cpu(io []string, cpu string) {
	if cpu == "cpu1" {
		insertQ(io, cpu1)
		cpu1 = ""
	} else if cpu == "cpu2" {
		insertQ(io, cpu2)
		cpu2 = ""
	}
	expire(cpu)
}

func use_ioSx(ioName string) {
	fq := ""
	switch ioName {
	case "1":
		fq = deleteQ(io1)
	case "2":
		fq = deleteQ(io2)
	case "3":
		fq = deleteQ(io3)
	case "4":
		fq = deleteQ(io4)
	default:
		return
	}
	if fq == "" {
		return
	}
	if cpu1 == "" {
		cpu1 = fq
	} else if cpu2 == "" {
		cpu2 = fq
	} else {
		insertQ(ready, fq)
	}
}

func main() {
	initx()
	for {
		showProcess()
		command := getCommand()
		commandx := strings.Split(command, " ")
		switch commandx[0] {
		case "exit":
			return
		case "new":
			for i := range commandx {
				if i == 0 {
					continue
				}
				newProcess(commandx[i])
			}
		case "terminate":
			for i := range commandx {
				if i == 0 {
					continue
				}
				terminate(commandx[i])
			}
		case "expire":
			for i := range commandx {
				if i == 0 {
					continue
				}
				expire(commandx[i])
			}
		case "io":
			use_ioS(commandx[1], commandx[2])

		case "iox":
			use_ioSx(commandx[1])

		default:
			fmt.Printf("\nERROR!! PLEASE TRY AGAIN...\n")
		}
	}

}
