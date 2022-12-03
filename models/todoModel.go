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
		if data.ActivityGroupId <= 0 {
			return errors.New("activity_group_id cannot be null")
		}
		if data.Title == "" {
			return errors.New("title cannot be null")
		}
		return nil
	default:
		if data.ActivityGroupId <= 0 {
			return errors.New("activity_group_id cannot be null")
		}
		if data.Title == "" {
			return errors.New("title cannot be null")
		}
		return nil
	}
}

func (data *Todo) M_CreateTodo() (*Todo, error) {

	data.IsActive = true
	data.Priority = "very-high"

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
			fmt.Println(err.Error())
			return nil, err
		}

		data = append(data, listData)
	}

	return &data, nil
}

func M_GetOneTodo(todoId string) (*Todo, error) {

	data := Todo{}

	err := db.Debug().Raw(`SELECT * FROM todos WHERE id = ?`, todoId).Scan(&data).Error

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &data, nil
}

func (data *Todo) M_UpdateTodo(todoId string) (*Todo, error) {

	err := db.Debug().Exec("UPDATE todos SET activity_group_id = COALESCE(NULLIF(?, null), activity_group_id), title = COALESCE(NULLIF(?,''), title), is_active = COALESCE(NULLIF(?, null), is_active), priority = COALESCE(NULLIF(?,''), priority) WHERE id = ?", data.ActivityGroupId, data.Title, data.IsActive, data.Priority, todoId).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	err = db.Debug().Raw(`SELECT * FROM todos WHERE id = ?`, todoId).Scan(&data).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return data, nil
}

func M_DeleteTodo(todoId string) (*Todo, error) {

	data := Todo{}

	err := db.Debug().Raw(`SELECT * FROM todos WHERE id = ?`, todoId).Scan(&data).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	err = db.Debug().Where("id = ?", todoId).Delete(&Todo{}).Error

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return nil, nil
}
