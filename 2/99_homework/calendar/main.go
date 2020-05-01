package main

import (
	"fmt"
	"time"
	"math"
)

// Calendar type
type Calendar struct {
	currentMonth int
	currentDay int
	currentYear int
}

// NewCalendar Calendar struct constructor
func NewCalendar(currentTime time.Time) *Calendar {
	calendar := new(Calendar)
	calendar.currentMonth = int(currentTime.Month())
	calendar.currentDay = currentTime.Day()
	calendar.currentYear = currentTime.Year()

	return calendar
}

// CurrentQuarter method
func (calendar *Calendar) CurrentQuarter() int {
	return int(math.Ceil(float64(calendar.currentMonth) / float64(3)))
}

func main() {
	parsed, _ := time.Parse("2006-01-02", "2015-01-15")

	fmt.Println(*NewCalendar(parsed))
}