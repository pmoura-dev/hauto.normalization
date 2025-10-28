package shelly

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/pmoura-dev/hauto.normalization/pkg/types"
)

func ShellyAvailabilityTranslator(
	payload []byte,
	deviceID uuid.UUID,
	deviceType types.DeviceType,
) ([]byte, error) {

	shellyPayload := string(payload)

	availabilityMessage := types.AvailabilityMessage{
		BaseStateMessage: types.BaseStateMessage{
			DeviceID:   deviceID,
			DeviceType: deviceType,
			Timestamp:  time.Now(),
		},
		Online: shellyPayload == "true",
	}

	translatedPayload, err := json.Marshal(availabilityMessage)
	if err != nil {
		return nil, err
	}

	return translatedPayload, nil
}
