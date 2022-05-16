package main

import (
	"fmt"
	"runtime/debug"
	"time"

	"golang.org/x/mod/module"
)

var version = "(devel)"

func main() {
	// read buildinfo
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		// shouldn't happen
		fmt.Println("unknown error reading build-info")
	}

	// prefer classic ldflags `version` overrides
	if version == "(devel)" {
		// no ldflags version provided, use buildinfo
		version = bi.Main.Version
	}

	// no specific version provided by ldflags OR buildinfo; construct a new one
	if version == "(devel)" {
		var vcsTime time.Time
		var vcsRevision string
		for _, s := range bi.Settings {
			switch s.Key {
			case "vcs.time":
				// If the format is invalid, we'll print a zero timestamp.
				vcsTime, _ = time.Parse(time.RFC3339Nano, s.Value)
			case "vcs.revision":
				vcsRevision = s.Value
				if len(vcsRevision) > 12 {
					vcsRevision = vcsRevision[:12]
				}
			}
		}
		if vcsRevision != "" {
			version = module.PseudoVersion("", "", vcsTime, vcsRevision)
		}
	}

	fmt.Printf("%s\n", version)
}
