package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {

	sched, err := time.Parse("1/02/2006 15:04:05", date)
	if err != nil {
		fmt.Println(err)
	}
	return sched
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	sched, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		fmt.Println(err)
	}
	if sched.Before(time.Now()) {
		return true
	}

	return false
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	sched, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		fmt.Println(err)
	}

	if sched.Hour() >= 12 && sched.Hour() < 18 {
		return true
	}

	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {

	sched, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("You have an appointment on %v", sched.Format("Monday, January 2, 2006, at 15:04."))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	year := fmt.Sprint(time.Now().Year())
	dateString := year + "-09-15"
	annDate, _ := time.Parse("2006-1-2", dateString)
	return annDate
}
