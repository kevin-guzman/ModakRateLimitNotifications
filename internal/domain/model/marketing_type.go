package model

import "modak_golang/internal/domain/model/valueobjects"

type marketingNotificationBuilder struct {
	notification Notification
}

func NewMarketingNotificationBuilder() *marketingNotificationBuilder {
	return &marketingNotificationBuilder{
		notification: Notification{
			nType:        Market,
			rateLimit:    3,
			rateInterval: valueobjects.Hour,
		},
	}
}

func (b *marketingNotificationBuilder) SetUserID(userID string) *marketingNotificationBuilder {
	b.notification.userId = userID
	return b
}

func (b *marketingNotificationBuilder) SetMessage(message string) *marketingNotificationBuilder {
	b.notification.message = message
	return b
}

func (b *marketingNotificationBuilder) SetRecipientID(recipientID string) *marketingNotificationBuilder {
	b.notification.recipientId = recipientID
	return b
}

func (b *marketingNotificationBuilder) Build() Notification {
	return b.notification
}
