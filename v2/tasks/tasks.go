package tasks

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"v2/services/auth"
	"v2/store"
	"v2/utils"

	"v2/types"

	"github.com/gorilla/mux"
)

var errNameRequired = errors.New("Name is required")
var errProjectIDRequired = errors.New("Project id is required")
var errUserIDRequired = errors.New("User id is required")

type TasksService struct {
	store store.Store
}

func NewTasksService(s store.Store) *TasksService {
	return &TasksService{store: s}
}

func (s *TasksService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", auth.WithJWTAuth(s.handleCreateTask, s.store)).Methods("POST")
	r.HandleFunc("/tasks/{id}", auth.WithJWTAuth(s.handleGetTask, s.store)).Methods("GET")
}

func (s *TasksService) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var task *types.Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest,  utils.ErrorResponse{Error: "Invalid request payload"})
		return
	}

	if err := validateTaskPayload(task); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}

	t, err := s.store.CreateTask(task)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError,  utils.ErrorResponse{Error: "Error creating task"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, t)
}

func (s *TasksService) handleGetTask(w http.ResponseWriter, r *http.Request) {

}
func validateTaskPayload(task *types.Task) error {
	if task.Name == "" {
		return errNameRequired
	}

	if task.ProjectID == 0 {
		return errProjectIDRequired
	}

	if task.AssignedToID == 0 {
		return errUserIDRequired
	}

	return nil
}
