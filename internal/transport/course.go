package http

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/likesense/task-service/internal/database/filters"
	"github.com/likesense/task-service/internal/models"
	service "github.com/likesense/task-service/internal/services"
)

type CourseHandler struct {
	service service.Course
}

func NewCourseHandler(service service.Course) *CourseHandler {
	return &CourseHandler{
		service: service,
	}
}

func (ch *CourseHandler) RegisterCourseRoutes(grp *gin.RouterGroup) {
	course := grp.Group("/course")
	{
		course.GET("/:ID", ch.handleGetCourseByID)
		course.PATCH("/:ID", ch.handleUpdateCourseByID)
		course.POST("/", ch.handleCreateNewCourse)
		// course.GET("/", ch.handleGetAllCourses)
		course.GET("/", ch.handleGetCoursesByFilterList)
		course.POST("/fill/:ID", ch.handleFillCourseByID)
	}
}

// handleGetCourseByID returns a course by its ID
//
// @Summary Get course by ID
// @Description Retrieves a course by its ID
// @Tags courses
// @Accept json
// @Produce json
// @Param ID path string true "Course ID" example(1)
// @Success 200 {object} models.Course "Course details"
// @Failure 400 "Invalid course ID"
// @Failure 404 "Course not found"
// @Failure 500 "Internal server error"
// @Router /api/task-service/course/{ID} [get]
func (ch *CourseHandler) handleGetCourseByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("ID"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	course, err := ch.service.GetCourseByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, course)
}

func (ch *CourseHandler) handleGetCoursesByFilterList(ctx *gin.Context) {
	opts := make([]func(any) any, 0)
	httpCode, err := ch.applyFilters(ctx.Request.URL.Query(), &opts)
	if err != nil {
		ctx.JSON(httpCode, gin.H{"error": err.Error()})
		return
	}
	courses, err := ch.service.GetCoursesByFilterList(opts...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, courses)
}

// handleGetAllCourses returns all available courses
//
// @Summary Get all courses
// @Description Retrieves all available courses
// @Tags courses
// @Accept json
// @Produce json
// @Success 200 {array} models.Course "List of all courses"
// @Failure 500 "Internal server error"
// @Router /api/task-service/course [get]
func (ch *CourseHandler) handleGetAllCourses(ctx *gin.Context) {
	courses, err := ch.service.GetAllCourses()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, courses)
}

// handleUpdateCourseByID updates an existing course
//
// @Summary Update course by ID
// @Description Updates an existing course by its ID
// @Tags courses
// @Accept json
// @Produce json
// @Param ID path string true "Course ID" example(1)
// @Param course body models.Course true "Updated course data"
// @Success 200 {object} models.Course "Course successfully updated"
// @Failure 400 "Invalid course ID or request body"
// @Failure 404 "Course not found"
// @Failure 500 "Internal server error"
// @Router /api/task-service/course/{ID} [patch]
func (ch *CourseHandler) handleUpdateCourseByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("ID"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var course models.Course
	if err := ctx.ShouldBindJSON(&course); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	patchedCourse, err := ch.service.UpdateCourseByID(id, course)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, patchedCourse)
}

// handleCreateNewCourse creates a new course
//
// @Summary Create new course
// @Description Creates a new course
// @Tags courses
// @Accept json
// @Produce json
// @Param course body models.Course true "Course object that needs to be created"
// @Success 200 {object} models.Course "Course successfully created"
// @Success 201 {object} models.Course "Course successfully created"
// @Failure 400 "Bad Request"
// @Failure 500 "Internal server error"
// @Router /api/task-service/course [post]
func (ch *CourseHandler) handleCreateNewCourse(ctx *gin.Context) {
	var course models.Course
	if err := ctx.ShouldBindJSON(&course); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newCourse, err := ch.service.CreateNewCourse(course)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, newCourse)
}

// handleFillCourseByID retrieves content for a course
//
// @Summary Get course content
// @Description Retrieves content for a specified course
// @Tags courses
// @Accept json
// @Produce json
// @Param ID path string true "Course ID" example(1)
// @Success 200 {array} dto.CourseContentResponse "Course content retrieved successfully"
// @Failure 400 "Invalid course ID"
// @Failure 500 "Internal server error"
// @Router /api/task-service/course/fill/{ID} [post]
func (ch *CourseHandler) handleFillCourseByID(ctx *gin.Context) {
	courseID, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	content, err := ch.service.FillCourseContent(courseID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, content)
}

func (ch *CourseHandler) applyFilters(parameters url.Values, opts *[]func(any) any) (int, error) {
	if parameters.Has("title") {
		*opts = append(*opts, filters.ByCourseTitle(parameters.Get("title")))
	}
	if parameters.Has("description") {
		*opts = append(*opts, filters.ByCourseDesription(parameters.Get("description")))
	}
	return 0, nil
}
