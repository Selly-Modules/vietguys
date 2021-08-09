package vietguys

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Log ...
type Log struct {
	ID          primitive.ObjectID `bson:"_id"`
	Carrier     string             `bson:"carrier"`
	Type        string             `bson:"type"`
	PhoneNumber string             `bson:"phoneNumber"`
	Code        string             `bson:"code"`
	IsCodeValid bool               `bson:"isCodeValid"`
	Content     string             `bson:"content"`
	IP          string             `bson:"ip"`
	Success     bool               `bson:"success"`
	Result      string             `bson:"result"`
	CreatedAt   time.Time          `bson:"createdAt"`

	tableName string
}

// Save log to db
func (s Service) saveLog(doc Log) {
	if _, err := s.DB.InsertOne(bgCtx, doc); err != nil {
		fmt.Println("*** Error when create log", err)
		fmt.Println("*** Log", doc)
	}
}
