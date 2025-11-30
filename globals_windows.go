//go:build windows

package main

import "github.com/tischda/timer/registry"

const REGISTRY_PATH_SOFTWARE = `SOFTWARE\Tischer`

const shell = "cmd"
const shellCmdFlag = "/c"

// using sleep command that comes with git for windows
const execTestCmd = "sleep 1"
const execTestRxp = `Total time: 1.\d*s`

var timer = &Timer{registry: registry.RealRegistry{}}
