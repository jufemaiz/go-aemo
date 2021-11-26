package nem12

import "time"

const (
	// nanosInSecond is the number of nanoseconds in a second.
	nanosInSecond = 1e9
	// secondsInMinute is the number of seconds in a minute.
	secondsInMinute = 60
	// minutesInHour is the number of minutes in an hour.
	minutesInHour = 60
	// hoursInDay is the number of hours in a day.
	hoursInDay = 24

	// Date8Format is the golang time format to use with time.Parse fields of
	// NEM12 format Date(8).
	Date8Format = "20060102"

	// DateTime12Format is the golang time format to use with time.Parse fields of
	// NEM12 format DateTime(12).
	DateTime12Format = "200601021504"

	// DateTime14Format is the golang time format to use with time.Parse fields of
	// NEM12 format DateTime(14).
	DateTime14Format = "20060102150405"
)

var (
	// nemTime is the fixed timezone adopted by AEMO in the NEM, set to UTC+10:00.
	nemTime = time.FixedZone("NEMTIME", (10 * minutesInHour * secondsInMinute)) //nolint:gochecknoglobals
)

// NEMTime returns NEM time, the fixed timezone adopted by AEMO in the NEM, set
// to UTC+10:00.
func NEMTime() *time.Location {
	return nemTime
}
