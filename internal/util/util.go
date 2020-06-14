package util

import (
	"time"

	"github.com/google/uuid"
	"github.com/goombaio/namegenerator"
)

type RequestId string
type UUID string

func GenerateID() UUID {
	return UUID(uuid.New().String())
}

func GenerateDisplayName() string {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	return nameGenerator.Generate()
}