package port

import (
	"modak_golang/internal/domain/model"
	"modak_golang/internal/domain/model/valueobjects"
	"time"
)

type NotificationRepository interface {
	SendNotification(n model.Notification) error
	GetLastNotificationDate(recipientId string) time.Time
	GetNotificationsCountForRecipient(recipientId string, interval valueobjects.IntervalType) int
}
