package service

import (
	"fmt"
	"modak_golang/internal/domain/model"
	"modak_golang/internal/domain/port"
)

type NotificationService struct {
	notificationRepository port.NotificationRepository
}

func NewNotificationService(notificationRepository port.NotificationRepository) *NotificationService {
	return &NotificationService{
		notificationRepository,
	}
}

func (ns NotificationService) SendNotification(n model.Notification) model.AbuseError {
	lastNotificationTime := ns.notificationRepository.GetLastNotificationDate(n.GetRecipientId())
	count := ns.notificationRepository.GetNotificationsCountForRecipient(n.GetRecipientId(), n.GetRateInterval())

	isUnderLimit := n.IsUnderLimit(lastNotificationTime, count)
	if !isUnderLimit {
		err := fmt.Errorf("limit reached")
		details := fmt.Sprintf("limit reached for user %s", n.GetRecipientId())
		return model.NewAbuseError(err, details)
	}

	err := ns.notificationRepository.SendNotification(n)
	if err != nil {
		details := fmt.Sprintf("failed to send notification to user %s", n.GetRecipientId())
		return model.NewAbuseError(err, details)
	}

	return nil
}
