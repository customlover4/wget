package utils

import (
	"strconv"
	"sync"
)

var counter = 0
var mu = &sync.Mutex{}

func NumeredName() string {
	mu.Lock()
	defer mu.Unlock()

	n := strconv.Itoa(counter)
	counter++
	return n
}
