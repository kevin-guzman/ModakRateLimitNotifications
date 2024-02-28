package service_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"modak_golang/internal/domain/model"
	"modak_golang/internal/domain/service"
	"modak_golang/test/utils"
	"modak_golang/test/utils/mocks"
)

var (
	notificationRepository *mocks.NotificationRepository
	notificationSerice     service.NotificationService
)

var _ = Describe("Testing Notification Service", func() {
	notification := model.NewMarketingNotificationBuilder().
		SetMessage("test").
		SetUserID("test").
		SetRecipientID("test").
		Build()

	BeforeEach(func() {
		notificationRepository = &mocks.NotificationRepository{}
		notificationSerice = *service.NewNotificationService(notificationRepository)
	})

	Context("Sending a notification", func() {
		It("Should return AbuseError because it is not under limit", func() {
			notificationRepository.On("GetLastNotificationDate", notification.GetRecipientId()).Return(time.Now())
			notificationRepository.On("GetNotificationsCountForRecipient", notification.GetRecipientId(), notification.GetRateInterval()).Return(3)

			err := notificationSerice.SendNotification(notification)

			Expect(err).To(HaveOccurred())
		})

		It("Should return nil because it is under limit", func() {
			notificationRepository.On("GetLastNotificationDate", notification.GetRecipientId()).Return(time.Now().Add(-time.Hour))
			notificationRepository.On("GetNotificationsCountForRecipient", notification.GetRecipientId(), notification.GetRateInterval()).Return(3)
			notificationRepository.On("SendNotification", notification).Return(nil)

			err := notificationSerice.SendNotification(notification)

			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("Given bad repositories", func() {
		It("Should return error on SendNotification", func() {
			notificationRepository.On("GetLastNotificationDate", notification.GetRecipientId()).Return(time.Now().Add(-time.Hour))
			notificationRepository.On("GetNotificationsCountForRecipient", notification.GetRecipientId(), notification.GetRateInterval()).Return(3)
			notificationRepository.On("SendNotification", notification).Return(utils.ErrInfrastructure)

			err := notificationSerice.SendNotification(notification)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal(utils.ErrInfrastructure.Error()))
			Expect(err.Details()).To(Equal("failed to send notification to user test"))
		})
	})
})
