package main

import (
	"fmt"
	"os"
	"time"
)

var TIMEOUT = 1 * time.Second

func main() {
	args, err := ParseArgs()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var memUsed chan string
	if args.Memory {
		memUsed = SubMemUsageChan()
	}
	
	var cpuUsed chan string
	if args.Cpu {
		cpuUsed = SubCpuUsageChan()
	}

	for {
		strToPrint := GetStrFormat(args)

		select {
		case usage := <-memUsed:
			strToPrint = fmt.Sprintf(strToPrint, usage)
		case usage := <-cpuUsed:
			strToPrint = fmt.Sprintf(strToPrint, usage)
		}

		fmt.Printf(strToPrint)
	}
}

func GetStrFormat(args *Args) string {
	if (args.Memory || args.Cpu) && !(args.Memory && args.Cpu) {
		return  "\r%s"
	}
	return ""
}
