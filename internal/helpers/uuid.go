package helpers

import "github.com/google/uuid"

func GetUUID() string {
	return uuid.NewString()
}
