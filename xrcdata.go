// +build pro free

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	matchData     = make(chan XRCMatchData)
	previousMatch = XRCMatchData{}
)

// XRCMatchData holds the data outputted to the files for a given match.
type XRCMatchData struct {
	RedScore         int
	BlueScore        int
	RedAuto          int
	BlueAuto         int
	RedTele          int
	BlueTele         int
	RedEnd           int
	BlueEnd          int
	RedPenalty       int
	RedPenaltyCount  int
	BluePenalty      int
	BluePenaltyCount int
	RedAdjust        int
	BlueAdjust       int
	Timer            string
	MatchStatus      string
	RedAlliance      []XRCPlayer
	BlueAlliance     []XRCPlayer
	Completed        time.Time
}

func (m *XRCMatchData) isMatchFinished() bool {
	return m.MatchStatus == "FINISHED"
}

// Equals replaces reflect.DeepEquals
func (m *XRCMatchData) Equals(o XRCMatchData) bool {
	red := true
	blue := true
	if len(m.BlueAlliance) == len(o.BlueAlliance) && len(o.RedAlliance) == len(m.RedAlliance) {
		for i := 0; i < len(m.RedAlliance); i++ {
			red = red && (m.RedAlliance[i].Equals(o.RedAlliance[i]))
		}
		for i := 0; i < len(m.BlueAlliance); i++ {
			blue = blue && (m.BlueAlliance[i].Equals(o.BlueAlliance[i]))
		}

	}
	equal := m.BlueAuto == o.BlueAuto &&
		m.BluePenalty == o.BluePenalty &&
		m.BlueScore == o.BlueScore &&
		m.BlueAdjust == o.BlueAdjust &&
		m.RedAuto == o.RedAuto &&
		m.RedPenalty == o.RedPenalty &&
		m.RedScore == o.RedScore &&
		m.RedAdjust == o.RedAdjust &&
		m.Timer == m.Timer &&
		m.MatchStatus == o.MatchStatus
		//&& red && blue
	return equal
}

// XRCPlayer holds the data for a given player in a given match.
type XRCPlayer struct {
	Name   string
	OPR    []int
	Wins   int
	Losses int
	Ties   int
}

func (p *XRCPlayer) AvgOPR() int {
	if len(p.OPR) == 0 {
		return 0
	}
	sum := 0
	for _, i := range p.OPR {
		sum += i
	}
	return sum / len(p.OPR)
}

func (p *XRCPlayer) RP() int {
	return p.Wins*2 + p.Ties
}

func (p *XRCPlayer) Update(o XRCPlayer) {
	if p.Name != o.Name {
		return
	}
	if len(p.OPR) == len(o.OPR) {
		// Player OPRs are the same, do not combine.
		return
	}

	p.Wins += o.Wins
	p.Ties += o.Ties
	p.Losses += o.Losses
	p.OPR = append(p.OPR, o.OPR...)
}

// Equals replaces deep reflection
func (p *XRCPlayer) Equals(o XRCPlayer) bool {
	return p.Name == o.Name
}

// exportMatchData writes out the per-match log files to the same directory as the
// executable is being run from.
func exportMatchData(data XRCMatchData) {
	export, _ := json.Marshal(data)
	path := filepath.FromSlash(Config.MatchConfig.LogfileDirectory + "/" + strconv.FormatInt(time.Now().Unix(), 10) + ".json")
	err := ioutil.WriteFile(path, export, 0775)
	if err != nil {
		fmt.Println("Could not write match archive data.")
		fmt.Println(err)
	}
}

// exportPlayers writes to disk the contents of the passed player list.
func exportPlayers(match XRCMatchData, seenPlayers *[]XRCPlayer, playerSet map[string]XRCPlayer) {
	// Need to parse and deduplicate.

	matchPlayers := append(match.BlueAlliance, match.RedAlliance...)
	*seenPlayers = append(*seenPlayers, matchPlayers...)

	for _, p := range matchPlayers {
		if p.Name == "" {
			continue
		}
		if reflect.DeepEqual(playerSet[p.Name], XRCPlayer{}) {
			playerSet[p.Name] = p
			continue
		}
		record := playerSet[p.Name]
		record.Update(p)
		playerSet[p.Name] = record
	}

	export, _ := json.Marshal(*seenPlayers)
	err := ioutil.WriteFile("players.json", export, 0775)
	if err != nil {
		fmt.Println("Could not write player master data.")
		fmt.Println(err)
	}

}

// exportMatches writes to disk the contents of the recorded matches.
func exportMatches(match XRCMatchData, matches *[]XRCMatchData) {
	*matches = append(*matches, match)
	export, _ := json.Marshal(*matches)
	err := ioutil.WriteFile("matches.json", export, 0775)
	if err != nil {
		fmt.Println("Could not write match master data.")
		fmt.Println(err)
	}

}

// readMatchData handles the main file read loop, getting all the data from the match files at the specified polling rate in the config.
func readMatchData(dataChannel chan XRCMatchData) {
	dataRead := XRCMatchData{}
	files, err := ioutil.ReadDir(Config.MatchDataDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".txt") {
			continue
		}
		path := filepath.FromSlash(Config.MatchDataDir + "/" + file.Name())
		value, err := ioutil.ReadFile(path)

		if err != nil {
			fmt.Println(err)
			return
		}
		// Parse game data files only
		switch file.Name() {
		case "GameState.txt":
			dataRead.MatchStatus = string(value)
			break
		case "Timer.txt":
			dataRead.Timer = string(value)
			break
		case "OPR.txt":
			// OPR.txt has a special format, each line is related to the alliance position
			// while blank lines mean no player was in that slot.
			// This behaviour needs to be checked for alliance sizes not equal to 3. ie. FTC.
			lines := strings.Split(string(value), "\n")
			var red []XRCPlayer
			var blue []XRCPlayer
			for i, line := range lines {
				player := XRCPlayer{}
				if line != "" {
					split := strings.Split(line, ": ")
					opr, _ := strconv.Atoi(split[1])
					player.Name = split[0]
					player.OPR = append(player.OPR, opr)
				}
				if i < 3 {
					red = append(red, player)
				} else if i >= 3 {
					blue = append(blue, player)
				}
			}
			dataRead.RedAlliance = red
			dataRead.BlueAlliance = blue
			break
		case "AutoR.txt":
			dataRead.RedAuto, _ = strconv.Atoi(string(value))
			break
		case "AutoB.txt":
			dataRead.BlueAuto, _ = strconv.Atoi(string(value))
			break
		case "TeleR.txt":
			dataRead.RedTele, _ = strconv.Atoi(string(value))
			break
		case "TeleB.txt":
			dataRead.BlueTele, _ = strconv.Atoi(string(value))
			break
		case "EndR.txt":
			dataRead.RedEnd, _ = strconv.Atoi(string(value))
			break
		case "EndB.txt":
			dataRead.BlueEnd, _ = strconv.Atoi(string(value))
			break
		case "PC_R.txt":
			dataRead.RedPenaltyCount, _ = strconv.Atoi(string(value))
			break
		case "PC_B.txt":
			dataRead.BluePenaltyCount, _ = strconv.Atoi(string(value))
			break
		case "PenR.txt":
			dataRead.RedPenalty, _ = strconv.Atoi(string(value))
			break
		case "PenB.txt":
			dataRead.BluePenalty, _ = strconv.Atoi(string(value))
			break
		case "ScoreR.txt":
			dataRead.RedScore, _ = strconv.Atoi(string(value))
			break
		case "ScoreB.txt":
			dataRead.BlueScore, _ = strconv.Atoi(string(value))
			break
		case "RedADJ.txt":
			dataRead.RedAdjust, _ = strconv.Atoi(string(value))
			break
		case "BlueADJ.txt":
			dataRead.BlueAdjust, _ = strconv.Atoi(string(value))
			break
		}
	}
	if dataRead.RedScore > dataRead.BlueScore {
		for _, p := range dataRead.RedAlliance {
			p.Wins += 1
		}
		for _, p := range dataRead.BlueAlliance {
			p.Losses += 1
		}
	} else if dataRead.BlueScore > dataRead.RedScore {
		for _, p := range dataRead.BlueAlliance {
			p.Wins += 1
		}
		for _, p := range dataRead.RedAlliance {
			p.Losses += 1
		}
	} else {
		for _, p := range append(dataRead.RedAlliance, dataRead.BlueAlliance...) {
			p.Ties += 1
		}
	}
	dataRead.Completed = time.Now()
	dataChannel <- dataRead
}

// XRCDataHandler is the main loop for reading new values and updating the existing ones.
// Whenever possible, the functions called are split into goroutines for concurrency.
func XRCDataHandler(speed int, quit chan struct{}) {

	ticker := time.NewTicker(time.Duration(speed) * time.Second)
	for {
		select {
		case received := <-matchData:
			if received.isMatchFinished() {
				// Check that we're not exporting a duplicate of the match.
				if received.Equals(previousMatch) {
					continue
				}
				fmt.Println("Received: ", received)
				previousMatch = received

				go IsScheduledMatch(&received, MasterSchedule)
				go exportMatchData(received)
				go exportPlayers(received, &PLAYERS, PLAYERSET)
				go exportMatches(received, &MATCHES)
			}
			break
		case <-ticker.C:
			go readMatchData(matchData)
			break
		case <-quit:
			ticker.Stop()
			return
		}
	}
}
