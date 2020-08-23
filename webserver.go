package main

import (
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

type WebData struct {
	TConf TournamentConfig
}

func startWebserver(port string) {
	router := gin.Default()
	router.GET("/", wIndex)
	router.GET("/matches", wMatches)
	router.GET("/players", wPlayers)
	router.GET("/rankings", wRankings)

	http.Handle("/", router)
	router.Run(":" + port)

}

func wIndex(c *gin.Context) {
	data := WebData{}

	tmpl := getPageTemplate("index.html", c)

	tmpl.Execute(c.Writer, data)
}

func wMatches(c *gin.Context) {
	data := WebData{}

	tmpl := getPageTemplate("index.html", c)

	tmpl.Execute(c.Writer, data)
}
func wPlayers(c *gin.Context) {
	data := WebData{}

	tmpl := getPageTemplate("index.html", c)

	tmpl.Execute(c.Writer, data)
}
func wRankings(c *gin.Context) {
	data := WebData{}

	tmpl := getPageTemplate("index.html", c)

	tmpl.Execute(c.Writer, data)
}

func getPageTemplate(page string, c *gin.Context) *template.Template {
	// reads html as a slice of bytes
	html := getData(page)
	tmpl, _ := template.New(page).Parse(html)
	return tmpl
}
