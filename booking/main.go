package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
// Input format: "7/25/2019 13:45:00"
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}
	}
	return t
}

// HasPassed returns whether a date has passed.
// Input format: "January 2, 2006 15:04:05"
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		return false
	}
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
// Input format: "Monday, January 2, 2006 15:04:05"
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		return false
	}
	h := t.Hour()
	return h >= 12 && h < 18
}

// Description returns a formatted string of the appointment time.
// Output: "You have an appointment on Thursday, July 25, 2019, at 13:45."
func Description(date string) string {
	layout := "1/2/2006 15:04:05" // same format as Schedule input
	t, err := time.Parse(layout, date)
	if err != nil {
		return ""
	}
	return "You have an appointment on " + t.Format("Monday, January 2, 2006, at 15:04") + "."
}

// AnniversaryDate returns a Time with this year's anniversary.
// (Exercism expects Sep 15 at midnight UTC.)
func AnniversaryDate() time.Time {
	y := time.Now().Year()
	return time.Date(y, time.September, 15, 0, 0, 0, 0, time.UTC)
}
