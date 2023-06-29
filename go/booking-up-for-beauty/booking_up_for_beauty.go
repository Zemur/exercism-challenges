package booking

import (
    "fmt"
    "time"
)

const (
    FULL_MONTH_DAY_YEAR_24_HOUR = "January 2, 2006 15:04:05"
    DAY_FULL_MONTH_YEAR_24_HOUR = "Monday, January 2, 2006 15:04:05"
    ISO_WITH_SLASHES = "1/02/2006 15:04:05"
    ISO_WITH_SLASHES_NO_DAY_PADDING = "1/2/2006 15:04:05"
)

func ParseTime(format string, date string) time.Time {
    t, err := time.Parse(format, date)
    
    if err != nil {
        panic(err)
    } else {
    	return t
    }
}

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    return ParseTime(ISO_WITH_SLASHES, date)
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	return ParseTime(FULL_MONTH_DAY_YEAR_24_HOUR, date).Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    var appointmentHour int = ParseTime(DAY_FULL_MONTH_YEAR_24_HOUR, date).Hour()
	return appointmentHour >= 12 && appointmentHour < 17
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	var parsedDate time.Time = ParseTime(ISO_WITH_SLASHES_NO_DAY_PADDING, date)
    return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", 
                       parsedDate.Weekday(), parsedDate.Month(), parsedDate.Day(), 
                       parsedDate.Year(), parsedDate.Hour(), parsedDate.Minute())
    
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
