package model

import "modak_golang/internal/domain/model/valueobjects"

type statusNotificationBuilder struct {
	notification Notification
}

func NewStatusNotificationBuilder() *statusNotificationBuilder {
	return &statusNotificationBuilder{
		notification: Notification{
			nType:        Status,
			rateLimit:    2,
			rateInterval: valueobjects.Minute,
		},
	}
}

func (b *statusNotificationBuilder) SetUserID(userID string) *statusNotificationBuilder {
	b.notification.userId = userID
	return b
}

func (b *statusNotificationBuilder) SetMessage(message string) *statusNotificationBuilder {
	b.notification.message = message
	return b
}

func (b *statusNotificationBuilder) SetRecipientID(recipientID string) *statusNotificationBuilder {
	b.notification.recipientId = recipientID
	return b
}

func (b *statusNotificationBuilder) Build() Notification {
	return b.notification
}
