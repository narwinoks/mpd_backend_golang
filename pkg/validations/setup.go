package validations

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func InitGinValidator(db *gorm.DB) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

			if name == "-" {
				return ""
			}
			return name
		})
		// REGISTER CUSTOM TAG
		_ = v.RegisterValidation("is_npwp", ValidateNPWP)
		dbValidator := &DBValidator{DB: db}
		_ = v.RegisterValidation("unique", dbValidator.ValidateUnique)
	}
}
