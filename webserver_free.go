// +build free

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// wIndex is not for glass, home page handler.
func wIndex(c *gin.Context) {
	executeContent(c, "index")
}

func wSchedule(c *gin.Context) {
	executeContent(c, "schedule")
}

// wMatches handles match view.
func wMatches(c *gin.Context) {
	if c.Param("match") != "" {
		// TODO
		matchNum := c.Param("match")
		num, err := strconv.Atoi(matchNum)
		if err != nil || num < 0 || num >= len(MATCHES) {
			fmt.Println(err)
			executeContent(c, "matches")
			return
		}
		executeContent(c, "match:"+strconv.Itoa(num))
	} else {
		executeContent(c, "matches")
	}
}

// wRankings handles rankings view.
func wRankings(c *gin.Context) {
	executeContent(c, "rankings")
}

func playersAPI(c *gin.Context) {
	c.Status(http.StatusUnauthorized)

}
func matchesAPI(c *gin.Context) {
	c.Status(http.StatusUnauthorized)

}
func scheduleAPI(c *gin.Context) {
	c.Status(http.StatusUnauthorized)
}

func wOBS(c *gin.Context) {
	c.Status(http.StatusUnauthorized)
}
