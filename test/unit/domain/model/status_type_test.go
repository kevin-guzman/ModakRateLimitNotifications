package model_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"modak_golang/internal/domain/model"
)

var _ = Describe("Testing Status Notification", func() {
	It("Should return true because it is under time range", func() {
		notification := model.NewStatusNotificationBuilder().
			SetMessage("test").
			SetUserID("test").
			SetRecipientID("test").
			Build()

		Expect(notification.IsUnderLimit(time.Now().Add(-time.Minute), 0)).To(Equal(true))
		Expect(notification.IsUnderLimit(time.Now().Add(-time.Hour), 0)).To(Equal(true))
		Expect(notification.IsUnderLimit(time.Now().Add(-time.Hour*24), 0)).To(Equal(true))
	})

	It("Should return true because it is under time range and count is under limit", func() {
		notification := model.NewStatusNotificationBuilder().
			SetMessage("test").
			SetUserID("test").
			SetRecipientID("test").
			Build()

		Expect(notification.IsUnderLimit(time.Now(), 1)).To(Equal(true))
	})

	It("Should return false because it is under time range and count is above limit", func() {
		notification := model.NewStatusNotificationBuilder().
			SetMessage("test").
			SetUserID("test").
			SetRecipientID("test").
			Build()

		Expect(notification.IsUnderLimit(time.Now(), 2)).To(Equal(false))
	})
})
