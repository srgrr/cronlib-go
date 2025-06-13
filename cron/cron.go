/*
	Minimalist Cron Parser.
	It accepts cron expressions of the form:
	P P P P P
	where P can be:
	- A single value (e.g. 5)
	- A range (e.g. 5-10)
	- A list of values (e.g. 5,10,15)
	- An asterisk (*) for "every" value
	- An "* / X" notation w.o. spaces (e.g. "every 5" for every 5th value)
	It won't accept:
	- Month names (e.g. Jan, Feb, etc.)
	- Weekday names (e.g. Mon, Tue, etc.)

	The meaning of every pattern is as follows:
	- Minute: 0-59
	- Hour: 0-23
	- Month Day: 1-31
	- Month: 1-12
	- Week Day: 0-6 (0 = Sunday, 1 = Monday, ..., 6 = Saturday)
*/
package cron
import (
	"fmt"
	"regexp"
	"strings"
	"log"
)


type Cron struct {
	mask string
	minute map[int]bool
	hour map[int]bool
	monthDay map[int]bool
	month map[int]bool
	weekDay map[int]bool
}

// General cron regex patterns
// By general I mean that the following patterns won't check specific ranges for minutes, hours, etc.
var (
	// General pattern for asterisk in cron expressions
	generalAsteriskPattern = "\\*"
	capturingAsteriskPattern = fmt.Sprintf("(%s)", generalAsteriskPattern)

	// General pattern for "every X" in cron expressions
	generalPeriodicPattern = "\\*/\\s*\\d+"
	capturingPeriodicPattern = fmt.Sprintf("(%s)", generalPeriodicPattern)

	// General pattern for matching single values in cron expressions
	generalSingleValuePattern = "\\d+"
	capturingSingleValuePattern = fmt.Sprintf("(%s)", generalSingleValuePattern)

	// General pattern for matching ranges in cron expressions
	generalRangePattern = "\\d+-\\d+"
	capturingRangePattern = fmt.Sprintf("(%s)", generalRangePattern)

	// General pattern for matching lists in cron expressions
	generalListPattern = "\\d+(?:,\\d+)*"
	capturingListPattern = fmt.Sprintf("(%s)", generalListPattern)

	// General position regex for matching cron fields
	// A field can be a single value, a range, or a list
	generalPositionPattern = fmt.Sprintf("%s|%s|%s", generalSingleValuePattern, generalRangePattern, generalListPattern)

	// General regex for matching cron expressions
	generalCronPattern = 
		fmt.Sprintf(
			"^%s\\s+%s\\s+%s\\s+%s\\s+%s$",
			generalPositionPattern,
			generalPositionPattern,
			generalPositionPattern,
			generalPositionPattern,
			generalPositionPattern)
)


// Pre-compile regexes for further use
// Check the patterns above for the meaning of each regex
var (
	rangeRegex = regexp.MustCompile(capturingRangePattern)
	singleValueRegex = regexp.MustCompile(capturingSingleValuePattern)
	listRegex = regexp.MustCompile(capturingListPattern)
	generalCronRegex = regexp.MustCompile(generalCronPattern)
)

// General check for cron expression
func checkGeneralCronRegex(cronString string) bool {
	return generalCronRegex.MatchString(cronString)
}

// Get all the matching values for a cron field
func getMinutes(minute string) map[int]bool {
	if rangeRegex.MatchString(minute) {
		_ = 0
	} else if singleValueRegex.MatchString(minute) {
		_ = 0
	} else if listRegex.MatchString(minute) {
		_ = 0
	}
	return make(map[int]bool) // Placeholder return, should be replaced with actual logic
	panic(fmt.Sprintf("Invalid minute value '%s' in cron expression", minute))
}

func getHours(hour string) map[int]bool {
	if rangeRegex.MatchString(hour) {
		_ = 0
	} else if singleValueRegex.MatchString(hour) {
		_ = 0
	} else if listRegex.MatchString(hour) {
		_ = 0
	}
	return make(map[int]bool) // Placeholder return, should be replaced with actual logic
	panic(fmt.Sprintf("Invalid hour value '%s' in cron expression", hour))
}

func getMonthDays(monthDay string, month string) map[int]bool {
	if rangeRegex.MatchString(monthDay) {
		_ = 0
	} else if singleValueRegex.MatchString(monthDay) {
		_ = 0
	} else if listRegex.MatchString(monthDay) {
		_ = 0
	} else if monthDay == "*" {
		_ = 0
	}
	return make(map[int]bool) // Placeholder return, should be replaced with actual logic
	panic(fmt.Sprintf("Invalid month day value '%s' in cron expression", monthDay))
}

func getMonths(month string) map[int]bool {
	if rangeRegex.MatchString(month) {
		_ = 0
	} else if singleValueRegex.MatchString(month) {
		_ = 0
	} else if listRegex.MatchString(month) {
		_ = 0
	} else if month == "*" {
		_ = 0
	}
	return make(map[int]bool) // Placeholder return, should be replaced with actual logic
	panic(fmt.Sprintf("Invalid month value '%s' in cron expression", month))
}

func getWeekDays(weekDay string) map[int]bool {
	if rangeRegex.MatchString(weekDay) {
		_ = 0
	} else if singleValueRegex.MatchString(weekDay) {
		_ = 0
	} else if listRegex.MatchString(weekDay) {
		_ = 0
	} else if weekDay == "*" {
		_ = 0
	}
	return make(map[int]bool) // Placeholder return, should be replaced with actual logic
	panic(fmt.Sprintf("Invalid week day value '%s' in cron expression", weekDay))
}


// TODO: Implement the specific checks for each cron field
// Particular checks for each cron field
func checkMinute(minute string) bool {
	return true
}

func checkHour(hour string) bool {
	return true
}

func checkMonthDay(monthDay string, month string) bool {
	return true
}

func checkMonth(month string) bool {
	return true
}

func checkWeekDay(weekDay string) bool {
	return true
}


// Exported functions

// CronFromString parses a cron expression string and returns a Cron struct.
// It panics if the cron expression is invalid.
// It expects the cron expression to be in the format:
// "P P P P P" where P can be a single value, a range, a list, an asterisk, or an "* / X" notation.
// It does not accept month or weekday names.
func CronFromString(cronString string) Cron {
	if !checkGeneralCronRegex(cronString) {
		panic(fmt.Sprintf("Cron expression '%s' doesn't follow general format for cron expressions", cronString))
	}

	// Now it's safe to split the cron string into its components
	parts := strings.Split(cronString, " ")

	log.Printf("Cron mask %s has %d parts: %s", cronString, len(parts), parts)

	minute := parts[0]
	hour := parts[1]
	monthDay := parts[2]
	month := parts[3]
	weekDay := parts[4]

	// Check each part of the cron expression
	if !checkMinute(minute) {
		panic(fmt.Sprintf("Invalid minute value '%s' in cron expression '%s'", minute, cronString))
	}
	if !checkHour(hour) {
		panic(fmt.Sprintf("Invalid hour value '%s' in cron expression '%s'", hour, cronString))
	}
	if !checkMonthDay(monthDay, month) {
		panic(fmt.Sprintf("Invalid month day value '%s' in cron expression '%s'", monthDay, cronString))
	}
	if !checkMonth(month) {
		panic(fmt.Sprintf("Invalid month value '%s' in cron expression '%s'", month, cronString))
	}
	if !checkWeekDay(weekDay) {
		panic(fmt.Sprintf("Invalid week day value '%s' in cron expression '%s'", weekDay, cronString))
	}

	return Cron{
		mask: cronString,
		minute: getMinutes(minute),
		hour: getHours(hour),
		monthDay: getMonthDays(monthDay, month),
		month: getMonths(month),
		weekDay: getWeekDays(weekDay),
	}
}


// CronMatchesTime checks if a given time matches the cron expression.
// It takes a Cron struct and the time components (minute, hour, monthDay, month, weekDay) as arguments.
// It returns true if the time matches the cron expression, false otherwise.
// TODO: This function is a placeholder and should be implemented with actual logic.
func CronMatchesTime(cron Cron, minute int, hour int, monthDay int, month int, weekDay int) bool {
	return true
}	