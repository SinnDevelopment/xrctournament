package main

import "time"

// TODO - Create Match Struct from MatchMaker CSV Output
type Match struct {
}

type ScheduleEntry struct {
	Time  time.Time
	Match Match
}

type Schedule struct {
	Matches []ScheduleEntry
	Type    string
}
