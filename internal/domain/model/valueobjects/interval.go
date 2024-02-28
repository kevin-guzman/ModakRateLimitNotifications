package valueobjects

import (
	"time"
)

type IntervalType string

const (
	Minute IntervalType = "minute"
	Hour   IntervalType = "hour"
	Day    IntervalType = "day"
)

var toDuration map[IntervalType]time.Duration = map[IntervalType]time.Duration{
	Minute: time.Minute,
	Hour:   time.Hour,
	Day:    time.Hour * 24,
}

func (it IntervalType) IsUnderTimeRage(comparisonDate time.Time) bool {
	now := time.Now()
	timeDifference := now.Sub(comparisonDate)
	duration := toDuration[it]

	return timeDifference >= duration
}

func (it IntervalType) String() string {
	return string(it)
}
