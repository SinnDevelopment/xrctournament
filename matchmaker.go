package main

import "time"

// Match is the struct of the MatchMaker CSV file. - TODO - Create Match Struct from MatchMaker CSV Output
type Match struct {
}

// ScheduleEntry is a Match paired with a time.
type ScheduleEntry struct {
	Time  time.Time
	Match Match
}

// Schedule is a series of Schedule Entries attached to a type (qual, playoff, practice)
type Schedule struct {
	Matches []ScheduleEntry
	Type    string
}
