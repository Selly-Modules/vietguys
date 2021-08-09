package vietguys

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/Selly-Modules/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDBConfig ...
type MongoDBConfig struct {
	Host, User, Password, DBName, Mechanism, Source string
}

// Config ...
type Config struct {
	// Endpoint which will use to call api
	Endpoint string
	// For auth
	User string
	// For auth
	Password string
	// Brand name
	From string
	// MongoDB config, for save documents
	MongoDB MongoDBConfig
	// Limit send time per ip
	IPMaxSendPerDay int
	// Limit send time per phone number
	PhoneMaxSendPerDay int
}

// Service ...
type Service struct {
	Config
	Client *http.Client
	DB     *mongo.Collection
}

var s *Service
var bgCtx = context.Background()

// NewInstance for using send sms method
func NewInstance(config Config) error {
	if config.Endpoint == "" || config.User == "" || config.Password == "" || config.From == "" || config.MongoDB.Host == "" {
		return errors.New("please provide all information that needed: endpoint, user, password, from, mongodb")
	}

	// Connect MongoDB
	err := mongodb.Connect(
		config.MongoDB.Host,
		config.MongoDB.User,
		config.MongoDB.Password,
		config.MongoDB.DBName,
		config.MongoDB.Mechanism,
		config.MongoDB.Source,
	)
	if err != nil {
		fmt.Println("Cannot init module VIETGUYS", err)
		return err
	}

	s = &Service{
		Config: config,
		Client: &http.Client{},
		DB:     mongodb.GetInstance().Collection(logCollectionName),
	}

	return nil
}

// GetInstance ...
func GetInstance() *Service {
	return s
}
