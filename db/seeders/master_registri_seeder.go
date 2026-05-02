package seeders

import (
	model "backend-app/internal/modules/master/model/registry"

	"gorm.io/gorm"
)

type RegistrySeed struct {
	Name      string
	Path      string
	Icon      string
	SortOrder int
	Children  []RegistrySeed
}

func SeedRegistries(db *gorm.DB) error {
	seeds := []RegistrySeed{
		{
			Name: "Matser User", Icon: "fas fa-users", SortOrder: 1,
			Children: []RegistrySeed{
				{Name: "Permission Aplikasi", Path: "/master/users/permissions", SortOrder: 1},
				{Name: "Kelompok User", Path: "/master/users/roles", SortOrder: 1},
				{Name: "Master User", Path: "/master/users", SortOrder: 1},
			},
		},
		{
			Name: "Master Farmasi", Icon: "ph-bold ph-pill", SortOrder: 2,
		},
		{
			Name: "Master Genaral", Icon: "fas fa-file-pdf", SortOrder: 2,
			Children: []RegistrySeed{
				{Name: "List Master Data", Path: "/master/registries", Icon: "fas fa-file-pdf", SortOrder: 1},
				{Name: "Infomation", Icon: "ph-fill ph-heartbeat", SortOrder: 5},
				{Name: "Agama", Path: "/master/general/religions", SortOrder: 2},
				{Name: "Jenis Kelamin", Path: "/master/general/genders", SortOrder: 3},
			},
		},
		{
			Name: "Pegawai", Icon: "fa fa-user", SortOrder: 4,
			Children: []RegistrySeed{
				{Name: "Pegawai", Path: "/master/employees", SortOrder: 1},
				{Name: "Ketegory Pegawai", Path: "/master/employee/job-category", SortOrder: 1},
				{Name: "Jenis Pegawai", Path: "/master/employee/job-titles", SortOrder: 2},
			},
		},
	}

	// 3. Eksekusi looping Parent. headID diisi nil karena ini adalah Menu Utama
	for _, seed := range seeds {
		if err := processSeed(db, seed, nil); err != nil {
			return err
		}
	}

	return nil
}

func processSeed(db *gorm.DB, seed RegistrySeed, headID *uint32) error {
	profileID := 1

	// 1. Siapkan model GORM dengan data lengkap (dipakai jika harus Create baru)
	registry := model.Registry{
		Name:      seed.Name,
		Path:      seed.Path,
		Icon:      seed.Icon,
		HeadID:    headID,
		BaseModel: createBaseModel(uint32(profileID), nil),
		SortOrder: seed.SortOrder,
	}

	// 2. Eksekusi Upsert menggunakan FirstOrCreate
	// Kita jadikan Name dan Path sebagai identifier (kunci pencarian)
	err := db.Where("name = ? AND path = ?", seed.Name, seed.Path).
		Assign(map[string]interface{}{
			// Assign berisi kolom-kolom yang akan di-update JIKA data sudah ditemukan
			"icon":       seed.Icon,
			"head_id":    headID,
			"sort_order": seed.SortOrder,
			"is_active":  true,
		}).
		FirstOrCreate(&registry).Error

	if err != nil {
		return err
	}

	// 3. Looping ke anak-anaknya secara rekursif
	for _, child := range seed.Children {
		// Ajaibnya FirstOrCreate: struct 'registry' otomatis terisi ID asli dari database,
		// baik dari hasil Create maupun dari hasil Find.
		if err := processSeed(db, child, &registry.ID); err != nil {
			return err
		}
	}

	return nil
}
