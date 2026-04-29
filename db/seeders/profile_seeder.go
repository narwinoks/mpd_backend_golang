package seeders

import (
	profile "backend-app/internal/modules/master/model/profile"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

type ProfileState struct {
	ProfileID uint32
}

func SeedProfiles(db *gorm.DB) (*ProfileState, error) {
	state := &ProfileState{}

	// RS Transmedic Profile
	rsProfile := profile.Profile{
		PostalCode:     gofakeit.Zip(),
		Email:          "info@transmedic.co.id",
		Name:           "RS Transmedic",
		Profile:        "Rumah Sakit Modern",
		GovernmentName: "RS Transmedic Indonesia",
		Phone:          gofakeit.Phone(),
		Telp:           gofakeit.Phone(),
		FullAddress:    gofakeit.Address().Address,
	}
	if err := db.Create(&rsProfile).Error; err != nil {
		return nil, err
	}
	state.ProfileID = rsProfile.ID

	// Profile Detail
	rsDetail := profile.ProfileDetail{
		BaseModel:        createBaseModel(state.ProfileID, nil),
		Website:          "www.transmedic.co.id",
		Longitude:        gofakeit.Longitude(),
		Latitude:         gofakeit.Latitude(),
		RegistrationDate: time.Now(),
		Moto:             "Serving with Heart",
	}
	if err := db.Create(&rsDetail).Error; err != nil {
		return nil, err
	}

	return state, nil
}
