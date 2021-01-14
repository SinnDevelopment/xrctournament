// +build free

package main

import "github.com/gin-gonic/gin"

func setVersion() {
	Config.CompetitionName = DefaultConfig.CompetitionName
	gin.SetMode(gin.ReleaseMode)
}
