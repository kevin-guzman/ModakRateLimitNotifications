package model_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"modak_golang/internal/domain/model"
	"modak_golang/internal/domain/model/valueobjects"
)

var _ = Describe("Testing Notification", func() {
	It("Should GetRecipientId and GetRateInterval", func() {
		notification := model.NewStatusNotificationBuilder().
			SetMessage("test").
			SetUserID("test").
			SetRecipientID("test").
			Build()

		Expect(notification.GetRecipientId()).To(Equal("test"))
		Expect(notification.GetRateInterval()).To(Equal(valueobjects.Minute))
	})
})
