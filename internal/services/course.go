package service

import (
	"fmt"

	"github.com/likesense/task-service/internal/dto"
	"github.com/likesense/task-service/internal/models"
	repository "github.com/likesense/task-service/internal/repositories"
)

type CourseService struct {
	repo repository.Course
}

func NewCourseService(repo *repository.Repositories) *CourseService {
	return &CourseService{
		repo: repo.Course,
	}
}

func (cs *CourseService) GetAllCourses() ([]*models.Course, error) {
	courses, err := cs.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all courses in service: %v", err)
	}
	return courses, nil
}

func (cs *CourseService) GetCourseByID(id uint64) (*models.Course, error) {
	course, err := cs.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get course by id: %v", err)
	}
	return course, err
}

func (cs *CourseService) UpdateCourseByID(id uint64, updatedCourse models.Course) (*models.Course, error) {
	existingCourse, err := cs.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("error getting course by ID while update it: %v", err)
	}
	if updatedCourse.Title != "" {
		existingCourse.Title = updatedCourse.Title
	}
	if updatedCourse.Description != "" {
		existingCourse.Description = updatedCourse.Description
	}
	existingCourse.IsActive = updatedCourse.IsActive

	patchedCourse, err := cs.repo.Update(existingCourse)
	if err != nil {
		return nil, fmt.Errorf("failed to update existing course: %v", err)
	}
	return patchedCourse, nil
}

func (cs *CourseService) CreateNewCourse(course models.Course) (*models.Course, error) {
	if course.Title == "" {
		return nil, fmt.Errorf("field 'title' is required")
	}
	if course.Description == "" {
		return nil, fmt.Errorf("field 'description' is required")
	}
	newCourse, err := cs.repo.Create(&course)
	if err != nil {
		return nil, fmt.Errorf("failed to create new course in service: %v", err)
	}
	return newCourse, nil
}

func (cs *CourseService) FillCourseContent(courseID int) (content []dto.CourseContentResponse, err error) {
	content, err = cs.repo.FillByID(courseID)
	if err != nil {
		return nil, fmt.Errorf("failed to fill course in service: %v", err)
	}
	return content, nil
}

func (cs *CourseService) GetCoursesByFilterList(filters ...func(any)any)([]*models.Course, error){
	courses, err := cs.repo.GetByFilterList(filters...)
	if err != nil {
		return nil, fmt.Errorf("failed to read courses from eduroam db: %v", err)
	}
	return courses, nil
}