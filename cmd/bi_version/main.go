package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	// read buildinfo
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		// shouldn't happen
		fmt.Println("unknown error reading build-info")
	}

	fmt.Println(bi.Main.Version)
}
