package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
}

func (m *XRCMatchData) isMatchFinished() bool {
	return m.MatchStatus == "FINISHED"
}

func (m *XRCMatchData) Equals(o XRCMatchData) bool {
	red := true
	blue := true
	for i := 0; i < 3; i++ {
		blue = blue && (m.BlueAlliance[i].Equals(o.BlueAlliance[i]))
		red = red && (m.RedAlliance[i].Equals(o.RedAlliance[i]))
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
		m.MatchStatus == o.MatchStatus &&
		red &&
		blue
	return equal
}

// XRCPlayer holds the OPR for a given player in a given match.
type XRCPlayer struct {
	Name string
	OPR  int
}

func (p *XRCPlayer) Equals(o XRCPlayer) bool {
	return p.Name == o.Name && p.OPR == o.OPR
}

func exportMatchData(data XRCMatchData) {
	export, _ := json.Marshal(data)
	ioutil.WriteFile(strconv.FormatInt(time.Now().Unix(), 10)+".json", export, 0775)
}

func readMatchData() {
	dataRead := XRCMatchData{}
	files, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {
		value, err := ioutil.ReadFile(file.Name())
		if err != nil {
			fmt.Println(err)
			return
		}

		switch file.Name() {
		case "GameState.txt":
			dataRead.MatchStatus = string(value)
			break
		case "Timer.txt":
			dataRead.Timer = string(value)
			break
		case "OPR.txt":
			lines := strings.Split(string(value), "\n")
			red := []XRCPlayer{}
			blue := []XRCPlayer{}
			for i, line := range lines {
				player := XRCPlayer{}
				if line != "" {
					split := strings.Split(line, ": ")
					opr, _ := strconv.Atoi(split[1])
					player.Name = split[0]
					player.OPR = opr
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

	matchData <- dataRead
}

func xrcDataHandler(speed int, quit chan struct{}) {
	// Run from within the same directory as the score files.
	ticker := time.NewTicker(time.Duration(speed) * time.Second)
	for {
		select {
		case received := <-matchData:
			fmt.Println("Received: ", received)
			if received.isMatchFinished() {
				// Check that we're not exporting a duplicate of the match.
				if received.Equals(previousMatch) {
					continue
				}
				previousMatch = received
				go exportMatchData(received)

				// Add to matches.json
				// Add to players.json
			}
			break
		case <-ticker.C:
			go readMatchData()
			break
		case <-quit:
			ticker.Stop()
			return
		}
	}
}
