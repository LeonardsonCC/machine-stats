package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const CPU_STAT = "/proc/stat"

func SubCpuUsageChan() chan string {
	cpuUsed := make(chan string)

	go func() {
		for {
			u := GetCpuUsage()
			cpuUsed <- u

			time.Sleep(TIMEOUT / 2) // needs a delta
		}
	}()

	return cpuUsed
}

var (
	idle     int
	lastIdle int

	sum     int
	lastSum int
)

func GetCpuUsage() string {
	dat, err := os.ReadFile(CPU_STAT)
	if err != nil {
		fmt.Println(err.Error())
	}

	data := string(dat)
	cpu := strings.Split(data, "\n")[0]
	cpuCols := strings.Split(cpu, " ")

	totalCpu := 0
	for i, col := range cpuCols {
		value, _ := strconv.Atoi(col)

		if i == 3 { // idle column
			idle = value
		}

		totalCpu += value
	}

	sum = totalCpu

	// Get the delta between two reads :)
  delta := sum - lastSum
  idleDelta := idle - lastIdle

  used := delta - idleDelta

  // usage := 100 * used / delta

  lastSum = sum
  lastIdle = idle
	return strconv.Itoa(used)
}
