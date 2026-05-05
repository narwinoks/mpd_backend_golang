package seeders

import (
	general "backend-app/internal/modules/master/model/general"
	job "backend-app/internal/modules/master/model/job"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

type MasterState struct {
	ReligionIDs  map[string]uint32
	GenderIDs    map[string]uint32
	MaritalIDs   map[string]uint32
	EmpStatusIDs map[string]uint32
	JobCatIDs    map[string]uint32
	JobTitleIDs  map[string]uint32
	PositionIDs  map[string]uint32
}

func SeedMasters(db *gorm.DB, profileID uint32) (*MasterState, error) {
	state := &MasterState{
		ReligionIDs:  make(map[string]uint32),
		GenderIDs:    make(map[string]uint32),
		MaritalIDs:   make(map[string]uint32),
		EmpStatusIDs: make(map[string]uint32),
		JobCatIDs:    make(map[string]uint32),
		JobTitleIDs:  make(map[string]uint32),
		PositionIDs:  make(map[string]uint32),
	}

	// Religions
	religions := []string{"Islam", "Kristen", "Katolik", "Hindu", "Buddha", "Khonghucu"}
	for _, r := range religions {
		m := general.Religion{
			BaseModel: createBaseModel(profileID, nil),
			Religion:  r,
		}
		if err := db.Create(&m).Error; err != nil {
			return nil, err
		}
		state.ReligionIDs[r] = m.ID
	}

	// Genders
	genders := []string{"Laki-laki", "Perempuan", "Tidak Diketahui"}
	for _, g := range genders {
		m := general.Gender{
			BaseModel: createBaseModel(profileID, nil),
			Gender:    g,
		}
		if err := db.Create(&m).Error; err != nil {
			return nil, err
		}
		state.GenderIDs[g] = m.ID
	}

	// Marital Status
	maritals := []string{"Belum Kawin", "Kawin", "Cerai Hidup", "Cerai Mati"}
	for _, s := range maritals {
		m := general.MaritalStatus{
			BaseModel:      createBaseModel(profileID, nil),
			MaterialStatus: s,
		}
		if err := db.Create(&m).Error; err != nil {
			return nil, err
		}
		state.MaritalIDs[s] = m.ID
	}

	// Employment Status
	empStatuses := []struct{ Code, Status string }{
		{"TETAP", "TETAP"},
		{"KONTRAK", "KONTRAK"},
		{"MITRA", "MITRA"},
	}
	for _, s := range empStatuses {
		m := job.EmploymentStatus{
			BaseModel:      createBaseModel(profileID, nil),
			Code:           s.Code,
			EmployeeStatus: s.Status,
		}
		if err := db.Create(&m).Error; err != nil {
			return nil, err
		}
		state.EmpStatusIDs[s.Code] = m.ID
	}

	// Job Categories
	jobCats := []struct{ Code, Cat string }{
		{"MED", "Tenaga Medis"},
		{"PAR", "Tenaga Paramedis"},
		{"ADM", "Tenaga Administrasi"},
	}
	for _, c := range jobCats {
		m := job.JobCategory{
			BaseModel:   createBaseModel(profileID, nil),
			Code:        c.Code,
			JobCategory: c.Cat,
		}
		if err := db.Create(&m).Error; err != nil {
			return nil, err
		}
		state.JobCatIDs[c.Code] = m.ID
	}

	// Job Titles
	jobTitles := []struct {
		Title string
		CatID uint32
	}{
		{"Dokter Umum", state.JobCatIDs["MED"]},
		{"Dokter Spesialis Bedah", state.JobCatIDs["MED"]},
	}
	for _, t := range jobTitles {
		m := job.JobTitle{
			BaseModel:     createBaseModel(profileID, nil),
			JobCategoryID: t.CatID,
			Code:          gofakeit.LetterN(5),
			JobTitle:      t.Title,
		}
		if err := db.Create(&m).Error; err != nil {
			return nil, err
		}
		state.JobTitleIDs[t.Title] = m.ID
	}

	// Positions
	positions := []string{"Direktur Utama", "Kepala Ruangan", "Staf Pelaksana"}
	for _, p := range positions {
		m := job.Position{
			BaseModel: createBaseModel(profileID, nil),
			Position:  p,
		}
		if err := db.Create(&m).Error; err != nil {
			return nil, err
		}
		state.PositionIDs[p] = m.ID
	}

	return state, nil
}
