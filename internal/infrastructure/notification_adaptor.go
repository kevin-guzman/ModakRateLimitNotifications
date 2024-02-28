package infrastructure

import (
	"modak_golang/internal/domain/model"
	"modak_golang/internal/domain/model/valueobjects"
	"modak_golang/internal/domain/port"
	"time"
)

type notificationAdaptor struct {
	// Dependencies like kvs, or caches
}

func NewNotificationAdaptor() port.NotificationRepository {
	return &notificationAdaptor{}
}

func (na notificationAdaptor) SendNotification(n model.Notification) error {
	return nil
}

func (na notificationAdaptor) GetLastNotificationDate(recipientId string) time.Time {
	return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 10, 20, 1, 0, time.Now().Location())
}

func (na notificationAdaptor) GetNotificationsCountForRecipient(recipientId string, interval valueobjects.IntervalType) int {
	return 2
}
