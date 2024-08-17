package config

import (
	"os"
)

type currentApp struct {
	RestfulAddress string
	GrpcPort       string
}

type kafka struct {
	Addr1 string
	Addr2 string
	Addr3 string
}

type midtrans struct {
	BaseUrl   string
	ServerKey string
}

type Config struct {
	CurrentApp *currentApp
	Kafka      *kafka
	Midtrans   *midtrans
}

var Conf *Config

// *config ini hanya berisi env variable
func init() {
	appStatus := os.Getenv("PRASORGANIC_APP_STATUS")

	if appStatus == "DEVELOPMENT" {

		Conf = setUpForDevelopment()
		return
	}

	Conf = setUpForNonDevelopment(appStatus)
}
