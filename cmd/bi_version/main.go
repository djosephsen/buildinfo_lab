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

	fmt.Printf("main version: %s\n", bi.Main.Version)
	for _, s := range bi.Settings {
		fmt.Printf("%16s %s\n", s.Key, s.Value)
	}
}
