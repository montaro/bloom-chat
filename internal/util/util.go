package util

import "github.com/google/uuid"

type UUID string

func GenerateID() UUID {
	return UUID(uuid.New().String())
}
