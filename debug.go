// +build debug

package main

import "fmt"

var DebugGroup = ""

func setDebugGroup(group string) {
	DebugGroup = group
}

func debug(message interface{}) {
	fmt.Println(DebugGroup, "DEBUG:", message)
}
