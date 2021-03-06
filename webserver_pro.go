// +build pro debug

package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// wIndex is not for glass, home page handler. Could probably be converted to one function for all 3.
func wIndex(c *gin.Context) {
	executeContent(c, "index")
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
	c.JSON(http.StatusOK, PLAYERSET)
}
func matchesAPI(c *gin.Context) {
	c.JSON(http.StatusOK, MATCHES)
}
func scheduleAPI(c *gin.Context) {
	c.JSON(http.StatusOK, MasterSchedule)
}
func wSchedule(c *gin.Context) {
	executeContent(c, "schedule")
}

func wOBS(c *gin.Context) {
	executeContent(c, "obs")
}
