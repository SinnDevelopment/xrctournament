// +build free pro debug

package main

import (
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

// WebData is the view data used for the subviews. Pointers are used to prevent data duplication.
type WebData struct {
	TConf    *TournamentConfig
	Schedule *Schedule
	Matches  []XRCMatchData
	Players  map[string]XRCPlayer
	Page     string
	Param    int
}

// startWebServer handles initialization and running of the gin webserver.
// HTTPS should be supported at some point, but running behind nginx or a similar proxy should cover that use.
func startWebserver(port string) {
	router := gin.Default()
	router.GET("/", wIndex)
	router.GET("/matches/:match", wMatches)
	router.GET("/matches", wMatches)
	router.GET("/schedule", wSchedule)
	router.GET("/rankings", wRankings)
	router.GET("/api/players", playersAPI)
	router.GET("/api/matches", matchesAPI)
	router.GET("/api/schedule", scheduleAPI)
	router.GET("/obs", wOBS)
	http.Handle("/", router)
	router.Run(":" + port)
}

// executeContent handles data display for all pages.
func executeContent(c *gin.Context, page string) {
	param := 0
	if strings.Contains(page, ":") {
		params := strings.Split(page, ":")
		page = params[0]
		param, _ = strconv.Atoi(params[1])
	}
	data := WebData{
		TConf:    &Config,
		Schedule: MasterSchedule,
		Matches:  MATCHES,
		Players:  PLAYERSET,
		Page:     page,
		Param:    param,
	}
	html := getData(page)
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
		"time": func() string {
			return time.Now().String()
		},
	}).Parse(html)
	tmpl.Execute(c.Writer, data)
}
