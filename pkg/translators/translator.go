package translators

import (
	"github.com/google/uuid"
)

type Translator interface {
	//TranslateDispatch()
	TranslateIngestion(deviceID uuid.UUID, deviceType string, message []byte) ([]byte, error)
}
