package util

import "testing"

func TestStringToTime(t *testing.T) {
	time := StringToTime("2024-05-06T01:44:36")

	if time.Year() != 2024 {
		t.Errorf("Year is not correct")
	}

	if time.Month() != 5 {
		t.Errorf("Month is not correct")
	}

	if time.Day() != 6 {
		t.Errorf("Day is not correct")
	}

	if time.Hour() != 1 {
		t.Errorf("Hour is not correct")
	}

	if time.Minute() != 44 {
		t.Errorf("Minute is not correct")
	}

	if time.Second() != 36 {
		t.Errorf("Second is not correct")
	}
}
