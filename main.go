package main

import (
	"github.com/tosone/logging"

	"github.com/tosone/golang-gin-template/pkg/cmd"
	"github.com/tosone/golang-gin-template/pkg/version"
)

// Version version command output msg
var Version = "no provided"

// BuildStamp version command output msg
var BuildStamp = "no provided"

// GitHash version command output msg
var GitHash = "no provided"

//go:generate swag init
// @title golang-gin-template API
// @version 1.0
// @description golang-gin-template API.

// @host golang-gin-template.com
// @BasePath /api
func main() {
	// set version command output
	version.Setting(Version, BuildStamp, GitHash)

	// init cobra commander
	if err := cmd.RootCmd.Execute(); err != nil {
		logging.Panic(err.Error())
	}
}
