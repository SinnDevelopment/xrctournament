package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

const ()

// XRCMatchData holds the data outputted to the files for a given match.
type XRCMatchData struct {
	RedScore    int
	BlueScore   int
	RedAuto     int
	BlueAuto    int
	RedTele     int
	BlueTele    int
	RedEnd      int
	BlueEnd     int
	RedPenalty  int
	BluePenalty int
	RedAdjust   int
	BlueAdjust  int
	Timer       string
	MatchStatus string
	OPR         []XRCOPRData
}

func (m *XRCMatchData) isMatchFinished() bool {
	return m.MatchStatus == "FINISHED"
}

func (m *XRCMatchData) Equals(o XRCMatchData) bool {

	equal := m.BlueAuto == o.BlueAuto &&
		m.BluePenalty == o.BluePenalty &&
		m.BlueScore == o.BlueScore &&
		m.BlueAdjust == o.BlueAdjust &&
		m.RedAuto == o.RedAuto &&
		m.RedPenalty == o.RedPenalty &&
		m.RedScore == o.RedScore &&
		m.RedAdjust == o.RedAdjust &&
		m.Timer == m.Timer &&
		len(m.OPR) == len(o.OPR) &&
		m.MatchStatus == o.MatchStatus
	return equal
}

// XRCOPRData holds the OPR for a given player.
type XRCOPRData struct {
	Name string
	OPR  string
}

var (
	matchData     = make(chan XRCMatchData)
	previousMatch = XRCMatchData{}
)

func main() {
	// Run from within the same directory as the score files.
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})

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
			players := []XRCOPRData{}
			for _, line := range lines {
				if line == "" {
					continue
				}
				split := strings.Split(line, ":")
				players = append(players, XRCOPRData{Name: split[0], OPR: split[1]})
			}
			dataRead.OPR = players

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
			break
		case "PC_B.txt":
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
