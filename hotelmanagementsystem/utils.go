package main

import "time"

func normalizeDate(t time.Time) time.Time {
	yr, m, d := t.Date()
	return time.Date(yr, m, d, 0, 0, 0, 0, time.UTC)
}

func daysBetween(start, end time.Time) int {
	start = normalizeDate(start)
	end = normalizeDate(end)
	return int(end.Sub(start).Hours() / 24)
}

func computeTotalPrice(basePrice float64, start, end time.Time) float64 {
	n := daysBetween(start, end)
	if n <= 0 {
		return 0
	}
	return basePrice * float64(n)
}