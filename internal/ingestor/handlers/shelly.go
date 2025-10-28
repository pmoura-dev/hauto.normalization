package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/pmoura-dev/beacon"
	shelly_translators "github.com/pmoura-dev/hauto.normalization/pkg/translators/shelly"
	"github.com/pmoura-dev/hauto.normalization/pkg/types"
)

var (
	ErrInternalError = errors.New("something bad happened")
)

func ShellyAvailability(ctx context.Context, publisher beacon.Publisher, message beacon.RoutedMessage) error {
	deviceData := ctx.Value("DEVICE_DATA").(types.DeviceData)

	rawTopic := fmt.Sprintf(types.AvailabilityTopicFormat, deviceData.DeviceID)
	topic, err := beacon.NewTopic(rawTopic)
	if err != nil {
		log.Printf("Error: [%v]\n", err)
		return ErrInternalError
	}

	payload, err := shelly_translators.ShellyAvailabilityTranslator(
		message.Payload,
		deviceData.DeviceID,
		deviceData.Type,
	)
	if err != nil {
		log.Printf("Error: [%v]\n", err)
		return ErrInternalError
	}

	publisher.Publish(topic, beacon.Message{Payload: payload})
	return nil
}

func ShellyColorBulbState(ctx context.Context, publisher beacon.Publisher, message beacon.RoutedMessage) error {
	deviceData := ctx.Value("DEVICE_DATA").(types.DeviceData)

	// internal topic
	rawTopic := fmt.Sprintf(types.StateTopicFormat, deviceData.DeviceID)
	topic, err := beacon.NewTopic(rawTopic)
	if err != nil {
		log.Printf("Error: [%v]\n", err)
		return ErrInternalError
	}

	var payload []byte

	switch deviceData.Type {
	case types.DeviceTypeLight:
		payload, err = shelly_translators.ShellyColorBulbToLightTranslator(
			message.Payload,
			deviceData.DeviceID,
			deviceData.Type,
		)
	default:
		log.Printf("Device type not implemented")
		return nil
	}

	if err != nil {
		log.Printf("Error: [%v]\n", err)
		return ErrInternalError
	}

	publisher.Publish(topic, beacon.Message{Payload: payload})
	return nil
}
