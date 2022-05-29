package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	const longForm string = "January 2, 2006 15:04:05"
	const realyLongForm string = "Monday, January 2, 2006 15:04:05"
	const shortForm1 string = "1/2/2006 15:04:05"
	var dateString string

	switch {
	case len(date) >= len(longForm) && len(date) <= len(longForm)+2:
		t, err := time.Parse(longForm, date)
		if err != nil {
			panic(err)
		}
		dateString = t.Format("2006-01-02 15:04:05 +00")

	case len(date) > len(longForm)+2:
		t, err := time.Parse(realyLongForm, date)
		if err != nil {
			panic(err)
		}
		dateString = t.Format("2006-01-02 15:04:05 +00")

	case len(date) <= len(shortForm1)+2:
		t, err := time.Parse(shortForm1, date)
		if err != nil {
			panic(err)
		}
		dateString = t.Format("2006-01-02 15:04:05 +00")
	default:
		fmt.Println("I don't know what time is it. (ツ)")
	}
	parsed, erro := time.Parse("2006-01-02 15:04:05 +00", dateString)
	if erro != nil {
		panic(erro)
	}
	return parsed
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	return Schedule(date).Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	return Schedule(date).Hour() >= 12 && Schedule(date).Hour() < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	tm := Schedule(date)
	return fmt.Sprintf("You have an appointment on %v, %v %v, %v, at %v:%v.", tm.Weekday(), tm.Month(), tm.Day(), tm.Year(), tm.Hour(), tm.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	currentYear := time.Date(time.Now().Year(), 9, 15, 00, 00, 0, 0, time.UTC)
	return currentYear
}
