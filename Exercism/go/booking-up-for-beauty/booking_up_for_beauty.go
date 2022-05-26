package booking

import (
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	t, err := time.Parse("1/02/2006 15:04:05", date)
	if err != nil {
		panic(err)
	}

	dateString := t.Format("2006-01-02 15:04:05 +00")

	parsed, erro := time.Parse("2006-01-02 15:04:05 +00", dateString)
	if erro != nil {
		panic(erro)
	}

	return parsed
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	panic("Please implement the HasPassed function")
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	panic("Please implement the IsAfternoonAppointment function")
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	panic("Please implement the Description function")
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	panic("Please implement the AnniversaryDate function")
}
