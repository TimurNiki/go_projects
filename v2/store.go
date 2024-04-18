package main

import "database/sql"

type Store interface {
	CreateUser(u *User) (*User, error)
	GetUserByID(id string) (*User, error)
	// Projects
	CreateProject(p *Project) error
	GetProject(id string) (*Project, error)
	DeleteProject(id string) error
	// Tasks
	CreateTask(t *Task) (*Task, error)
	GetTask(id string) (*Task, error)
}

type Storage struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) CreateUser(u *User) (*User, error) {
	return nil
}

func (s *Storage) GetUserByID(id string) (*User, error) {
	return nil
}

func (s *Storage) CreateProject(p *Project) error {
	return nil
}

func (s *Storage) GetProject(id string) (*Project, error) {
	return nil
}

func (s *Storage) DeleteProject(id string) error {
	return nil
}

func (s *Storage) CreateTask(t *Task) (*Task, error) {
	return nil
}

func (s *Storage) GetTask(id string) (*Task, error) {
	return nil
}

func (s *Storage) DeleteTask(id string) error {
	return nil
}
