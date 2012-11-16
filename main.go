package main

import (
	"./queens"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

var fieldsize = flag.Int("size", 8, "size of board to solve")
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {

	initialize()
	defer uninitialize()

	solution := queens.Solve(*fieldsize)
	if solution != nil {
		fmt.Printf("Solution found: %v\n", solution)
	}
}

func initialize() {

	flag.Parse()

	if cpuProfileFlagEnabled() {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
	}
}

func cpuProfileFlagEnabled() bool {
	return cpuProfileFlag() != ""
}

func cpuProfileFlag() string {
	return *cpuprofile
}

func uninitialize() {

	if cpuProfileFlagEnabled() {
		pprof.StopCPUProfile()
	}
}
