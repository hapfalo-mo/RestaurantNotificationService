package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Notification struct {
	Id             primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	NotificationID string             `json:"notificationID" bson:"NotificationID, omitempty"`
	UserId         int64              `json:"userId" bson:"UserID, omitempty"`
	Content        string             `json:"content" bson:"Content"`
	Message        string             `json:"message" bson:"Message"`
	CreatedAt      string             `json:"createdAt" bson:"CreatedAt"`
	UpdatedAt      string             `json:"updatedAt" bson:"UpdatedAt"`
	DeletedAt      string             `json:"deletedAt" bson:"DeletedAt"`
	Status         int64              `json:"status" bson:"status"`
}
