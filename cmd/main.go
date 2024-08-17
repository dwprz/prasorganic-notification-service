package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/dwprz/prasorganic-notification-service/src/core/broker"
	"github.com/dwprz/prasorganic-notification-service/src/core/restful"
	"github.com/dwprz/prasorganic-notification-service/src/service"
)

func HandleCloseApp(cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		cancel()
	}()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	HandleCloseApp(cancel)

	kafkaClient := broker.InitKafkaProducer()
	defer kafkaClient.Close()

	notifService := service.NewNotification(kafkaClient)

	restfulServer := restful.InitServer(notifService)
	defer restfulServer.Stop()

	go restfulServer.Run()

	<-ctx.Done()
}
