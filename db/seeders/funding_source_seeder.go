package seeders

import (
	model "backend-app/internal/modules/master/model/funding_source"

	"gorm.io/gorm"
)

func SeedFundingSources(db *gorm.DB, profileID uint32) error {
	sources := []string{"ANGGARAN", "DONASI", "GRATIS"}

	for _, s := range sources {
		m := model.FundingSource{
			BaseModel:     createBaseModel(profileID, nil),
			FundingSource: s,
		}
		if err := db.Where("funding_source = ?", s).FirstOrCreate(&m).Error; err != nil {
			return err
		}
	}

	return nil
}
