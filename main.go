package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// TournamentConfig is the master configuration
type TournamentConfig struct {
	CompetitionName string      `json:"competitionName"`
	EnableWebserver bool        `json:"enableWebserver"`
	FileReadSpeed   int         `json:"fileReadSpeed"`
	MatchDataDir    string      `json:"matchDataDir"`
	MatchConfig     MatchConfig `json:"matchConfig"`
	TwitchChannel   string      `json:"twitchChannel"`
	WebserverPort   int         `json:"webserverPort"`
}

// MatchConfig holds match specific configuration data.
type MatchConfig struct {
	LogfileDirectory      string `json:"logfileDirectory"`
	PlayoffSchedule       string `json:"playoffSchedule"`
	PlayoffsEnabled       bool   `json:"playoffsEnabled"`
	QualSchedule          string `json:"qualSchedule"`
	QualificationsEnabled bool   `json:"qualificationsEnabled"`
}

// MatchDataFile is a container for matches.json
type MatchDataFile struct {
	Matches []XRCMatchData
}

// PlayerDataFile is a container for players.json
type PlayerDataFile struct {
	Players []XRCPlayer
}

var (
	// DefaultConfig is the default config.
	DefaultConfig = TournamentConfig{
		CompetitionName: "xRC Tournament",
		EnableWebserver: true,
		FileReadSpeed:   5,
		MatchDataDir:    "./",
		MatchConfig: MatchConfig{
			LogfileDirectory:      "./",
			PlayoffsEnabled:       true,
			QualificationsEnabled: true,
			QualSchedule:          "schedule.csv",
			PlayoffSchedule:       "elimschedule.csv",
		},
		TwitchChannel: "SinnDevelopment",
		WebserverPort: 8080,
	}
	// MATCHES holds the current master list of matches played.
	MATCHES []XRCMatchData
	// PLAYERS holds the current master list of players seen.
	PLAYERS []XRCPlayer
	// Config is the currently active configuration
	Config TournamentConfig
	// MasterSchedule is the currently active event schedule
	MasterSchedule Schedule
)

func main() {
	_, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Could not open config.json. Using default values.")
		fmt.Println(err)
		// Write config.json out from default values.
		Config = DefaultConfig
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
	json.Unmarshal(configJSON, &Config)

	quit := make(chan struct{})

	if Config.EnableWebserver {

		usePlayers := true
		useMatches := true

		matchesJSON, err := ioutil.ReadFile("matches.json")
		if err != nil {
			fmt.Println("Could not read matches.json. Starting with no matches run.")
			useMatches = false
		}
		playerJSON, err := ioutil.ReadFile("players.json")
		if err != nil {
			fmt.Println("Could not read players.json. Starting with no players.")
			usePlayers = false
		}
		matchTemp := MatchDataFile{}
		playerTemp := PlayerDataFile{}
		if useMatches {
			json.Unmarshal(matchesJSON, &matchTemp)
		}
		if usePlayers {
			json.Unmarshal(playerJSON, &playerTemp)
		}

		MATCHES = matchTemp.Matches
		PLAYERS = playerTemp.Players

		// Qualifications and Playoffs are only usable when in webserver mode.
		if Config.MatchConfig.QualificationsEnabled {
			MasterSchedule = ImportSchedule(Config.MatchConfig.QualSchedule)
		}

		go XRCDataHandler(Config.FileReadSpeed, quit)
		startWebserver(strconv.Itoa(Config.WebserverPort))

	} else {
		// If the webserver is not enabled, we must block the main thread from exiting with the datahandler.
		XRCDataHandler(Config.FileReadSpeed, quit)
	}
}
