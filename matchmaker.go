package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"
)

// ScheduleEntry is a Match paired with a time.
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

// MatchesXRCMatch checks if the alliances are correct. Order _technically_ doesn't matter.
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

// Schedule is a series of Schedule Entries attached to a type (qual, playoff, practice)
type Schedule struct {
	Matches []ScheduleEntry
	Type    string
}

func isScheduledMatch(match XRCMatchData, schedule Schedule) (bool, ScheduleEntry) {
	for _, m := range schedule.Matches {
		if m.MatchesXRCMatch(match) {
			return true, m
		}
	}
	return false, ScheduleEntry{}
}

func importSchedule(file string) (schedule Schedule) {

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
		timeRaw, err := time.Parse(time.UnixDate, row[1])
		if err != nil {
			timeRaw = currentMatchTime
		}

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
