package controllers

import (
	"devcode/helpers"
	"devcode/models"

	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func C_CreateTodo(w http.ResponseWriter, r *http.Request) {

	data := &models.Todo{}

	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		resp := helpers.Message("Bad Request", err.Error())
		resp["data"] = map[string]interface{}{}
		helpers.Response(w, http.StatusBadRequest, resp)
		return
	}

	err = data.ValidateTodo("create")
	if err != nil {
		resp := helpers.Message("Bad Request", err.Error())
		resp["data"] = map[string]interface{}{}
		helpers.Response(w, http.StatusBadRequest, resp)
		return
	}

	result, err := data.M_CreateTodo()
	if err != nil {
		resp := helpers.Message("Bad Request", err.Error())
		resp["data"] = map[string]interface{}{}
		helpers.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helpers.Message("Success", "Success")
	resp["data"] = result
	helpers.Response(w, http.StatusCreated, resp)
}

func C_GetAllTodo(w http.ResponseWriter, r *http.Request) {

	activityGroupId := r.URL.Query().Get("activity_group_id")

	result, err := models.M_GetAllTodo(activityGroupId)

	if err != nil {
		resp := helpers.Message("Bad Request", err.Error())
		resp["data"] = map[string]interface{}{}
		helpers.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helpers.Message("Success", "Success")
	resp["data"] = result
	helpers.Response(w, http.StatusCreated, resp)
}

func C_GetOneTodo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	todoId := vars["id"]

	result, err := models.M_GetOneTodo(todoId)
	if err != nil {
		resp := map[string]interface{}{}
		if strings.Contains(err.Error(), "record not found") {
			resp = helpers.Message("Not Found", "Todo with ID " + todoId + " Not Found")
			resp["data"] = map[string]interface{}{}
			helpers.Response(w, http.StatusNotFound, resp)
		} else {
			resp = helpers.Message("Bad Request", err.Error())
			resp["data"] = map[string]interface{}{}
			helpers.Response(w, http.StatusBadRequest, resp)
		}
		return
	}

	resp := helpers.Message("Success", "Success")
	resp["data"] = result
	helpers.Response(w, http.StatusOK, resp)
}

func C_UpdateTodo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	todoId := vars["id"]

	data := &models.Todo{}

	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		resp := helpers.Message("Bad Request", err.Error())
		resp["data"] = map[string]interface{}{}
		helpers.Response(w, http.StatusBadRequest, resp)
		return
	}

	result, err := data.M_UpdateTodo(todoId)
	if err != nil {
		resp := map[string]interface{}{}
		if strings.Contains(err.Error(), "record not found") {
			resp = helpers.Message("Not Found", "Todo with ID " + todoId + " Not Found")
			resp["data"] = map[string]interface{}{}
			helpers.Response(w, http.StatusNotFound, resp)
		} else {
			resp = helpers.Message("Bad Request", err.Error())
			resp["data"] = map[string]interface{}{}
			helpers.Response(w, http.StatusBadRequest, resp)
		}
		return
	}

	resp := helpers.Message("Success", "Success")
	resp["data"] = result
	helpers.Response(w, http.StatusOK, resp)
}

func C_DeleteTodo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	todoId := vars["id"]

	_, err := models.M_DeleteTodo(todoId)
	if err != nil {
		resp := map[string]interface{}{}
		if strings.Contains(err.Error(), "record not found") {
			resp = helpers.Message("Not Found", "Todo with ID " + todoId + " Not Found")
			resp["data"] = map[string]interface{}{}
			helpers.Response(w, http.StatusNotFound, resp)
		} else {
			resp = helpers.Message("Bad Request", err.Error())
			resp["data"] = map[string]interface{}{}
			helpers.Response(w, http.StatusBadRequest, resp)
		}
		return
	}

	resp := helpers.Message("Success", "Success")
	resp["data"] = map[string]interface{}{}
	helpers.Response(w, http.StatusOK, resp)
}
