package datamunging_test

import (
	"testing"

	datamunging "github.com/go-curious/code-katas/data-munging"
)

type ParseWeatherRecordTests []struct {
	input   string
	day     int
	maxTemp int
	minTemp int
}

func TestParseWeatherRecord(t *testing.T) {
	tests := ParseWeatherRecordTests{
		{
			input:   "1  88    59    74          53.8       0.00 F       280  9.6 270  17  1.6  93 23 1004.5",
			day:     1,
			maxTemp: 88,
			minTemp: 59,
		},
		{
			input:   "26  97*   64    81          70.4       0.00 H       050  5.1 200  12  4.0 107 45 1014.9",
			day:     26,
			maxTemp: 97,
			minTemp: 64,
		},
	}

	for _, test := range tests {
		result := datamunging.ParseWeatherRecord(test.input)
		if test.day != result.Day {
			t.Fail()
		}
		if test.maxTemp != result.MaxTemp {
			t.Fail()
		}
		if test.minTemp != result.MinTemp {
			t.Fail()
		}
	}

}
