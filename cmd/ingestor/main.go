package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/pmoura-dev/beacon"
	"github.com/pmoura-dev/beacon/publishers"
	"github.com/pmoura-dev/beacon/subscribers"
	"github.com/pmoura-dev/hauto.normalization/internal/ingestor/handlers"
	"github.com/pmoura-dev/hauto.normalization/internal/ingestor/middleware"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	mqttURL := "mqtt://broker.emqx.io:1883"

	subscriber := subscribers.NewMQTTSubscriber(mqttURL)
	publisher := publishers.NewMQTTPublisher(mqttURL)

	r := beacon.NewRouter(
		beacon.NewBroker(subscriber, publisher),
	)

	r.UseMiddleware(middleware.GetDeviceData)

	addShellySubscriptions(r)

	if err := r.Start(); err != nil {
		log.Fatal(err)
	}

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := r.Shutdown(shutdownCtx); err != nil {
		log.Fatal("Error shutting down Beacon.")
	}
}

func addShellySubscriptions(r *beacon.Router) {
	_ = r.AddSubscription("shellies/{external_id}/online", handlers.ShellyAvailability)

	_ = r.AddSubscription("shellies/{external_id}/color/0/status", handlers.ShellyColorBulbState)
}
