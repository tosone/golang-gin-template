package version

import (
	"fmt"
	"runtime"

	"github.com/spf13/viper"
)

// VERSION version
var VERSION = "no provided"

// BUILDSTAMP BuildStamp
var BUILDSTAMP = "no provided"

// GITHASH GitHash
var GITHASH = "no provided"

// Setting get the version command output msg
func Setting(version, buildStamp, gitHash string) {
	VERSION = version
	BUILDSTAMP = buildStamp
	GITHASH = gitHash
}

// Initialize version command entry
func Initialize() {
	fmt.Printf("%s %s/%s %s\n", viper.GetString("AppName"), runtime.GOOS, runtime.GOARCH, runtime.Version())
	fmt.Printf("Version: %s\n", VERSION)
	fmt.Printf("BuildDate: %s\n", BUILDSTAMP)
	fmt.Printf("BuildHash: %s\n", GITHASH)
}
