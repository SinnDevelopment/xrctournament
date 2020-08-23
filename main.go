package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var (
	DefaultConfig = TournamentConfig{
		CompetitionName: "xRC Tournament",
		EnableWebserver: true,
		FileReadSpeed:   5,
		MatchConfig: MatchConfig{
			LogfileDirectory:      "./",
			PlayoffsEnabled:       true,
			QualificationsEnabled: true,
			QualSchedule:          "",
			PlayoffSchedule:       "",
		},
		TwitchChannel: "SinnDevelopment",
		WebserverPort: 8080,
	}
)

type TournamentConfig struct {
	CompetitionName string      `json:"competitionName"`
	EnableWebserver bool        `json:"enableWebserver"`
	FileReadSpeed   int         `json:"fileReadSpeed"`
	MatchConfig     MatchConfig `json:"matchConfig"`
	TwitchChannel   string      `json:"twitchChannel"`
	WebserverPort   int         `json:"webserverPort"`
}
type MatchConfig struct {
	LogfileDirectory      string `json:"logfileDirectory"`
	PlayoffSchedule       string `json:"playoffSchedule"`
	PlayoffsEnabled       bool   `json:"playoffsEnabled"`
	QualSchedule          string `json:"qualSchedule"`
	QualificationsEnabled bool   `json:"qualificationsEnabled"`
}

func main() {
	config := TournamentConfig{}
	_, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Could not open config.json. Using default values.")
		fmt.Println(err)
		// Write config.json out from default values.
		config = DefaultConfig
		defaultConfigJSON, _ := json.Marshal(DefaultConfig)
		ioutil.WriteFile("config.json", defaultConfigJSON, 0775)
		return
	}

	configJSON, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("Could not read config.json.")
		fmt.Println(err)
		return
	}
	json.Unmarshal(configJSON, &config)
	quit := make(chan struct{})
	go xrcDataHandler(config.FileReadSpeed, quit)

	if config.EnableWebserver {
		// Check if data files exist, if they do, load them

		startWebserver(strconv.Itoa(config.WebserverPort))
	}
}
