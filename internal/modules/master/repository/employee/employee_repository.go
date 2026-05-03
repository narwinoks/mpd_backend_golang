package employee

import (
	model "backend-app/internal/modules/master/model/employee"
	"backend-app/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type employeeRepositoryImpl struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepositoryImpl{db: db}
}

type EmployeeWithCount struct {
	model.Employee
	TotalCount int64 `gorm:"column:total_count"`
}

func (r *employeeRepositoryImpl) FindAll(ctx context.Context, req pagination.BaseRequest) ([]model.Employee, int64, error) {
	var results []EmployeeWithCount
	var items []model.Employee
	var total int64

	err := r.db.WithContext(ctx).Model(&model.Employee{}).
		Preload("Gender", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "gender")
		}).
		Preload("JobTitle", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "job_title")
		}).
		Preload("EmployeeStatus", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "employee_status")
		}).
		Preload("Religion", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "religion")
		}).
		Scopes(pagination.PaginateScope(req)).
		Scopes(req.SearchScope("full_name")).
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	if len(results) > 0 {
		total = results[0].TotalCount
		for _, res := range results {
			items = append(items, res.Employee)
		}
	}

	return items, total, nil
}

func (r *employeeRepositoryImpl) FindByID(ctx context.Context, id uint32) (*model.Employee, error) {
	var employee model.Employee
	err := r.db.WithContext(ctx).
		Preload("Gender", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "gender")
		}).
		Preload("JobTitle", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "job_title")
		}).
		Preload("EmployeeStatus", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "employee_status")
		}).
		Preload("Religion", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "religion")
		}).
		Preload("Detail").
		Preload("Addresses").
		Preload("Educations").
		First(&employee, id).Error
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *employeeRepositoryImpl) Create(ctx context.Context, employee *model.Employee) error {
	return r.db.WithContext(ctx).Session(&gorm.Session{FullSaveAssociations: true}).Create(employee).Error
}

func (r *employeeRepositoryImpl) Update(ctx context.Context, employee *model.Employee) error {
	return r.db.WithContext(ctx).Session(&gorm.Session{FullSaveAssociations: true}).Save(employee).Error
}

func (r *employeeRepositoryImpl) Delete(ctx context.Context, id uint32) error {
	var employee model.Employee
	if err := r.db.WithContext(ctx).First(&employee, id).Error; err != nil {
		return err
	}
	return employee.SetNonActive(r.db.WithContext(ctx).Model(&employee))
}

func (r *employeeRepositoryImpl) FindByUuid(ctx context.Context, Uuid string) (*model.Employee, error) {
	var employee model.Employee
	err := r.db.WithContext(ctx).
		Preload("Gender", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "gender")
		}).
		Preload("JobTitle", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "job_title")
		}).
		Preload("EmployeeStatus", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "employee_status")
		}).
		Preload("Religion", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "religion")
		}).
		Preload("Detail").
		Preload("Addresses").
		Preload("Educations").
		Where("uuid = ?", Uuid).First(&employee).Error
	if err != nil {
		return nil, err
	}
	return &employee, nil
}
