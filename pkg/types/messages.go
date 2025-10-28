package types

import (
	"time"

	"github.com/google/uuid"
)

type BaseStateMessage struct {
	DeviceID   uuid.UUID  `json:"device_id"`
	DeviceType DeviceType `json:"device_type"`
	Timestamp  time.Time  `json:"timestamp"`
}

type AvailabilityMessage struct {
	BaseStateMessage
	Online bool `json:"online"`
}

type LightState struct {
	IsOn bool `json:"is_on"`
}

type LightStateMessage struct {
	BaseStateMessage
	State LightState `json:"state"`
}
