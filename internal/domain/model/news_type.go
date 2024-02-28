package model

import "modak_golang/internal/domain/model/valueobjects"

type newsNotificationBuilder struct {
	notification Notification
}

func NewNewsNotificationBuilder() *newsNotificationBuilder {
	return &newsNotificationBuilder{
		notification: Notification{
			nType:        News,
			rateLimit:    1,
			rateInterval: valueobjects.Day,
		},
	}
}

func (b *newsNotificationBuilder) SetUserID(userID string) *newsNotificationBuilder {
	b.notification.userId = userID
	return b
}

func (b *newsNotificationBuilder) SetMessage(message string) *newsNotificationBuilder {
	b.notification.message = message
	return b
}

func (b *newsNotificationBuilder) SetRecipientID(recipientID string) *newsNotificationBuilder {
	b.notification.recipientId = recipientID
	return b
}

func (b *newsNotificationBuilder) Build() Notification {
	return b.notification
}
