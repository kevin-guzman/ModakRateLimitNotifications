package model

import (
	"modak_golang/internal/domain/model/valueobjects"
	"time"
)

type NotificationType string

const (
	Status NotificationType = "status"
	News   NotificationType = "news"
	Market NotificationType = "market"
)

type Notification struct {
	nType        NotificationType
	userId       string
	message      string
	rateLimit    int
	rateInterval valueobjects.IntervalType
	recipientId  string
}

func (n Notification) IsUnderLimit(lastNotificationTime time.Time, coutn int) bool {
	isUnderTimeRange := n.rateInterval.IsUnderTimeRage(lastNotificationTime)
	if !isUnderTimeRange {
		return coutn < n.rateLimit
	}

	return true
}

func (n Notification) GetRecipientId() string {
	return n.recipientId
}

func (n Notification) GetRateInterval() valueobjects.IntervalType {
	return n.rateInterval
}
