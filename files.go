// +build pro free

package main

import (
	_ "embed"
)

// Requires 1.16

var (
	//go:embed "web/index.html"
	index string
	//go:embed "web/match.html"
	match string
	//go:embed "web/matches.html"
	matches string
	//go:embed "web/rankings.html"
	rankings string
	//go:embed "web/schedule.html"
	schedule string
	//go:embed "web/template.html"
	_template string
)

func getData(name string) string {
	switch name {
	case "index":
		return _template + index
	case "match":
		return _template + match
	case "matches":
		return _template + matches
	case "rankings":
		return _template + rankings
	case "schedule":
		return _template + schedule
	case "template":
		return _template + _template
	default:
		return ""
	}
}
