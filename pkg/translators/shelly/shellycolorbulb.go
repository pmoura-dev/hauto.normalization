package shelly

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/pmoura-dev/hauto.normalization/pkg/types"
)

func ShellyColorBulbToLightTranslator(
	payload []byte,
	deviceID uuid.UUID,
	deviceType types.DeviceType,
) ([]byte, error) {

	shellyPayload := shellyColorBulbStateMessage{}
	err := json.Unmarshal(payload, &shellyPayload)
	if err != nil {
		return nil, err
	}

	lightMessage := types.LightStateMessage{
		BaseStateMessage: types.BaseStateMessage{
			DeviceID:   deviceID,
			DeviceType: deviceType,
			Timestamp:  time.Now(),
		},
		State: types.LightState{
			IsOn: shellyPayload.IsOn,
		},
	}

	translatedPayload, err := json.Marshal(lightMessage)
	if err != nil {
		return nil, err
	}

	return translatedPayload, nil
}

func ShellyColorBulbToColorLightTranslator() {

}

type shellyColorBulbStateMessage struct {
	IsOn           bool   `json:"ison"`
	HasTimer       bool   `json:"has_timer"`
	TimerStarted   int64  `json:"timer_started"`
	TimerDuration  int    `json:"timer_duration"`
	TimerRemaining int    `json:"timer_remaining"`
	Mode           string `json:"mode"`
	Red            int    `json:"red"`
	Green          int    `json:"green"`
	Blue           int    `json:"blue"`
	White          int    `json:"white"`
	Gain           int    `json:"gain"`
	Temp           int    `json:"temp"`
	Brightness     int    `json:"brightness"`
	Effect         int    `json:"effect"`
}
