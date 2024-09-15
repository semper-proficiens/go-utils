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
