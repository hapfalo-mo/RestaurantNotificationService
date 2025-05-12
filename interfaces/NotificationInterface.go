package interfaces

import (
	"restaurantnotificationservice/custom"
	"restaurantnotificationservice/dto"
	"restaurantnotificationservice/models"
)

type NotificationInterface interface {
	CreateNewNotification(request *dto.NotificationRequest) (custom.Data[models.Notification], custom.ErrorCustomStruct)
}
