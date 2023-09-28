package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type BusPreview struct {
	Id primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	EndingPointName string `json:"ending-point-name"`
	BusName         string `json:"bus-name"`
	Duration        int    `json:"duration"`
	TicketPrice     int    `json:"ticket-price"`
}
