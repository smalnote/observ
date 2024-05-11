// Package build defines build information that should be injected during
// compiling, with go build --ldflags "-X xxx/build.Version="1.0"
package debug

import (
	"runtime/debug"
	"sync"
)

// Version app version.
var Version string = "n/a"

// Time build time.
var Time string = "n/a"

// Commit build git commit.
var Commit string = "n/a"

// BuildInfo build information.
type BuildInfo struct {
	Version   string `json:"version"`
	Time      string `json:"time"`
	Commit    string `json:"commit"`
	GoVersion string `json:"goVersion"`
	// Main describes the module that contains the main package for the binary.
	Main debug.Module `json:"main"`
}

var once sync.Once
var buildInfo BuildInfo

// ReadBuildInfo returns build information.
func ReadBuildInfo() BuildInfo {
	once.Do(func() {
		buildInfo.Version = Version
		buildInfo.Time = Time
		buildInfo.Commit = Commit
		if bi, ok := debug.ReadBuildInfo(); ok {
			buildInfo.GoVersion = bi.GoVersion
			buildInfo.Main = bi.Main
		}
	})
	return buildInfo
}
