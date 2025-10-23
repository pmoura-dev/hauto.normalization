package types

import "github.com/google/uuid"

type DeviceInfo struct {
	DeviceID uuid.UUID
	Model    string
	Type     string
}
