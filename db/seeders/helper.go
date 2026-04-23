package seeders

import (
	base "backend-app/internal/base/models"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func createBaseModel(profileID uint32, creatorID *uint32) base.BaseModel {
	return base.BaseModel{
		UUID:         gofakeit.UUID(),
		IsActive:     true,
		ProfileID:    &profileID,
		ExternalCode: gofakeit.LetterN(10),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		CreatedBy:    creatorID,
		UpdatedBy:    creatorID,
	}
}

func uint32Ptr(v uint32) *uint32 {
	return &v
}
