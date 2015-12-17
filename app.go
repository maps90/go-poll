package main

import (
	"fmt"
	"github.com/maps90/go-poll/routes"
	"runtime"
)

func main() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	fmt.Printf("Running with %d CPUs\n", numCPU)

	routes.PourGin()
}
