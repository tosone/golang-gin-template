package main

import (
	"github.com/tosone/logging"

	"github.com/tosone/golang-gin-template/pkg/cmd"
	"github.com/tosone/golang-gin-template/pkg/version"
)

// VERSION version command output msg
var VERSION = "no provided"

// BUILDSTAMP version command output msg
var BUILDSTAMP = "no provided"

// GITHASH version command output msg
var GITHASH = "no provided"

//go:generate swag init
// @title golang-gin-template API
// @version 1.0
// @description golang-gin-template API.

// @host golang-gin-template.com
// @BasePath /api
func main() {
	// set version command output
	version.Setting(VERSION, BUILDSTAMP, GITHASH)

	// init cobra commander
	if err := cmd.RootCmd.Execute(); err != nil {
		logging.Panic(err.Error())
	}
}
