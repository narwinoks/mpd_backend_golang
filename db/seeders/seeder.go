package seeders

import (
	"gorm.io/gorm"
)

type Seeder struct {
	db *gorm.DB
}

func NewSeeder(db *gorm.DB) *Seeder {
	return &Seeder{db: db}
}

func (s *Seeder) Run() error {
	if err := SeedUser(s.db); err != nil {
		return err
	}
	// Add other seeders here as needed

	return nil
}
