package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

const MEMINFO_FILE = "/proc/meminfo"

func SubMemUsageChan() chan string {
	memUsed := make(chan string)

	go func() {
		for {
			u := GetMemoryUsage()
			memUsed <- u

			time.Sleep(TIMEOUT)
		}
	}()

	return memUsed
}

func GetMemoryUsage() string {
	var (
		// Memory Total
		memtotal float32

		// Memory Free
		memfree float32

		// Memory Cached
		memcached float32

		// Memory Buffers
		membuffers float32

		memused float32
	)

	dat, err := os.ReadFile(MEMINFO_FILE)
	if err != nil {
		fmt.Println(err.Error())
	}

	data := string(dat)
	for _, line := range strings.Split(data, "\n") {
		spl := strings.Split(line, ":")

		switch spl[0] {
		case "MemTotal":
			memtotal = FromMeminfoToNumber(spl[1])
		case "MemFree":
			memfree = FromMeminfoToNumber(spl[1])

		case "Buffers":
			membuffers = FromMeminfoToNumber(spl[1])

		case "Cached", "SReclaimable":
			memcached += FromMeminfoToNumber(spl[1])
		case "Shmem":
			memcached -= FromMeminfoToNumber(spl[1])
		}
	}
	memused = memtotal - (memfree + memcached + membuffers)

	return fmt.Sprintf("%f GB", memused)
}

func FromMeminfoToNumber(value string) float32 {
	cleanValue := strings.TrimRight(value, " kB")
	cleanValue = strings.Trim(cleanValue, " ")

	valueFloat, err := strconv.ParseFloat(cleanValue, 64)
	if err != nil {
		fmt.Println(err.Error())
	}

	return float32(valueFloat / math.Pow(1024.0, 2.0))
}
