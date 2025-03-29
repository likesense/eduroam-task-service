package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/likesense/task-service/internal/dto"
	"github.com/likesense/task-service/internal/models"
)

type Task interface {
	GetById(id uint64) (*models.Task, error)
	GetAll() (tasks []*models.Task, err error)
	GetAllThemes() ([]string, error)
	GetByFilterList(filters ...func(any) any) (tasks []*models.Task, err error)
	Create(task *models.Task) (newTask *models.Task, err error)
	Update(task *models.Task) (patchTask *models.Task, err error)
}

type Hint interface {
	GetByID(ID uint64) (*models.Hint, error)
	GetAllByTaskID(taskID uint64) (hints []*models.Hint, err error)
	Create(hint *models.Hint) (newHint *models.Hint, err error)
	Update(hint *models.Hint) (patchedHint *models.Hint, err error)
}

type Course interface {
	Create(course *models.Course) (newCourse *models.Course, err error)
	Update(course *models.Course) (patchedCourse *models.Course, err error)
	GetByID(id uint64) (course *models.Course, err error)
	GetAll() ([]*models.Course, error)
	FillByID(courseID int) ([]dto.CourseContentResponse, error)
	GetByFilterList(filters ...func(any) any) (courses []*models.Course, err error)
}

type Theory interface {
	GetByID(id uint64) (theory *models.Theory, err error)
	Create(theory *models.Theory) (newTheory *models.Theory, err error)
}

type Repositories struct {
	Task   Task
	Hint   Hint
	Course Course
	Theory Theory
}

func NewRepositories(postgresDb *sqlx.DB) *Repositories {
	return &Repositories{
		Task:   NewTaskRepository(postgresDb),
		Hint:   NewHintRepository(postgresDb),
		Course: NewCourseRepository(postgresDb),
		Theory: NewTheoryRepository(postgresDb),
	}
}
