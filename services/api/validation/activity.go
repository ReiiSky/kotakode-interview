package validation

import (
	"github.com/Satssuki/Go-Service-Boilerplate/models"
	"github.com/softbrewery/gojoi/pkg/joi"
)

var activitySchema = joi.Struct().Keys(joi.StructKeys{
	"Title":    joi.String().Max(64).Min(6),
	"Priority": joi.Int().Min(0).Max(3),
	"Group":    joi.String().Max(16).Min(4),
	"Detail":   joi.String().Min(6),
})

func ValidateSingleActivity(activity *models.Activity) error {
	return joi.Validate(*activity, schema)
}
