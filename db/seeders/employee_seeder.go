package seeders

import (
	employee "backend-app/internal/modules/master/model/employee"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

type EmployeeState struct {
	WinarnoID uint32
	BudiID    uint32
}

func SeedEmployees(db *gorm.DB, profileID uint32, masters *MasterState) (*EmployeeState, error) {
	state := &EmployeeState{}

	// Winarno (First Employee, CreatedBy: nil)
	winarno := employee.Employee{
		BaseModel:          createBaseModel(profileID, nil),
		ReligionID:         masters.ReligionIDs["Islam"],
		GenderID:           masters.GenderIDs["Laki-laki"],
		JobTitleID:         masters.JobTitleIDs["Dokter Umum"],
		EmploymentStatusID: masters.EmpStatusIDs["TETAP"],
		FullName:           "Winarno",
		IdentityNumber:     gofakeit.DigitN(16),
		NIP:                "195001",
		NPWP:               gofakeit.DigitN(15),
		BirthPlace:         "Bandung",
		BirthDate:          time.Date(1980, 1, 1, 0, 0, 0, 0, time.Local),
	}
	if err := db.Create(&winarno).Error; err != nil {
		return nil, err
	}
	state.WinarnoID = winarno.ID

	// Update Winarno's audit fields to point to himself (optional but common)
	db.Model(&winarno).Updates(map[string]interface{}{
		"created_by": state.WinarnoID,
		"updated_by": state.WinarnoID,
	})

	// dr. Budi Santoso, Sp.B
	budi := employee.Employee{
		BaseModel:          createBaseModel(profileID, &state.WinarnoID),
		ReligionID:         masters.ReligionIDs["Islam"],
		GenderID:           masters.GenderIDs["Laki-laki"],
		JobTitleID:         masters.JobTitleIDs["Dokter Spesialis Bedah"],
		EmploymentStatusID: masters.EmpStatusIDs["TETAP"],
		FullName:           "dr. Budi Santoso, Sp.B",
		IdentityNumber:     gofakeit.DigitN(16),
		NIP:                "195002",
		NPWP:               gofakeit.DigitN(15),
		BirthPlace:         "Jakarta",
		BirthDate:          time.Date(1975, 5, 20, 0, 0, 0, 0, time.Local),
	}
	if err := db.Create(&budi).Error; err != nil {
		return nil, err
	}
	state.BudiID = budi.ID

	// Employee Details
	emps := []uint32{state.WinarnoID, state.BudiID}
	for _, id := range emps {
		detail := employee.EmployeeDetail{
			BaseModel:            createBaseModel(profileID, &state.WinarnoID),
			EmployeeID:           &id,
			MaritalStatusID:      masters.MaritalIDs["Kawin"],
			FunctionalPositionID: masters.PositionIDs["Staf Pelaksana"],
			StructuralPositionID: masters.PositionIDs["Staf Pelaksana"],
			JoinDate:             time.Now().AddDate(-1, 0, 0),
		}
		if err := db.Create(&detail).Error; err != nil {
			return nil, err
		}
	}

	return state, nil
}
