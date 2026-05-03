package validations

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type DBValidator struct {
	DB *gorm.DB
}

func (v *DBValidator) ValidateUnique(fl validator.FieldLevel) bool {
	param := fl.Param()

	params := strings.Split(param, ".")
	if len(params) != 2 {
		return false
	}

	tableName := params[0]
	columnName := params[1]

	value := fl.Field().String()
	if value == "" {
		return true
	}

	var count int64
	v.DB.Table(tableName).Where(columnName+" = ?", value).Count(&count)

	return count == 0
}
