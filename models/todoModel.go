package models

import (
	"devcode/entities"

	"errors"
	"fmt"
	"strings"
)

type Todo entities.Todo

func (data *Todo) ValidateTodo(action string) error {
	switch strings.ToLower(action) {
	case "create":
		if data.ActivityGroupId == "" {
			return errors.New("ActivityGroupId cannot be null")
		}
		if data.Title == "" {
			return errors.New("Title cannot be null")
		}
		if data.Priority == "" {
			return errors.New("Priority cannot be null")
		}
		if data.IsActive == "" {
			return errors.New("IsActive cannot be null")
		}
		return nil
	default:
		if data.ActivityGroupId == "" {
			return errors.New("ActivityGroupId cannot be null")
		}
		if data.Title == "" {
			return errors.New("Title cannot be null")
		}
		if data.Priority == "" {
			return errors.New("Priority cannot be null")
		}
		if data.IsActive == "" {
			return errors.New("IsActive cannot be null")
		}
		return nil
	}
}

func (data *Todo) M_CreateTodo() (*Todo, error) {

	err := db.Debug().Create(&data).Error

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return data, nil
}

func M_GetAllTodo(activityGroupId string) (*[]Todo, error) {

	listData := Todo{}
	data := make([]Todo, 0)

	where := "WHERE deleted_at IS NULL"

	if activityGroupId != "" {
		where += " AND activity_group_id = " + activityGroupId + ""
	}

	rows, err := GetDB().Debug().Raw(`SELECT * FROM todos ` + where + ``).Rows()

	for rows.Next() {

		rows.Scan(&listData.Id, &listData.ActivityGroupId, &listData.Title, &listData.IsActive, &listData.Priority, &listData.CreatedAt, &listData.UpdatedAt, &listData.DeletedAt)

		if err != nil {
			return nil, err
		}

		data = append(data, listData)
	}

	return &data, nil
}

func M_GetOneTodo(todoId string) (*Todo, error) {

	data := Todo{}

	err := db.Where("id = ?", todoId).Find(&data).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (data *Todo) M_UpdateTodo(todoId string) (*Todo, error) {

	err := db.Debug().Model(Todo{}).Where("id = ?", todoId).Update(&data).Error
	if err != nil {
		return nil, err
	}

	err = db.Where("id = ?", todoId).Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func M_DeleteTodo(todoId string) (*Todo, error) {

	err := db.Debug().Where("id = ?", todoId).Delete(&Todo{}).Error

	if err != nil {
		return nil, err
	}

	return nil, nil
}
