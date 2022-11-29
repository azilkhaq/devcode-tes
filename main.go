package main

import (
	"devcode/controllers"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println(err.Error())
	}

	router := mux.NewRouter()

	// Activity
	router.HandleFunc("/activity-groups", controllers.C_CreateActivity).Methods("POST")
	router.HandleFunc("/activity-groups", controllers.C_GetAllActivity).Methods("GET")
	router.HandleFunc("/activity-groups/{id}", controllers.C_GetOneActivity).Methods("GET")
	router.HandleFunc("/activity-groups/{id}", controllers.C_UpdateActivity).Methods("PATCH")
	router.HandleFunc("/activity-groups/{id}", controllers.C_DeleteActivity).Methods("DELETE")

	// Todo
	router.HandleFunc("/todo-items", controllers.C_CreateTodo).Methods("POST")
	router.HandleFunc("/todo-items", controllers.C_GetAllTodo).Methods("GET")
	router.HandleFunc("/todo-items/{id}", controllers.C_GetOneTodo).Methods("GET")
	router.HandleFunc("/todo-items/{id}", controllers.C_UpdateTodo).Methods("PATCH")
	router.HandleFunc("/todo-items/{id}", controllers.C_DeleteTodo).Methods("DELETE")

	port := os.Getenv("PORT")
	handler := router
	server := new(http.Server)
	server.Handler = handler
	server.Addr = ":" + port
	fmt.Println("Starting server at", server.Addr)
	server.ListenAndServe()
}
