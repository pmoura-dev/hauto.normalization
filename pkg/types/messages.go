package types

import (
	"time"

	"github.com/google/uuid"
)

type BaseStateMessage struct {
	DeviceID  uuid.UUID `json:"device_id"`
	Timestamp time.Time `json:"timestamp"`
}

type LightStateMessage struct {
	BaseStateMessage
	IsOn bool `json:"is_on"`
}
