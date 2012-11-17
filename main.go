package main

import (
	"./queens"
	"flag"
	"log"
	"os"
	"runtime/pprof"
)

var size = flag.Int("size", 8, "size of board to solve")
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {

	initialize()
	defer uninitialize()

	queens.Solve(*size)
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
	return *cpuprofile != ""
}

func uninitialize() {

	if cpuProfileFlagEnabled() {
		pprof.StopCPUProfile()
	}
}
