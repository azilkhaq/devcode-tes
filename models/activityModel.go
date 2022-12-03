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

	err := db.Create(&data).Error

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return data, nil
}

func M_GetAllActivity() (*[]Activity, error) {

	data := []Activity{}
	
	err := db.Find(&data).Error

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &data, nil
}

func M_GetOneActivity(activityId string) (*Activity, error) {

	data := Activity{}

	err := db.Raw(`SELECT * FROM activities WHERE id = ?`, activityId).Scan(&data).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &data, nil
}

func (data *Activity) M_UpdateActivity(activityId string) (*Activity, error) {

	err := db.Exec("UPDATE activities SET email = COALESCE(NULLIF(?,''), email), title = COALESCE(NULLIF(?,''), title) WHERE id = ?", data.Email, data.Title, activityId).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	err = db.Raw(`SELECT * FROM activities WHERE id = ?`, activityId).Scan(&data).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return data, nil
}

func M_DeleteActivity(activityId string) (*Activity, error) {

	data := Activity{}

	err := db.Raw(`SELECT * FROM activities WHERE id = ?`, activityId).Scan(&data).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	err = db.Where("id = ?", activityId).Delete(&Activity{}).Error

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return nil, nil
}