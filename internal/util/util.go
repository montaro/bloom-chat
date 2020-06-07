package util

import "github.com/google/uuid"

type RequestId string
type UUID string

func GenerateID() UUID {
	return UUID(uuid.New().String())
}
