package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gin-gonic/gin"
)

// WebData is the view data used for the subviews. Pointers are used to prevent data duplication.
type WebData struct {
	TConf    *TournamentConfig
	Schedule *Schedule
	Matches  []XRCMatchData
	Players  map[string]XRCPlayer
	Page     string
}

// startWebServer handles initialization and running of the gin webserver.
// HTTPS should be supported at some point, but running behind nginx or a similar proxy should cover that use.
func startWebserver(port string) {
	router := gin.Default()
	router.GET("/", wIndex)
	router.GET("/matches/:match", wMatches)
	router.GET("/matches", wMatches)
	router.GET("/playoffs", wPlayoffs)
	router.GET("/quals", wQualifications)
	router.GET("/rankings", wRankings)
	router.GET("/api/players", playersAPI)
	router.GET("/api/matches", matchesAPI)
	router.GET("/api/schedule", scheduleAPI)

	http.Handle("/", router)
	router.Run(":" + port)
}

// wIndex is not for glass, home page handler. Could probably be converted to one function for all 3.
func wIndex(c *gin.Context) {
	executeContent(c, "home")
}

func wPlayoffs(c *gin.Context) {

}

func wQualifications(c *gin.Context) {

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
		match := MATCHES[num]
		c.JSON(http.StatusOK, match)
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

// executeContent handles data display for all pages.
func executeContent(c *gin.Context, page string) {
	data := WebData{
		TConf:    &Config,
		Schedule: MasterSchedule,
		Matches:  MATCHES,
		Players:  PLAYERSET,
		Page:     page,
	}
	html := getData("index.html")
	tmpl, _ := template.New(page).Funcs(template.FuncMap{
		"avgOPR": func(p XRCPlayer) int {
			if len(p.OPR) == 0 {
				return 0
			}
			sum := 0
			for _, i := range p.OPR {
				sum += i
			}
			return sum / len(p.OPR)
		},
		"rankPoints": func(p XRCPlayer) int {
			return p.Wins*2 + p.Ties
		},
	}).Parse(html)
	tmpl.Execute(c.Writer, data)
}
