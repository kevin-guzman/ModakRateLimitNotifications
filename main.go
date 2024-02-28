package main

import (
	"fmt"
	"modak_golang/internal/domain/model"
	"modak_golang/internal/domain/service"
	"modak_golang/internal/infrastructure"
)

func main() {
	notificationAdaptor := infrastructure.NewNotificationAdaptor()
	notificationService := service.NewNotificationService(notificationAdaptor)

	statusNotification := model.NewStatusNotificationBuilder().
		SetMessage("hello").
		SetRecipientID("12").
		SetUserID("1").
		Build()
	err := notificationService.SendNotification(statusNotification)
	fmt.Println("status notification result:", err)

	newsNotification := model.NewNewsNotificationBuilder().
		SetMessage("hello").
		SetRecipientID("12").
		SetUserID("1").
		Build()

	err = notificationService.SendNotification(newsNotification)
	fmt.Println("news notification result:", err)

	marketingNotification := model.NewMarketingNotificationBuilder().
		SetMessage("hello").
		SetRecipientID("12").
		SetUserID("1").
		Build()

	err = notificationService.SendNotification(marketingNotification)
	fmt.Println("marketing notification result:", err)

}
