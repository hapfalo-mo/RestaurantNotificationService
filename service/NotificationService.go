package service

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"restaurantnotificationservice/custom"
	"restaurantnotificationservice/dto"
	"restaurantnotificationservice/interfaces"
	"restaurantnotificationservice/models"
	"time"
)

var _ interfaces.NotificationInterface = &NotificationService{}

type NotificationService struct{}

func (s *NotificationService) CreateNewNotification(request *dto.NotificationRequest) (response custom.Data[models.Notification], err custom.ErrorCustomStruct) {
	if request.Content == "" || request.Message == "" {
		err.Message = "<MNK>"
		err.ErrorField = "Null Value"
		err.Field = "Service and Notification"
		return custom.Data[models.Notification]{}, err
	}
	newNotification := models.Notification{
		Id:             primitive.NewObjectID(),
		NotificationID: uuid.New().String(),
		UserId:         request.UserId,
		Content:        request.Content,
		Message:        request.Message,
		CreatedAt:      time.Now().String(),
		UpdatedAt:      "",
		DeletedAt:      "",
		Status:         1,
	}
	return custom.Data[models.Notification]{newNotification}, custom.ErrorCustomStruct{}
}
