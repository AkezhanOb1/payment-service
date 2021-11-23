package configs

import (
	"github.com/AkezhanOb1/payment/configs/log"
	"github.com/AkezhanOb1/payment/configs/mongo"
	"github.com/joho/godotenv"
)

func init() {
	var err = godotenv.Load("prod.env")
	if err != nil {
		log.Logger.Fatal("Error loading .env file")
	}
	log.Logger.Info("configs are loaded")

	mongo.InitMongo()
}
