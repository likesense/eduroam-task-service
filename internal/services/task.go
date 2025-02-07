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

func (ts *TaskService) CreateNewTask(task models.Task) (newTask *models.Task, err error) {
	if task.Theme == "" {
		return nil, fmt.Errorf("field 'theme' is requiered")
	}
	if task.TaskText == "" {
		return nil, fmt.Errorf("field 'task_text' is required")
	}
	if task.Complexity <= 0 {
		return nil, fmt.Errorf("field 'complexity' must be UINT and more then 0")
	}
	newTask, err = ts.repo.Create(&task)
	if err != nil {
		return nil, fmt.Errorf("failed to create task: %v", err)
	}
	return newTask, nil
}

func (ts *TaskService) UpdateTaskByID(taskID uint64, newTask models.Task) (patchedTask *models.Task, err error) {
	existingTask, err := ts.repo.GetById(taskID)
	if err != nil {
		return nil, fmt.Errorf("error getting task with ID: %v", err)
	}
	if newTask.Complexity > 0 {
		existingTask.Complexity = newTask.Complexity
	}
	if newTask.Theme != "" {
		existingTask.Theme = newTask.Theme
	}
	if newTask.TaskText != "" {
		existingTask.TaskText = newTask.TaskText
	}
	existingTask.IsFinished = newTask.IsFinished
	existingTask.Attempts = newTask.Attempts

	patchedTask, err = ts.repo.Update(existingTask)
	if err != nil {
		return nil, fmt.Errorf("error patching task: %v", err)
	}
	return patchedTask, nil
}

func (ts *TaskService) GetTasksByFilterList(filters ...func(any) any) ([]*models.Task, error) {
	tasks, err := ts.repo.GetByFilterList(filters...)
	if err != nil {
		return nil, fmt.Errorf("failed to read tasks from eduroamm db: %v", err)
	}
	return tasks, nil
}
