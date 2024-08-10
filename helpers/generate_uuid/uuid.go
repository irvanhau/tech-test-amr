package generate_uuid

import uuid2 "github.com/google/uuid"

type GenerateUUIDInterface interface {
	GenerateUUID() uuid2.UUID
}

type GenerateUUID struct {
}

func InitUUID() GenerateUUIDInterface {
	return &GenerateUUID{}
}

func (uuid *GenerateUUID) GenerateUUID() uuid2.UUID {
	return uuid2.New()
}
