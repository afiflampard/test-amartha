package helper

import (
	"fmt"

	"github.com/google/uuid"
)

func CastingToUID(value interface{}) (uuid.UUID, error) {
	str, ok := value.(string)
	if !ok {
		return uuid.UUID{}, fmt.Errorf("value is not a string")
	}

	return uuid.Parse(str)
}
