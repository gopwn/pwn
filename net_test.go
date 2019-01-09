package pwn

import (
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// get a random port from min 1024 to max 65535
func randPort() string {
	var port int
	for port <= 1024 {
		port = rand.Intn(65535)
	}
	return strconv.Itoa(port)
}
