package utils

import (
	"time"

	"fmt"

	"github.com/google/uuid"
)

func ParseJamString(jamStr string) (time.Time, error) {
	layout := "15:04"
	return time.Parse(layout, jamStr)
}

func ParseUUID(idStr string) (uuid.UUID, error) {
	uid, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s tidak valid: %w", idStr, err)
	}
	return uid, nil
}

