package seeders

import (
	auth "backend-app/internal/modules/auth/models"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func SeedMenus(db *gorm.DB, profileID uint32, employees *EmployeeState) error {
	winarnoID := employees.WinarnoID
	type SubMenuData struct {
		Title string
		Href  string
	}

	type MenuData struct {
		Title    string
		Href     string
		Icon     string
		SubItems []SubMenuData
	}

	rawMenus := []MenuData{
		{
			Title: "Pendaftaran", Href: "null", Icon: "ph-bold ph-identification-card",
			SubItems: []SubMenuData{
				{Title: "Pasien Baru", Href: "/pendaftaran/baru"},
				{Title: "List Antrean", Href: "/pendaftaran/antrean"},
				{Title: "Registrasi Rawat Jalan", Href: "/pendaftaran/rawat-jalan"},
			},
		},
		{
			Title: "Rawat Jalan", Href: "/rawat-jalan", Icon: "ph-bold ph-stethoscope",
			SubItems: []SubMenuData{
				{Title: "Pemeriksaan Dokter", Href: "/rawat-jalan/pemeriksaan"},
				{Title: "Input Tindakan", Href: "/rawat-jalan/tindakan"},
				{Title: "E-Resep", Href: "/rawat-jalan/resep"},
			},
		},
		{
			Title: "Rawat Inap", Href: "/rawat-inap", Icon: "ph-bold ph-bed",
			SubItems: []SubMenuData{
				{Title: "Manajemen Bangsal", Href: "/rawat-inap/bangsal"},
				{Title: "Pindah Kamar", Href: "/rawat-inap/pindah"},
				{Title: "Monitoring Pasien", Href: "/rawat-inap/monitoring"},
			},
		},
		{
			Title: "Farmasi", Href: "/farmasi", Icon: "ph-bold ph-pill",
			SubItems: []SubMenuData{
				{Title: "Antrean Resep", Href: "/farmasi/antrean"},
				{Title: "Stok Obat", Href: "/farmasi/stok"},
				{Title: "Laporan Penjualan", Href: "/farmasi/laporan"},
			},
		},
		{
			Title: "Laboratorium", Href: "/laboratorium", Icon: "ph-bold ph-flask",
			SubItems: []SubMenuData{
				{Title: "Input Hasil Lab", Href: "/laboratorium/input"},
				{Title: "Validasi Hasil", Href: "/laboratorium/validasi"},
			},
		},
		{
			Title: "Billing", Href: "/billing", Icon: "ph-bold ph-credit-card",
			SubItems: []SubMenuData{
				{Title: "Pembayaran Kasir", Href: "/billing/kasir"},
				{Title: "Piutang Asuransi", Href: "/billing/asuransi"},
			},
		},
	}

	// 2. Buat Modul Utama (Parent tertinggi)
	mainModule := auth.AppModule{
		BaseModel: createBaseModel(profileID, &winarnoID),
		Code:      "MOD_HIS",
		Name:      "Hospital Information System",
		Category:  "Core",
		SortOrder: 1,
	}
	if err := db.Create(&mainModule).Error; err != nil {
		return err
	}

	// 3. Eksekusi Looping Data Menu
	for i, menu := range rawMenus {
		// Generate Code yang bersih, misal: "MENU_PENDAFTARAN"
		parentCode := fmt.Sprintf("MENU_%s", strings.ToUpper(strings.ReplaceAll(menu.Title, " ", "_")))

		// A. Insert Parent Menu
		parentMenu := auth.AppMenu{
			BaseModel:   createBaseModel(profileID, &winarnoID),
			AppModuleID: mainModule.ID,
			Code:        parentCode,
			Name:        menu.Title,
			Path:        menu.Href,
			Icon:        menu.Icon,
			SortOrder:   int(uint32(i + 1)),
		}

		if err := db.Create(&parentMenu).Error; err != nil {
			return err
		}

		// B. Insert Child Menus (SubItems)
		for j, sub := range menu.SubItems {
			childCode := fmt.Sprintf("%s_%s", parentCode, strings.ToUpper(strings.ReplaceAll(sub.Title, " ", "_")))

			childMenu := auth.AppMenu{
				BaseModel:   createBaseModel(profileID, &winarnoID),
				AppModuleID: mainModule.ID,
				ParentID:    &parentMenu.ID,
				Code:        childCode,
				Name:        sub.Title,
				Path:        sub.Href,
				SortOrder:   int(uint32(j + 1)),
			}

			if err := db.Create(&childMenu).Error; err != nil {
				return err
			}
		}
	}

	fmt.Println("✅ Berhasil melakukan seeding struktur Menu SIMRS!")
	return nil
}
