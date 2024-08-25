package config

import (
	"context"
	"os"
	"strings"

	"github.com/dwprz/prasorganic-notification-service/src/common/log"
	vault "github.com/hashicorp/vault/api"
	"github.com/sirupsen/logrus"
)

func setUpForNonDevelopment(appStatus string) *Config {
	defaultConf := vault.DefaultConfig()
	defaultConf.Address = os.Getenv("PRASORGANIC_CONFIG_ADDRESS")

	client, err := vault.NewClient(defaultConf)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "config.setUpForNonDevelopment", "section": "vault.NewClient"}).Fatal(err)
	}

	client.SetToken(os.Getenv("PRASORGANIC_CONFIG_TOKEN"))

	mountPath := "prasorganic-secrets" + "-" + strings.ToLower(appStatus)

	notifServiceSecrets, err := client.KVv2(mountPath).Get(context.Background(), "notification-service")
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "config.setUpForNonDevelopment", "section": "KVv2.Get"}).Fatal(err)
	}

	kafkaSecrets, err := client.KVv2(mountPath).Get(context.Background(), "kafka")
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "config.setUpForNonDevelopment", "section": "KVv2.Get"}).Fatal(err)
	}

	midtransSecrets, err := client.KVv2(mountPath).Get(context.Background(), "midtrans")
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "config.setUpForNonDevelopment", "section": "KVv2.Get"}).Fatal(err)
	}

	shipperSecrets, err := client.KVv2(mountPath).Get(context.Background(), "shipper")
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "config.setUpForNonDevelopment", "section": "KVv2.Get"}).Fatal(err)
	}

	ngrokSecrets, err := client.KVv2(mountPath).Get(context.Background(), "ngrok")
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "config.setUpForNonDevelopment", "section": "KVv2.Get"}).Fatal(err)
	}

	currentAppConf := new(currentApp)
	currentAppConf.RestfulAddress = notifServiceSecrets.Data["RESTFUL_ADDRESS"].(string)
	currentAppConf.GrpcPort = notifServiceSecrets.Data["GRPC_PORT"].(string)

	kafkaConf := new(kafka)
	kafkaConf.Addr1 = kafkaSecrets.Data["ADDRESS_1"].(string)
	kafkaConf.Addr2 = kafkaSecrets.Data["ADDRESS_2"].(string)
	kafkaConf.Addr3 = kafkaSecrets.Data["ADDRESS_3"].(string)

	midtransConf := new(midtrans)
	midtransConf.BaseUrl = midtransSecrets.Data["BASE_URL"].(string)
	midtransConf.ServerKey = midtransSecrets.Data["SERVER_KEY"].(string)

	shipperConf := new(shipper)
	shipperConf.BaseUrl = shipperSecrets.Data["BASE_URL"].(string)
	shipperConf.ApiKey = shipperSecrets.Data["API_KEY"].(string)

	ngrokConf := new(ngrok)
	ngrokConf.BaseUrl = ngrokSecrets.Data["BASE_URL"].(string)

	return &Config{
		CurrentApp: currentAppConf,
		Kafka:      kafkaConf,
		Midtrans:   midtransConf,
		Shipper:    shipperConf,
		Ngrok:      ngrokConf,
	}
}
