package main

import (
	"testing"
)

func Test_exportMatches(t *testing.T) {
	type args struct {
		match XRCMatchData
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exportMatches(tt.args.match)
		})
	}
}

func Test_readMatchData(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			readMatchData()
		})
	}
}

func Test_exportPlayers(t *testing.T) {
	type args struct {
		match XRCMatchData
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exportPlayers(tt.args.match)
		})
	}
}

func Test_checkSchedule(t *testing.T) {
	type args struct {
		data XRCMatchData
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkSchedule(tt.args.data)
		})
	}
}