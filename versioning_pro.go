// +build pro

package main

// Required to support Codemeter.

//import "C"
import "github.com/gin-gonic/gin"

func setVersion() {
	gin.SetMode(gin.ReleaseMode)
}
