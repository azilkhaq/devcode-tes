package controllers

import (
	"devcode/helpers"
	"devcode/models"

	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func C_CreateActivity(w http.ResponseWriter, r *http.Request) {

	data := &models.Activity{}

	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		resp := helpers.Message("Bad Request", err.Error())
		resp["data"] = map[string]interface{}{}
		helpers.Response(w, http.StatusBadRequest, resp)
		return
	}

	err = data.Validate("create")
	if err != nil {
		resp := helpers.Message("Bad Request", err.Error())
		resp["data"] = map[string]interface{}{}
		helpers.Response(w, http.StatusBadRequest, resp)
		return
	}

	result, err := data.M_CreateActivity()
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

func C_GetAllActivity(w http.ResponseWriter, r *http.Request) {

	result, err := models.M_GetAllActivity()

	if err != nil {
		resp := helpers.Message("Bad Request", err.Error())
		resp["data"] = map[string]interface{}{}
		helpers.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helpers.Message("Success", "Success")
	resp["data"] = result
	helpers.Response(w, http.StatusOK, resp)
}

func C_GetOneActivity(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	activityId := vars["id"]

	result, err := models.M_GetOneActivity(activityId)
	if err != nil {
		resp := map[string]interface{}{}
		if strings.Contains(err.Error(), "record not found") {
			resp = helpers.Message("Not Found", "Activity with ID " + activityId + " Not Found")
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

func C_UpdateActivity(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	activityId := vars["id"]

	data := &models.Activity{}

	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		resp := helpers.Message("Bad Request", err.Error())
		resp["data"] = map[string]interface{}{}
		helpers.Response(w, http.StatusBadRequest, resp)
		return
	}

	result, err := data.M_UpdateActivity(activityId)
	if err != nil {
		resp := map[string]interface{}{}
		if strings.Contains(err.Error(), "record not found") {
			resp = helpers.Message("Not Found", "Activity with ID " + activityId + " Not Found")
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

func C_DeleteActivity(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	activityId := vars["id"]

	_, err := models.M_DeleteActivity(activityId)
	if err != nil {
		resp := map[string]interface{}{}
		if strings.Contains(err.Error(), "record not found") {
			resp = helpers.Message("Not Found", "Activity with ID " + activityId + " Not Found")
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
