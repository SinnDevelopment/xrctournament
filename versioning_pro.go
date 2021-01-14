// +build pro

package main

//// Required to support Codemeter.
import "C"

func setVersion() {
	gin.SetMode(gin.ReleaseMode)
}
