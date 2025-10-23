package handlers

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pmoura-dev/beacon"
	"github.com/pmoura-dev/hauto.normalization/pkg/translators"
	"github.com/pmoura-dev/hauto.normalization/pkg/translators/shelly"
	"github.com/pmoura-dev/hauto.normalization/pkg/types"
)

func getDeviceInfo(externalID string) (types.DeviceInfo, error) {

	devices := map[string]types.DeviceInfo{
		"shellycolorbulb-12345": types.DeviceInfo{
			DeviceID: uuid.MustParse("b70bb447-825e-4692-8248-bf7cc3564fd9"),
			Model:    "shelly_color_bulb",
			Type:     "light",
		},
		"shellycolorbulb-67890": types.DeviceInfo{
			DeviceID: uuid.MustParse("d62820d7-4b63-4949-8e05-3b9187dc261b"),
			Model:    "shelly_color_bulb",
			Type:     "light",
		},
	}

	return devices[externalID], nil
}

const (
	modelShellyColorBulb = "shelly_color_bulb"
)

func ShellyState(publisher beacon.Publisher, message beacon.RoutedMessage) error {

	externalID := message.GetTopicParam("shelly_id")

	// 1. Get device info from registry
	deviceInfo, err := getDeviceInfo(externalID)
	if err != nil {
		return err
	}

	var translator translators.Translator
	switch deviceInfo.Model {
	case modelShellyColorBulb:
		// shellycolorbulb translator
		translator = shelly.ShellyColorBulbTranslator{}
	}

	translatedPayload, err := translator.TranslateIngestion(deviceInfo.DeviceID, deviceInfo.Type, message.Payload)
	if err != nil {
		return err
	}

	// publish message
	fmt.Println(string(translatedPayload))
	return nil
}
