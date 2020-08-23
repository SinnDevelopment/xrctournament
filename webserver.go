package main

import (
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

type WebData struct {
	TConf   TournamentConfig
	Matches []XRCMatchData
	Players []XRCPlayer
	Page    string
}

func startWebserver(port string) {
	router := gin.Default()
	router.GET("/", wIndex)
	router.GET("/matches", wMatches)
	router.GET("/rankings", wRankings)

	http.Handle("/", router)
	router.Run(":" + port)

}

func wIndex(c *gin.Context) {
	executeContent(c, "home")
}

func wMatches(c *gin.Context) {
	executeContent(c, "matches")

}
func wRankings(c *gin.Context) {
	executeContent(c, "rankings")
}

func executeContent(c *gin.Context, page string) {
	data := WebData{
		TConf:   Config,
		Matches: MATCHES,
		Players: PLAYERS,
		Page:    page,
	}
	tmpl := getPageTemplate("index.html", c)

	tmpl.Execute(c.Writer, data)
}

func getPageTemplate(page string, c *gin.Context) *template.Template {
	// reads html as a slice of bytes
	html := getData(page)
	tmpl, _ := template.New(page).Parse(html)
	return tmpl
}
