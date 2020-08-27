package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"
)

// ScheduleEntry is a Match paired with a time. Linked to a match data result when a match is completed.
type ScheduleEntry struct {
	Number    int
	Time      time.Time
	Red1      string
	Red2      string
	Red3      string
	Blue1     string
	Blue2     string
	Blue3     string
	Completed bool
	MatchData *XRCMatchData
}

// MatchesXRCMatch checks if the expected alliance members are the ones that were reported from the match data.
func (m *ScheduleEntry) MatchesXRCMatch(x XRCMatchData) bool {
	blue := true
	red := true
	for _, b := range x.BlueAlliance {
		if m.Blue1 == b.Name || m.Blue2 == b.Name || m.Blue3 == b.Name {
			continue
		}
		blue = false
	}
	for _, r := range x.RedAlliance {
		if m.Red1 == r.Name || m.Red2 == r.Name || m.Red3 == r.Name {
			continue
		}
		red = false
	}
	return red && blue
}

// Schedule is a series of Schedule Entries attached to a type (qual, playoff, practice).
// Supporting multiple types of matches on the view page, or in the master schedule is not implemented.
type Schedule struct {
	Matches []ScheduleEntry
	Type    string
}

// IsScheduledMatch handles checks for whether or not a match is within the given schedule.
func IsScheduledMatch(match *XRCMatchData, schedule Schedule) (bool, ScheduleEntry) {
	for _, m := range schedule.Matches {
		if m.MatchesXRCMatch(*match) {
			m.MatchData = match
			m.Completed = true
			return true, m
		}
	}
	return false, ScheduleEntry{}
}

// ImportSchedule handles conversion of the csv of matches into the ScheduleEntry and Schedule structs.
func ImportSchedule(file string) (schedule Schedule) {

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		return
	}
	rows, err := csv.NewReader(f).ReadAll()
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
	currentMatchTime := time.Now()
	for _, row := range rows {

		matchNum, _ := strconv.Atoi(row[0])
		unixTS, _ := strconv.Atoi(row[1])
		timeRaw := time.Unix(int64(unixTS), 0)

		scheduleEntry := ScheduleEntry{
			Number: matchNum + 1,
			Time:   timeRaw,
			Red1:   row[2],
			Red2:   row[3],
			Red3:   row[4],
			Blue1:  row[5],
			Blue2:  row[6],
			Blue3:  row[7],
		}
		schedule.Matches = append(schedule.Matches, scheduleEntry)
		currentMatchTime = currentMatchTime.Add(5 * time.Minute)
	}
	return
}
