package users

import (
	"errors"

	"net/http"

	"v2/store"
	"v2/types"

	"github.com/gorilla/mux"
)

var errEmailRequired = errors.New("email is required")
var errFirstNameRequired = errors.New("first name is required")
var errLastNameRequired = errors.New("last name is required")
var errPasswordRequired = errors.New("password is required")


type UserService struct {
	store store.Store
}

func NewUserService(s store.Store) *UserService {
	return &UserService{store: s}
}


func (s *UserService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users/register", s.handleUserRegister).Methods(http.MethodPost)
	r.HandleFunc("/users/login", s.handleUserLogin).Methods(http.MethodPost)
}

func (s *UserService) handleUserRegister(w http.ResponseWriter, r *http.Request){

}

func (s *UserService) handleUserLogin(w http.ResponseWriter, r *http.Request){}
func validateUserPayload(user *types.User) error {
	if user.Email == "" {
		return errEmailRequired
	}

	if user.FirstName == "" {
		return errFirstNameRequired
	}

	if user.LastName == "" {
		return errLastNameRequired
	}

	if user.Password == "" {
		return errPasswordRequired
	}

	return nil
}