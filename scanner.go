package pwn

import (
	"sync"
	"time"
	"net"
	"strconv"
)

const SLEEP_TIME = time.Millisecond * 100;	// 100 milliseconds
const MAX_PORTS = 16777214;	// this is the maximum THEORETICAL number of ports, event though

var wg1 sync.WaitGroup;

func PortScan(host string, portFrom int, portTo int, threadCount int) []int {
	threadsPer := (portTo - portFrom) / threadCount;
	curPortsScanning := 0;
	var openPorts []int
	openPorts = make([]int, 0, MAX_PORTS)
	for x := 0; x < threadsPer; x++ {
		go scanRange(host, curPortsScanning, curPortsScanning + threadsPer, &openPorts)
		curPortsScanning += threadsPer
	}
	wg1.Wait()
	return openPorts
}

// scanRange - This is used by its wrapper function, PortScan, in order
// to individually scan a port range.
func scanRange(host string, portFrom int, portTo int, resultBuff *[]int) {
	wg1.Add(1)
	defer wg1.Done()
	for x := portFrom; x < portTo; x++ {
		_, err := net.Dial("tcp", host+":"+strconv.Itoa(x))
		if err == nil {
			*resultBuff = append(*resultBuff, x)
		}
		if x % 5 == 0 {
			// for every five connections tried, try a different
			// range for SLEEP_TIME
			time.Sleep(SLEEP_TIME)
		}
	}
}
