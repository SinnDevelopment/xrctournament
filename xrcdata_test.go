// +build free pro debug

package main

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestReadMatchesJSON(t *testing.T) {

}

func TestReadPlayersJSON(t *testing.T) {

}

func TestUpdateMatchWLT(t *testing.T) {

}

func TestXRCPlayer_Update(t *testing.T) {
	player1 := XRCPlayer{
		Name:        "Player",
		OPR:         []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		Wins:        0,
		Losses:      0,
		Ties:        0,
		MatchesSeen: nil,
	}
	player2 := XRCPlayer{
		Name:        "Player",
		OPR:         []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		Wins:        1,
		Losses:      2,
		Ties:        3,
		MatchesSeen: nil,
	}
	player1.Update(player2)

}

func TestXRCPlayer_AvgOPR(t *testing.T) {
	player := XRCPlayer{
		Name:        "Player",
		OPR:         []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		Wins:        0,
		Losses:      0,
		Ties:        0,
		MatchesSeen: nil,
	}
	if player.AvgOPR() != 5 {
		t.Fail()
	}
}

func TestXRCPlayer_RP(t *testing.T) {
	player := XRCPlayer{
		Name:        "Player",
		OPR:         []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		Wins:        1,
		Losses:      1,
		Ties:        1,
		MatchesSeen: nil,
	}

	if player.RP() != 3 {
		t.Fail()
	}
}

func TestXRCPlayer_Equals(t *testing.T) {
	player1 := XRCPlayer{
		Name:        "Player",
		OPR:         []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		Wins:        1,
		Losses:      1,
		Ties:        1,
		MatchesSeen: nil,
	}
	player2 := XRCPlayer{
		Name:        "Player",
		OPR:         []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		Wins:        1,
		Losses:      1,
		Ties:        1,
		MatchesSeen: nil,
	}
	player3 := XRCPlayer{
		Name:        "Player3",
		OPR:         []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		Wins:        1,
		Losses:      1,
		Ties:        1,
		MatchesSeen: nil,
	}

	if player1.Equals(player3) {
		t.Fail()
	}

	if !player1.Equals(player2) {
		t.Fail()
	}
}

func TestXRCMatchData_Equals(t *testing.T) {

	match1 := XRCMatchData{
		RedScore:         0,
		BlueScore:        0,
		RedAuto:          0,
		BlueAuto:         0,
		RedTele:          0,
		BlueTele:         0,
		RedEnd:           0,
		BlueEnd:          0,
		RedPenalty:       0,
		RedPenaltyCount:  0,
		BluePenalty:      0,
		BluePenaltyCount: 0,
		RedAdjust:        0,
		BlueAdjust:       0,
		Timer:            "",
		MatchStatus:      "",
		RedAlliance: []XRCPlayer{
			{
				Name:        "Red1",
				OPR:         nil,
				Wins:        1,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
			{
				Name:        "Red2",
				OPR:         nil,
				Wins:        0,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
			{
				Name:        "Red3",
				OPR:         nil,
				Wins:        0,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
		},
		BlueAlliance: []XRCPlayer{
			{
				Name:        "Blue1",
				OPR:         nil,
				Wins:        1,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
			{
				Name:        "Blue2",
				OPR:         nil,
				Wins:        0,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
			{
				Name:        "Blue3",
				OPR:         nil,
				Wins:        0,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
		},
		Completed: time.Time{},
	}
	match2 := XRCMatchData{
		RedScore:         0,
		BlueScore:        0,
		RedAuto:          0,
		BlueAuto:         0,
		RedTele:          0,
		BlueTele:         0,
		RedEnd:           0,
		BlueEnd:          0,
		RedPenalty:       0,
		RedPenaltyCount:  0,
		BluePenalty:      0,
		BluePenaltyCount: 0,
		RedAdjust:        0,
		BlueAdjust:       0,
		Timer:            "",
		MatchStatus:      "",
		RedAlliance: []XRCPlayer{
			{
				Name:        "Red1",
				OPR:         nil,
				Wins:        1,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
			{
				Name:        "Red2",
				OPR:         nil,
				Wins:        0,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
			{
				Name:        "Red3",
				OPR:         nil,
				Wins:        0,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
		},
		BlueAlliance: []XRCPlayer{
			{
				Name:        "Blue1",
				OPR:         nil,
				Wins:        1,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
			{
				Name:        "Blue2",
				OPR:         nil,
				Wins:        0,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
			{
				Name:        "Blue3",
				OPR:         nil,
				Wins:        0,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
		},
		Completed: time.Time{},
	}
	match3 := XRCMatchData{
		RedScore:         0,
		BlueScore:        0,
		RedAuto:          0,
		BlueAuto:         0,
		RedTele:          0,
		BlueTele:         0,
		RedEnd:           1,
		BlueEnd:          0,
		RedPenalty:       0,
		RedPenaltyCount:  0,
		BluePenalty:      0,
		BluePenaltyCount: 0,
		RedAdjust:        0,
		BlueAdjust:       0,
		Timer:            "",
		MatchStatus:      "",
		RedAlliance: []XRCPlayer{
			{
				Name:        "Red1",
				OPR:         nil,
				Wins:        1,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
			{
				Name:        "Red2",
				OPR:         nil,
				Wins:        0,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
			{
				Name:        "Red3",
				OPR:         nil,
				Wins:        0,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
		},
		BlueAlliance: []XRCPlayer{
			{
				Name:        "Blue1",
				OPR:         nil,
				Wins:        1,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
			{
				Name:        "Blue2",
				OPR:         nil,
				Wins:        0,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
			{
				Name:        "Blue3",
				OPR:         nil,
				Wins:        0,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
		},
		Completed: time.Time{},
	}

	if !match1.Equals(match2) {
		debug("Match 1 does not equal Match 2 - FAIL")
		t.Fail()
	}

	if match2.Equals(match3) {
		debug("Match 2 equals Match 3 - FAIL")
		t.Fail()
	}

}

func TestXRCMatchData_WriteMatchArchive(t *testing.T) {
	match := XRCMatchData{
		RedScore:         100,
		BlueScore:        200,
		RedAuto:          1,
		BlueAuto:         2,
		RedTele:          3,
		BlueTele:         4,
		RedEnd:           5,
		BlueEnd:          6,
		RedPenalty:       7,
		RedPenaltyCount:  8,
		BluePenalty:      9,
		BluePenaltyCount: 10,
		RedAdjust:        11,
		BlueAdjust:       12,
		Timer:            "0:00",
		MatchStatus:      "FINISHED",
		RedAlliance: []XRCPlayer{
			{
				Name:        "Red1",
				OPR:         nil,
				Wins:        1,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
			{
				Name:        "Red2",
				OPR:         nil,
				Wins:        0,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
			{
				Name:        "Red3",
				OPR:         nil,
				Wins:        0,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
		},
		BlueAlliance: []XRCPlayer{
			{
				Name:        "Blue1",
				OPR:         nil,
				Wins:        1,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
			{
				Name:        "Blue2",
				OPR:         nil,
				Wins:        0,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
			{
				Name:        "Blue3",
				OPR:         nil,
				Wins:        0,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
		},
		Completed: time.Now(),
	}

	path := match.WriteMatchArchive("testdata")
	file, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fail()
	}
	filecontents := string(file)
	debug("Test Match Data Output: " + filecontents)
	err = os.Remove(path)
	if err != nil {
		t.Fail()
	}
}

func TestXRCMatchData_Summary(t *testing.T) {

}
