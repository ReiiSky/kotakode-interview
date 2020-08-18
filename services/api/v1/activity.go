package v1

import (
	"github.com/Satssuki/Go-Service-Boilerplate/models"
	"github.com/Satssuki/Go-Service-Boilerplate/services/api/validation"
)

// ActivityService struct that wrapper the user model api
type ActivityService struct {
	Activity models.Activity
}

// CreateSingleActivityService just a shorthand create object from struct
func CreateSingleActivityService() ActivityService {
	return ActivityService{models.Activity{}}
}

// Insert implementation of function in base interface
func (activity *ActivityService) Insert() error {
	currentActivity := &activity.Activity
	err := validation.ValidateSingleActivity(currentActivity)
	if err == nil {
		err = currentActivity.GetCollection().Create(currentActivity)
	}
	return err
}
