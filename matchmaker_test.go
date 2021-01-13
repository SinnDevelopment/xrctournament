// +build pro free debug

package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestImportSchedule_3v3(t *testing.T) {
	expected := Schedule{
		Matches: []ScheduleEntry{
			{Number: 1, Red1: "Red1", Red2: "Red2", Red3: "Red3",
				Blue1: "Blue1", Blue2: "Blue2", Blue3: "Blue3", Completed: false, MatchData: nil, Time: time.Unix(0, 0)},
			{Number: 2, Red1: "Red1-2", Red2: "Red2-2", Red3: "Red3-2",
				Blue1: "Blue1-2", Blue2: "Blue2-2", Blue3: "Blue3-2", Completed: false, MatchData: nil, Time: time.Unix(1, 0)},
		},
	}

	schedule := ImportSchedule("testdata/unittest/3v3.csv")
	if !reflect.DeepEqual(expected, schedule) {
		fmt.Println("Expected schedule for 3v3 was not equal to the imported one.")
		fmt.Printf("Expected: ")
		fmt.Println(expected)
		fmt.Printf("Real: ")
		fmt.Println(schedule)
		t.Fail()
	}
}

func TestImportSchedule_2v2(t *testing.T) {
	expected := Schedule{
		Matches: []ScheduleEntry{
			{Number: 1, Red1: "Red1", Red2: "Red2",
				Blue1: "Blue1", Blue2: "Blue2", Completed: false, MatchData: nil, Time: time.Unix(0, 0)},
			{Number: 2, Red1: "Red1-2", Red2: "Red2-2",
				Blue1: "Blue1-2", Blue2: "Blue2-2", Completed: false, MatchData: nil, Time: time.Unix(1, 0)},
		},
	}

	schedule := ImportSchedule("testdata/unittest/2v2.csv")
	if !reflect.DeepEqual(expected, schedule) {
		fmt.Println("Expected schedule for 2v2 was not equal to the imported one.")
		fmt.Printf("Expected: ")
		fmt.Println(expected)
		fmt.Printf("Real: ")
		fmt.Println(schedule)
		t.Fail()
	}
}

func TestScheduleEntry_MatchesXRCMatch(t *testing.T) {
	schedule := Schedule{
		Type: "Qualification",
		Matches: []ScheduleEntry{
			{
				Number:      0,
				Time:        time.Time{},
				Red1:        "1",
				Red2:        "2",
				Red3:        "3",
				Blue1:       "4",
				Blue2:       "5",
				Blue3:       "6",
				Completed:   false,
				MatchData:   nil,
				MasterIndex: 0,
			},
		}}

	match := XRCMatchData{
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
				Name:        "1",
				OPR:         nil,
				Wins:        0,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
			{
				Name:        "2",
				OPR:         nil,
				Wins:        0,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
			{
				Name:        "3",
				OPR:         nil,
				Wins:        0,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
		},
		BlueAlliance: []XRCPlayer{
			{
				Name:        "4",
				OPR:         nil,
				Wins:        0,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
			{
				Name:        "5",
				OPR:         nil,
				Wins:        0,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
			{
				Name:        "6",
				OPR:         nil,
				Wins:        0,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
		},
		Completed: time.Now(),
	}
	if !schedule.Matches[0].MatchesXRCMatch(match) {
		t.Fail()
	}
}

func TestIsScheduledMatch(t *testing.T) {
	schedule := Schedule{
		Type: "Qualification",
		Matches: []ScheduleEntry{
			{
				Number:      0,
				Time:        time.Time{},
				Red1:        "1",
				Red2:        "2",
				Red3:        "3",
				Blue1:       "4",
				Blue2:       "5",
				Blue3:       "6",
				Completed:   false,
				MatchData:   nil,
				MasterIndex: 0,
			},
			{
				Number:      1,
				Time:        time.Time{},
				Red1:        "11",
				Red2:        "12",
				Red3:        "13",
				Blue1:       "14",
				Blue2:       "15",
				Blue3:       "16",
				Completed:   false,
				MatchData:   nil,
				MasterIndex: 0,
			},
		}}

	match := XRCMatchData{
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
				Name:        "1",
				OPR:         nil,
				Wins:        0,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
			{
				Name:        "2",
				OPR:         nil,
				Wins:        0,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
			{
				Name:        "3",
				OPR:         nil,
				Wins:        0,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
		},
		BlueAlliance: []XRCPlayer{
			{
				Name:        "4",
				OPR:         nil,
				Wins:        0,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
			{
				Name:        "5",
				OPR:         nil,
				Wins:        0,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
			{
				Name:        "6",
				OPR:         nil,
				Wins:        0,
				Losses:      0,
				Ties:        0,
				MatchesSeen: nil,
			},
		},
		Completed: time.Now(),
	}
	result, index := IsScheduledMatch(&match, schedule.Matches)
	if !result || index != 0 {
		t.Fail()
	}
}
