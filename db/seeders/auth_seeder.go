package seeders

import (
	auth "backend-app/internal/modules/auth/models"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

func SeedAuth(db *gorm.DB, profileID uint32, employees *EmployeeState) error {
	winarnoID := employees.WinarnoID

	// Roles
	roles := []string{"Super Admin", "Dokter", "Kasir"}
	roleIDs := make(map[string]uint32)
	for _, r := range roles {
		m := auth.Role{
			BaseModel: createBaseModel(profileID, &winarnoID),
			Role:      r,
		}
		if err := db.Create(&m).Error; err != nil {
			return err
		}
		roleIDs[r] = m.ID
	}

	// AppModules
	modules := []struct{ Code, Name string }{
		{"HRD", "Human Resource"},
		{"OUT", "Outpatient"},
		{"BIL", "Billing"},
	}
	modIDs := make(map[string]uint32)
	for _, mod := range modules {
		m := auth.AppModule{
			BaseModel: createBaseModel(profileID, &winarnoID),
			Code:      mod.Code,
			Name:      mod.Name,
			Category:  "Core",
			SortOrder: 1,
		}
		if err := db.Create(&m).Error; err != nil {
			return err
		}
		modIDs[mod.Code] = m.ID
	}

	// AppMenus (Hierarchical)
	for _, modID := range modIDs {
		parent := auth.AppMenu{
			BaseModel:   createBaseModel(profileID, &winarnoID),
			AppModuleID: modID,
			Code:        gofakeit.LetterN(5),
			Name:        "Dashboard",
			Path:        "/dashboard",
			SortOrder:   1,
		}
		db.Create(&parent)
		
		child := auth.AppMenu{
			BaseModel:   createBaseModel(profileID, &winarnoID),
			AppModuleID: modID,
			ParentID:    &parent.ID,
			Code:        gofakeit.LetterN(5),
			Name:        "Reports",
			Path:        "/reports",
			SortOrder:   2,
		}
		db.Create(&child)
	}

	// Users
	users := []struct {
		EmpID    uint32
		Username string
		RoleID   uint32
	}{
		{employees.WinarnoID, "narno_dev", roleIDs["Super Admin"]},
		{employees.BudiID, "dr_budi", roleIDs["Dokter"]},
	}
	for _, u := range users {
		m := auth.User{
			BaseModel:  createBaseModel(profileID, &winarnoID),
			RoleID:     u.RoleID,
			EmployeeID: &u.EmpID,
			Username:   u.Username,
			Email:      gofakeit.Email(),
			Password:   gofakeit.Password(true, true, true, false, false, 12),
		}
		if err := db.Create(&m).Error; err != nil {
			return err
		}
	}

	// RoleModules (Pivot)
	// Super Admin to all
	for _, modID := range modIDs {
		m := auth.RoleModule{
			BaseModel: createBaseModel(profileID, &winarnoID),
			RoleID:    roleIDs["Super Admin"],
			ModulesID: modID,
		}
		db.Create(&m)
	}
	// Dokter to Outpatient
	db.Create(&auth.RoleModule{
		BaseModel: createBaseModel(profileID, &winarnoID),
		RoleID:    roleIDs["Dokter"],
		ModulesID: modIDs["OUT"],
	})

	return nil
}
