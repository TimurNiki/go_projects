package projects

import (
	"encoding/json"
	"net/http"

	"github.com/TimurNiki/go_api_tutorial/v2/store"
	"github.com/TimurNiki/go_api_tutorial/v2/utils"
	"github.com/gorilla/mux"
)

func NewProjectService(s store.Store) *ProjectService {
	return &ProjectService{store: s}
}

func (s *ProjectService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/projects", WithJWTAuth(s.handleCreateProject, s.store)).Methods("POST")
	r.HandleFunc("/projects/{id}", WithJWTAuth(s.handleGetProject, s.store)).Methods("GET")
	r.HandleFunc("/projects/{id}", WithJWTAuth(s.handleDeleteProject, s.store)).Methods("DELETE")
}

func (s *ProjectService) handleCreateProject(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var project *Project
	err = json.Unmarshal(body, &project)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid request payload"})
		return
	}

	if project.Name == "" {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Name is required"})
		return
	}

	err = s.store.CreateProject(project)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating project"})
		return
	}

	WriteJSON(w, http.StatusCreated, project)
}

func (s *ProjectService) handleGetProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	project, err := s.store.GetProject(id)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError,utils.ErrorResponse{Error: "Error getting project"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, project)
}

func (s *ProjectService) handleDeleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	project, err := s.store.DeleteProject(id)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError,utils.ErrorResponse{Error: "Error deleting project"})
		return
	}
	utils.WriteJSON(w, http.StatusNoContent, nil)
}