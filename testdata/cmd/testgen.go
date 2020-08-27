package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	genFRCSchedule(50)
	genFTCSchedule(50)
}

// 3v3 matches.
func genFRCSchedule(matches int) {
	for i := 0; i < matches; i++ {
		fmt.Printf("%d,%d,%s,%s,%s,%s,%s,%s\n", i, time.Now().Unix(),
			"Team"+strconv.Itoa(i), "Team"+strconv.Itoa(i+1), "Team"+strconv.Itoa(i+2),
			"Team"+strconv.Itoa(i+3), "Team"+strconv.Itoa(i+4), "Team"+strconv.Itoa(i+5))
	}
}

// 2v2 matches, using positions 1/2.
func genFTCSchedule(matches int) {
	for i := 0; i < matches; i++ {
		fmt.Printf("%d,%d,%s,%s,,%s,%s,\n", i, time.Now().Unix(),
			"Team"+strconv.Itoa(i), "Team"+strconv.Itoa(i+1),
			"Team"+strconv.Itoa(i+2), "Team"+strconv.Itoa(i+3))
	}
}
