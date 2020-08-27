package main

import (
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

// WebData is the view data used for the subviews. Pointers are used to prevent data duplication.
type WebData struct {
	TConf    *TournamentConfig
	Schedule *Schedule
	Matches  *[]XRCMatchData
	Players  *[]XRCPlayer
	Page     string
}

// startWebServer handles initialization and running of the gin webserver.
// HTTPS should be supported at some point, but running behind nginx or a similar proxy should cover that use.
func startWebserver(port string) {
	router := gin.Default()
	router.GET("/", wIndex)
	router.GET("/matches", wMatches)
	router.GET("/rankings", wRankings)

	http.Handle("/", router)
	router.Run(":" + port)
}

// wIndex is not for glass, home page handler. Could probably be converted to one function for all 3.
func wIndex(c *gin.Context) {
	executeContent(c, "home")
}

// wMatches handles match view.
func wMatches(c *gin.Context) {
	executeContent(c, "matches")
}

// wRankings handles rankings view.
func wRankings(c *gin.Context) {
	executeContent(c, "rankings")
}

// executeContent handles data display for all pages.
func executeContent(c *gin.Context, page string) {
	data := WebData{
		TConf:    &Config,
		Schedule: &MasterSchedule,
		Matches:  &MATCHES,
		Players:  &PLAYERS,
		Page:     page,
	}
	html := getData(page)
	tmpl, _ := template.New(page).Parse(html)
	tmpl.Execute(c.Writer, data)
}
