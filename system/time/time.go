package time

import (
	"fmt"
	"time"
)

// TimeUntilNextRun returns the duration until a given schedule hour and minute based on location.
//
// locations: "America/New_York", "America/ Los_Angeles", or "UTC". Use anything in IANA Time DB
//
// scheduledHour: these are single int. e.g. For 6am, don't put 06, but 6.
func TimeUntilNextRun(location string, scheduledHour, scheduledMin int) (time.Duration, error) {
	loc, err := time.LoadLocation(location)
	if err != nil {
		fmt.Println("Error loading location:", err)
		return 0, err
	}

	for {
		now := time.Now().In(loc)
		dateTimeScheduled := time.Date(now.Year(), now.Month(), now.Day(), scheduledHour, scheduledMin, 0, 0, loc)
		if now.After(dateTimeScheduled) {
			dateTimeScheduled = dateTimeScheduled.Add(24 * time.Hour)
		}

		return time.Until(dateTimeScheduled), nil
	}
}

// GetFirstDayOfThisMonth returns the 1st day of this month in time.Time, just format to the desired format.
//
// e.g. log.Println("First day of this month:", firstDayOfCurrentMonth.Format("2006-01-02"))
//
// e.g. log.Println("First day of the month in Unix:", firstDayOfCurrentMonth.Unix() )
func GetFirstDayOfThisMonth() time.Time {
	now := time.Now()
	firstDayOfCurrentMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	//log.Println("First day of this month:", firstDayOfCurrentMonth.Format("2006-01-02"))
	return firstDayOfCurrentMonth
}

// GetLastDayOfLastMonth returns the 1st day of this month in time.Time, just format to the desired format.
//
// e.g. log.Println("First day of this month:", lastDayOfLastMonth.Format("2006-01-02"))
//
// e.g. log.Println("First day of the month in Unix:", lastDayOfLastMonth.Unix() )
func GetLastDayOfLastMonth() time.Time {
	firstDayOfCurrentMonth := GetFirstDayOfThisMonth()
	lastDayOfLastMonth := firstDayOfCurrentMonth.AddDate(0, 0, -1)
	//log.Println("Last day of last month:", lastDayOfLastMonth.Format("2006-01-02"))
	return lastDayOfLastMonth
}

// GetFirstDayOfLastMonth returns the 1st day of this month in time.Time, just format to the desired format.
//
// e.g. log.Println("First day of this month:", firstDayOfLastMonth.Format("2006-01-02"))
//
// e.g. log.Println("First day of the month in Unix:", firstDayOfLastMonth.Unix() )
func GetFirstDayOfLastMonth() time.Time {
	lastDayOfLastMonth := GetLastDayOfLastMonth()
	firstDayOfLastMonth := time.Date(lastDayOfLastMonth.Year(), lastDayOfLastMonth.Month(), 1, 0, 0, 0, 0, lastDayOfLastMonth.Location())
	//log.Println("First day of last month:", firstDayOfLastMonth.Format("2006-01-02"))
	return firstDayOfLastMonth
}
