package service

import (
	"fmt"

	"github.com/likesense/task-service/internal/models"
	repository "github.com/likesense/task-service/internal/repositories"
)

type TaskService struct {
	repo repository.Task
}

func NewTaskService(repo *repository.Repositories) *TaskService {
	return &TaskService{
		repo: repo.Task,
	}
}

func (ts *TaskService) GetTaskById(taskID uint64) (*models.Task, error) {
	task, err := ts.repo.GetById(taskID)
	if err != nil {
		return nil, fmt.Errorf("error getting task by ID: %v", err)
	}
	return task, nil
}

func (ts *TaskService) GetAllTasks() (tasks []*models.Task, err error) {
	tasks, err = ts.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error getting tasks: %v", err)
	}
	return tasks, nil
}

func (ts *TaskService) GetAllThemes() ([]string, error) {
	themes, err := ts.repo.GetAllThemes()
	if err != nil {
		return nil, fmt.Errorf("error getting themes: %v", err)
	}
	return themes, nil
}

func (ts *TaskService) GetTasksByFilterList(filters ...func(any) any) ([]*models.Task, error) {
	tasks, err := ts.repo.GetByFilterList(filters...)
	if err != nil {
		return nil, fmt.Errorf("failed to read tasks from eduroamm db: %v", err)
	}
	return tasks, nil
}
