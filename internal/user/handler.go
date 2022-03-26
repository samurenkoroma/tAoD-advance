package user

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"tAoD-advance/internal/handlers"
	"tAoD-advance/pkg/logging"
)

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{logger: logger}

}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.GET(userURL, h.GetUserById)
	router.POST(usersURL, h.CreateUser)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PatchUpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("list of users")
	w.WriteHeader(200)
	w.Write([]byte("list of users"))
}

func (h *handler) GetUserById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("user by id")
	w.WriteHeader(200)
	w.Write([]byte("user by id"))
}
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("create user")
	w.WriteHeader(201)
	w.Write([]byte("create user"))
}
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("update user")
	w.WriteHeader(204)
	w.Write([]byte("update user"))

}
func (h *handler) PatchUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("patch update user")
	w.WriteHeader(204)
	w.Write([]byte("patch update user"))

}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, para httprouter.Params) {
	h.logger.Info("delete user")
	w.WriteHeader(204)
	w.Write([]byte("delete user"))
}
