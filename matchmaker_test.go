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
