package shelly

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/pmoura-dev/hauto.normalization/pkg/types"
)

type ShellyColorBulbTranslator struct{}

func (t ShellyColorBulbTranslator) TranslateIngestion(deviceID uuid.UUID, deviceType string, message []byte) ([]byte, error) {

	shellyMsg := shellyColorBulbStateMessage{}
	err := json.Unmarshal(message, &shellyMsg)
	if err != nil {
		return nil, err
	}

	switch deviceType {
	case "light":
		return t.toLightStateMessage(deviceID, shellyMsg)
	}

	return nil, nil
}

func (t ShellyColorBulbTranslator) toLightStateMessage(deviceID uuid.UUID, shellyMsg shellyColorBulbStateMessage) ([]byte, error) {
	lightMsg := types.LightStateMessage{
		BaseStateMessage: types.BaseStateMessage{
			DeviceID:  deviceID,
			Timestamp: time.Now(),
		},
		IsOn: shellyMsg.IsOn,
	}

	payload, err := json.Marshal(lightMsg)
	if err != nil {
		return nil, err
	}

	return payload, nil
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
