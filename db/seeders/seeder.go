package seeders

import (
	"fmt"

	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// 2. Profiles
		fmt.Println("Seeding Profiles...")
		profiles, err := SeedProfiles(tx)
		if err != nil {
			return fmt.Errorf("failed to seed profiles: %w", err)
		}
		// 1. Locations
		fmt.Println("Seeding Locations...")
		_, err = SeedLocations(tx)
		if err != nil {
			return err
		}
		if err != nil {
			return fmt.Errorf("failed to seed locations: %w", err)
		}

		// 3. Masters
		fmt.Println("Seeding Masters...")
		masters, err := SeedMasters(tx, profiles.ProfileID)
		if err != nil {
			return fmt.Errorf("failed to seed masters: %w", err)
		}

		// 4. Employees
		fmt.Println("Seeding Employees...")
		employees, err := SeedEmployees(tx, profiles.ProfileID, masters)
		if err != nil {
			return fmt.Errorf("failed to seed employees: %w", err)
		}

		// 5. Auth & RBAC
		fmt.Println("Seeding Auth & RBAC...")
		if err := SeedAuth(tx, profiles.ProfileID, employees); err != nil {
			return fmt.Errorf("failed to seed auth: %w", err)
		}

		fmt.Println("Seeding Menu...")
		err = SeedMenus(tx, profiles.ProfileID, employees)
		if err != nil {
			return fmt.Errorf("failed to seed menu: %w", err)
		}

		return nil
	})
}
