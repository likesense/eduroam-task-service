package repository

import (
	"fmt"

	sqrl "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/likesense/task-service/internal/database/queries"
	"github.com/likesense/task-service/internal/dto"
	"github.com/likesense/task-service/internal/models"
)

type CourseRepository struct {
	db *sqlx.DB
}

func NewCourseRepository(db *sqlx.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

func (cr *CourseRepository) Create(course *models.Course) (newCourse *models.Course, err error) {
	newCourse = new(models.Course)
	err = cr.db.Get(newCourse, queries.CreateNewCourse, course.Title, course.Description, course.IsActive)
	if err != nil {
		return nil, fmt.Errorf("failed to create new course: %v", err)
	}
	return newCourse, nil
}

func (cr *CourseRepository) Update(course *models.Course) (patchedCourse *models.Course, err error) {
	patchedCourse = new(models.Course)
	err = cr.db.Get(patchedCourse, queries.UpdateCourseByID, course.Title, course.Description, course.IsActive)
	if err != nil {
		return nil, fmt.Errorf("failed to update course: %v", err)
	}
	return patchedCourse, nil
}

func (cr *CourseRepository) GetByID(id uint64) (course *models.Course, err error) {
	course = new(models.Course)
	err = cr.db.Get(course, queries.GetCourseByID, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get course by id: %v", err)
	}
	return course, nil
}

func (cr *CourseRepository) GetAll() ([]*models.Course, error) {
	var courses []*models.Course
	err := cr.db.Select(&courses, queries.GetAllCourses)
	if err != nil {
		return nil, fmt.Errorf("failed to get all courses: %v", err)
	}
	return courses, nil
}

func (cr *CourseRepository) FillByID(courseID int) ([]dto.CourseContentResponse, error) {
	var courseContent []dto.CourseContent

	err := cr.db.Select(&courseContent, queries.GetCourseContent, courseID)
	if err != nil {
		return nil, fmt.Errorf("failed to get course content: %v", err)
	}

	var courseContentResponse []dto.CourseContentResponse
	for _, content := range courseContent {
		response := dto.CourseContentResponse{
			ID:          content.ID,
			CourseID:    content.ContentID,
			ContentType: content.ContentType,
			OrderNumber: content.OrderNumber,
		}
		switch content.ContentType {
		case "theory":
			var theory models.Theory
			err := cr.db.Get(&theory, queries.GetTheoryByID, content.ContentID)
			if err != nil {
				return nil, fmt.Errorf("failed to get theory: %v", err)
			}
			response.Theory = &theory
		case "task":
			var task models.Task
			err := cr.db.Get(&task, queries.GetTaskById, content.ContentID)
			if err != nil {
				return nil, fmt.Errorf("failed to get task: %v", err)
			}
			response.Task = &task
		case "hint":
			var hints []*models.Hint
			err := cr.db.Get(&hints, queries.GetAllHintsByTaskID, content.CourseID)
			if err != nil {
				return nil, fmt.Errorf("failed to get hint: %v", err)
			}
			response.Hints = hints
		}
		courseContentResponse = append(courseContentResponse, response)
	}
	return courseContentResponse, nil
}

func (cr *CourseRepository) GetByFilterList(filters ...func(any) any) (courses []*models.Course, err error) {
	const courseSelectQuery = `c.id, c.title, c.description, c.is_active`

	sb := sqrl.Select(courseSelectQuery).
		From("eduroam.public.course c").
		PlaceholderFormat(sqrl.Dollar)
	for _, filter := range filters {
		sb = filter(sb).(sqrl.SelectBuilder)
	}
	query, args, err := sb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build course query: %v", err)
	}
	err = cr.db.Select(&courses, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error reading courses: %v", err)
	}
	return courses, nil
}
