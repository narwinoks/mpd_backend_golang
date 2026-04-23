package seeders

import (
	location "backend-app/internal/modules/master/model/location"

	"gorm.io/gorm"
)

type LocationState struct {
	ProvinceIDs    map[string]uint32
	CityIDs        map[string]uint32
	SubdistrictIDs map[string]uint32
}

func SeedLocations(db *gorm.DB) (*LocationState, error) {
	state := &LocationState{
		ProvinceIDs:    make(map[string]uint32),
		CityIDs:        make(map[string]uint32),
		SubdistrictIDs: make(map[string]uint32),
	}

	// Provinces
	provinces := []string{"Jawa Barat", "DKI Jakarta", "Jawa Tengah"}
	for _, p := range provinces {
		m := location.Province{
			BaseModel: createBaseModel(1, nil),
			Province:  p,
		}
		if err := db.Create(&m).Error; err != nil {
			return nil, err
		}
		state.ProvinceIDs[p] = m.ID
	}

	// Cities
	cities := []struct {
		Name       string
		ProvinceID uint32
	}{
		{"Kota Bandung", state.ProvinceIDs["Jawa Barat"]},
		{"Kab. Bandung", state.ProvinceIDs["Jawa Barat"]},
		{"Jakarta Selatan", state.ProvinceIDs["DKI Jakarta"]},
	}
	for _, c := range cities {
		m := location.City{
			BaseModel:  createBaseModel(1, nil),
			ProvinceID: c.ProvinceID,
			City:       c.Name,
		}
		if err := db.Create(&m).Error; err != nil {
			return nil, err
		}
		state.CityIDs[c.Name] = m.ID
	}

	// Subdistricts
	subdistricts := []struct {
		Name   string
		CityID uint32
	}{
		{"Cicendo", state.CityIDs["Kota Bandung"]},
		{"Coblong", state.CityIDs["Kota Bandung"]},
	}
	for _, s := range subdistricts {
		m := location.Subdistrict{
			BaseModel:   createBaseModel(1, nil),
			CityID:      s.CityID,
			ProvinceID:  state.ProvinceIDs["Jawa Barat"],
			Subdistrict: s.Name,
		}
		if err := db.Create(&m).Error; err != nil {
			return nil, err
		}
		state.SubdistrictIDs[s.Name] = m.ID
	}

	// Villages
	villages := []struct {
		Name          string
		SubdistrictID uint32
	}{
		{"Pasir Kaliki", state.SubdistrictIDs["Cicendo"]},
		{"Arjuna", state.SubdistrictIDs["Cicendo"]},
	}
	for _, v := range villages {
		m := location.Village{
			BaseModel:     createBaseModel(1, nil),
			SubdistrictID: v.SubdistrictID,
			Village:       v.Name,
		}
		if err := db.Create(&m).Error; err != nil {
			return nil, err
		}
	}

	return state, nil
}
