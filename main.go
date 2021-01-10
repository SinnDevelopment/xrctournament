// +build pro free debug

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
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
	WebsiteURL      string      `json:"websiteUrl"`
}

// MatchConfig holds match specific configuration data.
type MatchConfig struct {
	LogfileDirectory      string `json:"logfileDirectory"`
	PlayoffSchedule       string `json:"playoffSchedule"`
	PlayoffsEnabled       bool   `json:"playoffsEnabled"`
	QualSchedule          string `json:"qualSchedule"`
	QualificationsEnabled bool   `json:"qualificationsEnabled"`
}

var (
	// DefaultConfig is the default config.
	DefaultConfig = TournamentConfig{
		CompetitionName: "xRC Tournament",
		EnableWebserver: true,
		FileReadSpeed:   5,
		MatchDataDir:    "./",
		WebsiteURL:      "localhost",
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
	// PLAYERSET holds the player master list.
	PLAYERSET = make(map[string]XRCPlayer)
	// Config is the currently active configuration
	Config TournamentConfig
	// QualSchedule is the imported qual schedule
	QualSchedule Schedule
	// PlayoffSchedule is the imported playoff scheule
	PlayoffSchedule Schedule
	// MasterSchedule is the currently active event schedule
	MasterSchedule *Schedule
)

func main() {
	_, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Could not open config.json. Using default values.")
		fmt.Println(err)
		// Write config.json out from default values.
		Config = DefaultConfig
		defaultConfigJSON, _ := json.Marshal(DefaultConfig)
		err = ioutil.WriteFile("config.json", defaultConfigJSON, 0775)
		if err != nil {
			fmt.Println("Could not write default config.json.")
			fmt.Println(err)
		}
		return
	}

	configJSON, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("Could not read config.json.")
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(configJSON, &Config)
	if err != nil {
		fmt.Println("Could not parse config.json. Please correct linting errors.")
		fmt.Println(err)
		return
	}

	quit := make(chan struct{})

	if Config.EnableWebserver {

		// Qualifications and Playoffs are only usable when in webserver mode.
		if Config.MatchConfig.QualificationsEnabled {
			QualSchedule = ImportSchedule(Config.MatchConfig.QualSchedule)
			QualSchedule.Type = "Qualification"
			MasterSchedule = &QualSchedule
		}
		if Config.MatchConfig.PlayoffsEnabled {
			PlayoffSchedule = ImportSchedule(Config.MatchConfig.PlayoffSchedule)
			PlayoffSchedule.Type = "Playoff"
			MasterSchedule = &PlayoffSchedule
		}

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

		if useMatches {
			err = json.Unmarshal(matchesJSON, &MATCHES)
			if err != nil {
				fmt.Println(err)
			}

			debug("Read in matches successfully. Parsing to find scheduled matches.")
			expected := 0
			for _, m := range MATCHES {
				matchFound, schedule := IsScheduledMatch(&m, MasterSchedule.Matches)
				if matchFound && !MasterSchedule.Matches[schedule].Completed {
					expected++
					MasterSchedule.Matches[schedule].Completed = true
					MasterSchedule.Matches[schedule].MatchData = &m
					updateWLT(m, PLAYERSET)
				}
			}

			imported := 0
			for _, m := range MasterSchedule.Matches {
				if m.MatchData != nil {
					imported++
				}
			}
			fmt.Println(MasterSchedule)
			debug("Real Scheduled matches found that were completed: " + strconv.Itoa(imported))
			debug("Potentially Matching Scheduled matches found that were completed: " + strconv.Itoa(expected))
		}
		if usePlayers {
			err = json.Unmarshal(playerJSON, &PLAYERS)
			if err != nil {
				fmt.Println(err)
			}

			debug("Reading in master player lists.")
			for _, p := range PLAYERS {
				if p.Name == "" {
					debug("Found player with empty string name. Likely a non-full match.")
					continue
				}
				if reflect.DeepEqual(PLAYERSET[p.Name], XRCPlayer{}) {
					debug("Found new player.")
					PLAYERSET[p.Name] = p
					continue
				}
				debug("Found player that already exists. " + p.Name)
				player := PLAYERSET[p.Name]
				player.Update(p)
				PLAYERSET[p.Name] = player
			}
		}

		setVersion()
		go XRCDataHandler(Config.FileReadSpeed, quit)
		startWebserver(strconv.Itoa(Config.WebserverPort))

	} else {
		// If the webserver is not enabled, we must block the main thread from exiting with the datahandler.
		XRCDataHandler(Config.FileReadSpeed, quit)
	}
}
