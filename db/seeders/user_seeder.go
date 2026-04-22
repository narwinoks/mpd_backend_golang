package seeders

import (
	"backend-app/internal/modules/master/model/user"
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) error {
	var count int64
	db.Model(&user.User{}).Count(&count)

	if count > 0 {
		fmt.Println("Users table already seeded, skipping...")
		return nil
	}

	fmt.Println("Seeding users...")

	for i := 0; i < 10; i++ {
		u := user.User{
			Username: gofakeit.Username(),
			Email:    gofakeit.Email(),
			Password: "password123", // In a real app, this should be hashed
			FullName: gofakeit.Name(),
			NIP:      gofakeit.DigitN(10),
			Role:     "staff",
			IsActive: true,
		}

		if err := db.Create(&u).Error; err != nil {
			return err
		}
	}

	fmt.Println("Users seeded successfully.")
	return nil
}
