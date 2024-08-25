package config

import (
	"os"

	"github.com/dwprz/prasorganic-notification-service/src/common/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func setUpForDevelopment() *Config {
	err := os.Chdir(os.Getenv("PRASORGANIC_NOTIFICATION_SERVICE_WORKSPACE"))
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "config.setUpForDevelopment", "section": "os.Chdir"}).Fatal(err)
	}

	viper := viper.New()
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "config.setUpForDevelopment", "section": "viper.ReadInConfig"}).Fatal(err)
	}

	currentAppConf := new(currentApp)
	currentAppConf.RestfulAddress = viper.GetString("CURRENT_APP_RESTFUL_ADDRESS")
	currentAppConf.GrpcPort = viper.GetString("CURRENT_APP_GRPC_PORT")

	kafkaConf := new(kafka)
	kafkaConf.Addr1 = viper.GetString("KAFKA_ADDRESS_1")
	kafkaConf.Addr2 = viper.GetString("KAFKA_ADDRESS_2")
	kafkaConf.Addr3 = viper.GetString("KAFKA_ADDRESS_3")

	midtransConf := new(midtrans)
	midtransConf.BaseUrl = viper.GetString("MIDTRANS_BASE_URL")
	midtransConf.ServerKey = viper.GetString("MIDTRANS_SERVER_KEY")

	shipperConf := new(shipper)
	shipperConf.BaseUrl = viper.GetString("SHIPPER_BASE_URL")
	shipperConf.ApiKey = viper.GetString("SHIPPER_API_KEY")

	ngrokConf := new(ngrok)
	ngrokConf.BaseUrl = viper.GetString("NGROK_BASE_URL")

	return &Config{
		CurrentApp: currentAppConf,
		Kafka:      kafkaConf,
		Midtrans:   midtransConf,
		Shipper:    shipperConf,
		Ngrok:      ngrokConf,
	}
}
