package models

import (
	"devcode/entities"

	"fmt"
	"strings"
	"errors"
)

type Activity entities.Activity

func (data *Activity) Validate(action string) error {
	switch strings.ToLower(action) {
	case "create":
		if data.Email == "" {
			return errors.New("email cannot be null")
		}
		if data.Title == "" {
			return errors.New("title cannot be null")
		}
		return nil
	default:
		if data.Email == "" {
			return errors.New("email cannot be null")
		}
		if data.Title == "" {
			return errors.New("title cannot be null")
		}
		return nil
	}
}

func (data *Activity) M_CreateActivity() (*Activity, error) {

	err := db.Debug().Create(&data).Error

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return data, nil
}

func M_GetAllActivity() (*[]Activity, error) {

	data := []Activity{}
	
	err := db.Debug().Find(&data).Error

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &data, nil
}

func M_GetOneActivity(activityId string) (*Activity, error) {

	data := Activity{}
	
	err := db.Debug().Where("id = ?", activityId).Find(&data).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (data *Activity) M_UpdateActivity(activityId string) (*Activity, error) {

	err := db.Debug().Exec("UPDATE activities SET email = ?, title = ? WHERE id = ?", data.Email, data.Title, activityId).Error
	if err != nil {
		return nil, err
	}
	
	err = db.Where("id = ?", activityId).Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func M_DeleteActivity(activityId string) (*Activity, error) {

	err := db.Debug().Where("id = ?", activityId).Delete(&Activity{}).Error

	if err != nil {
		return nil, err
	}

	return nil, nil
}