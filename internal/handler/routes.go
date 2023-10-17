package handler

import (
	"restApi/internal/service"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	service service.Service
}

func NewUserHandler(service service.Service) *UserHandler {
	return &UserHandler{
		service: service,
	}
}
func (h *UserHandler) InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users", h.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", h.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", h.GetUser).Methods("GET")

	return router
}
