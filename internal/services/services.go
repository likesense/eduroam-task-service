package service

import (
	"github.com/likesense/task-service/internal/models"
	repository "github.com/likesense/task-service/internal/repositories"
)

type Task interface {
	GetTaskById(taskID uint64) (*models.Task, error)
	GetAllTasks() (tasks []*models.Task, err error)
	GetTasksByFilterList(filters ...func(any) any) ([]*models.Task, error)
	GetAllThemes() ([]string, error)
	CreateNewTask(task models.Task) (newTask *models.Task, err error)
	UpdateTaskByID(taskID uint64, newTask models.Task) (patchedTask *models.Task, err error)
}

type Hint interface {
	GetHintByID(hintID uint64) (*models.Hint, error)
	GetAllHints(taskID uint64) (hints []*models.Hint, err error)
	CreateNewHint(hint models.Hint) (newHint *models.Hint, err error)
	UpdateHintByID(hintID uint64, newHint models.Hint) (patchedHint *models.Hint, err error)
}

type Services struct {
	Task Task
	Hint Hint
}

func NewServices(repo *repository.Repositories) *Services {
	return &Services{
		Task: NewTaskService(repo),
		Hint: NewHintService(repo),
	}
}
