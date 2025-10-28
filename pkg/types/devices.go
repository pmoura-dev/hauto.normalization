package types

import "github.com/google/uuid"

type DeviceType string

const (
	DeviceTypeLight      = "LIGHT"
	DeviceTypeColorLight = "COLOR_LIGHT"
)

type DeviceData struct {
	DeviceID     uuid.UUID  `json:"device_id"`
	Type         DeviceType `json:"type"`
	Manufacturer string     `json:"manufacturer"`
	Model        string     `json:"model"`
}
