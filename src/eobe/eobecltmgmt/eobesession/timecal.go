package eobesession

import "time"

type TimeCalculator struct {
}

//Left Time is before right time, true. Accurate to seconds
func (tc TimeCalculator) TimeBefore(left, right time.Time) bool {
	if left.Year() < right.Year() {
		return true
	}

	if left.Year() > right.Year() {
		return false
	}

	if left.Month() < right.Month() {
		return true
	}

	if left.Month() > right.Month() {
		return false
	}

	if left.YearDay() < right.YearDay() {
		return true
	}

	if left.YearDay() > right.YearDay() {
		return false
	}

	if left.Hour() < right.Hour() {
		return true
	}

	if left.Hour() > right.Hour() {
		return false
	}
	if left.Minute() < right.Minute() {
		return true
	}

	if left.Minute() > right.Minute() {
		return false
	}
	if left.Second() < right.Second() {
		return true
	}

	if left.Second() > right.Second() {
		return false
	}
	return false
}
