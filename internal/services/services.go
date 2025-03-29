package service

import (
	"github.com/likesense/task-service/internal/dto"
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

type Course interface {
	GetAllCourses() ([]*models.Course, error)
	GetCourseByID(id uint64) (*models.Course, error)
	UpdateCourseByID(id uint64, updatedCourse models.Course) (*models.Course, error)
	CreateNewCourse(course models.Course) (*models.Course, error)
	FillCourseContent(courseID int) (content []dto.CourseContentResponse, err error)
	GetCoursesByFilterList(filters ...func(any) any) ([]*models.Course, error)
}

type Theory interface {
	CreatenewTheory(theory *models.Theory) (newTheory *models.Theory, err error)
	GetTheoryByID(id uint64) (theory *models.Theory, err error)
}

type Services struct {
	Task   Task
	Hint   Hint
	Course Course
	Theory Theory
}

func NewServices(repo *repository.Repositories) *Services {
	return &Services{
		Task:   NewTaskService(repo),
		Hint:   NewHintService(repo),
		Course: NewCourseService(repo),
		Theory: NewTheoryService(repo),
	}
}
