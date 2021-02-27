package main

import "testing"

func TestTwoTimes(t *testing.T) {
	if timesTwo(4) != 15 {
		t.Error("timesTwo(4) is not 15")
	}
}
