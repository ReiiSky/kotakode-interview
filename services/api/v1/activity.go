package v1

import (
	"fmt"

	"github.com/Kamva/mgm/v3"
	"github.com/Satssuki/Go-Service-Boilerplate/models"
	"github.com/Satssuki/Go-Service-Boilerplate/services/api/validation"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ActivityService struct that wrapper the user model api
type ActivityService struct {
	Activity models.Activity
}

// ActivityServices struct that wrapper the user model api
type ActivityServices struct {
	Activities []models.Activity `json:"list"`
}

// CreateSingleActivityService just a shorthand create object from struct
func CreateSingleActivityService() ActivityService {
	return ActivityService{models.Activity{}}
}

// CreateBulkActivityService just a shorthand create object from struct
func CreateBulkActivityService() ActivityServices {
	return ActivityServices{[]models.Activity{}}
}

// Insert implementation of function in base interface
func (activity *ActivityServices) Insert() error {
	var err error
	if len(activity.Activities) == 0 {
		err = fmt.Errorf("Please insert document length more than 0")
	} else {
		listInterface := []interface{}{}
		for _, y := range activity.Activities {
			listInterface = append(listInterface, y)
		}
		_, Err := activity.Activities[0].GetCollection().InsertMany(mgm.Ctx(), listInterface)
		err = Err
	}
	return err
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

// Find implementation of function in base interface
func (activity *ActivityServices) Find(query bson.M) (*ActivityServices, error) {
	act := CreateSingleActivityService()
	err := act.Activity.GetCollection().SimpleFind(&activity.Activities, query)
	return activity, err
}

// FindByID implementation of function in base interface
func (activity *ActivityService) FindByID(id string) (*ActivityService, error) {
	act := CreateSingleActivityService()
	err := act.Activity.GetCollection().FindByID(id, &activity.Activity)
	return activity, err
}

// DeleteOne implementation of function in base interface
func (activity *ActivityService) DeleteOne(id string) int {
	activity.Activity.Group = "nil"
	hex, err := primitive.ObjectIDFromHex(id)
	if err == nil {
		result, err := activity.Activity.GetCollection().DeleteOne(mgm.Ctx(), bson.M{
			"_id": hex,
		})
		if err == nil {
			return int(result.DeletedCount)
		}
	}
	return 0
}
