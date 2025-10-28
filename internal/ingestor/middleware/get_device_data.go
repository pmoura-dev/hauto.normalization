package middleware

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/pmoura-dev/beacon"
	"github.com/pmoura-dev/hauto.normalization/pkg/types"
)

var (
	ErrDeviceDataNotFound = errors.New("device data not found in registry")
)

func GetDeviceData(next beacon.HandlerFunc) beacon.HandlerFunc {

	return func(ctx context.Context, publisher beacon.Publisher, message beacon.RoutedMessage) error {

		// simulate call to the registry/database
		devices := map[string]types.DeviceData{
			"shellycolorbulb-12345": {
				DeviceID:     uuid.MustParse("b70bb447-825e-4692-8248-bf7cc3564fd9"),
				Type:         "LIGHT",
				Manufacturer: "shelly",
				Model:        "shelly_color_bulb",
			},
			"shellycolorbulb-67890": {
				DeviceID:     uuid.MustParse("d62820d7-4b63-4949-8e05-3b9187dc261b"),
				Type:         "LIGHT",
				Manufacturer: "shelly",
				Model:        "shelly_color_bulb",
			},
		}

		externalID := message.GetTopicParam("external_id")

		deviceData, ok := devices[externalID]
		if !ok {
			return ErrDeviceDataNotFound
		}

		// store device data in context
		const deviceDataKey = "DEVICE_DATA"
		ctx = context.WithValue(ctx, deviceDataKey, deviceData)

		next(ctx, publisher, message)
		return nil
	}
}
