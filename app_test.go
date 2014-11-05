package main

import (
	"testing"
)

func TestConvertFromPhp(t *testing.T) {
	var phpTestMap = map[string]string{
		"Y-m-d":           "2006-01-02",
		"r":               "Mon, 02 Jan 2006 15:04:05 -0700",
		"F, y M. j. ga T": "Mon, 06 Jan. 2. 4pm MST",
	}

	for phpString, goStrig := range phpTestMap {
		convertedString := ConvertFromPhp(phpString)
		if convertedString != goStrig {
			t.Errorf("Wrong format! '%v' should '%v' but it's '%v'", phpString, goStrig, convertedString)
		}
	}
}
