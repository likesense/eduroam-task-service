package http

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/likesense/task-service/internal/database/filters"
	service "github.com/likesense/task-service/internal/services"
)

type TaskHandler struct {
	service service.Task
}

func NewTaskHandler(service *service.Services) *TaskHandler {
	return &TaskHandler{
		service: service.Task,
	}
}

func (th *TaskHandler) RegisterTaskRoutes(grp *gin.RouterGroup) {
	task := grp.Group("/task")
	{
		task.GET("", th.handleGetAllTasks)
		task.GET("filter", th.handleGetTaskByFilterList)
		task.GET("/:ID", th.handleGetTaskByID)
		task.GET("/themes", th.handleGetAllThemes)
	}
}

// handleGetAllTasks returns all tasks from the system
//
// @Summary Get all tasks
// @Description Retrieves a list of all tasks from the system
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {array} models.Task "List of tasks"
// @Failure 500 "Internal server error"
// @Router /api/task-service/task [get]
func (th *TaskHandler) handleGetAllTasks(ctx *gin.Context) {
	task, err := th.service.GetAllTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "can not get all tasks"})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

// handleGetAllThemes returns all themes for filters
// @Summary Get themes for filters
// @Description Retrieves a list of themes for tasks filters
// @Tags tasks
// @Produce json
// @Success 200 "List of themes"
// @Failure 500 "Internal Server Error"
// @Router /api/task-service/task/themes [get]
func (th *TaskHandler) handleGetAllThemes(ctx *gin.Context) {
	themes, err := th.service.GetAllThemes()
	if err != nil {
		log.Printf("Error getting themes for tasks: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "can not get all themes"})
		return
	}
	ctx.JSON(http.StatusOK, themes)
}

// handleGetByFilterList returns filtered tasks
// @Summary Get tasks by filters
// @Description Retrieves a list of tasks based on the applied filters
// @Tags tasks
// @Accept json
// @Produce json
// @Param theme query string false "Task theme"
// @Param isFinished query boolean false "Is task finished"
// @Param minComplexity query integer false "Minimum complexity"
// @Param maxComplexity query integer false "Maximum complexity"
// @Success 200 {array} models.Task "List of filtered tasks"
// @Failure 400 "Bad Request"
// @Failure 500 "Internal Server Error"
// @Router /api/task-service/task/filter [get]
func (th *TaskHandler) handleGetTaskByFilterList(ctx *gin.Context) {
	opts := make([]func(any) any, 0)

	httpCode, err := th.applyFilters(ctx.Request.URL.Query(), &opts)
	if err != nil {
		ctx.JSON(httpCode, gin.H{"error": err.Error()})
		return
	}

	tasks, err := th.service.GetTasksByFilterList(opts...)
	if err != nil {
		log.Printf("Error getting filtered tasks: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

// handleGetTaskByID returns a specific task by its ID
//
// @Summary Get task by ID
// @Description Retrieves a specific task using its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param ID path string true "Task ID" example(1)
// @Success 200 {object} models.Task "Task object"
// @Failure 400 "Invalid task ID format"
// @Failure 500 "Internal server error"
// @Router /api/task-service/task/{ID} [get]
func (th *TaskHandler) handleGetTaskByID(ctx *gin.Context) {
	taskIDStr := ctx.Param("ID")
	if strings.TrimSpace(taskIDStr) == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}
	taskID, err := strconv.ParseUint(taskIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "can not parse ID"})
		return
	}
	task, err := th.service.GetTaskById(taskID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "can not get task by ID"})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func (th *TaskHandler) applyFilters(parameters url.Values, opts *[]func(any) any) (int, error) {
	if parameters.Has("theme") {
		*opts = append(*opts, filters.ByTaskTheme(parameters.Get("theme")))
	}
	if parameters.Has("isFinished") {
		isFinished, err := strconv.ParseBool(parameters.Get("isFinished"))
		if err != nil {
			return http.StatusBadRequest, fmt.Errorf("invalid isFinished value: %v", err)
		}
		*opts = append(*opts, filters.ByTaskNotFinished(isFinished))
	}
	if parameters.Has("minComplexity") || parameters.Has("maxComplexity") {
		var min, max uint16 = 0, 255

		if minStr := parameters.Get("minComplexity"); minStr != "" {
			minVal, err := strconv.ParseUint(minStr, 10, 8)
			if err != nil {
				return http.StatusBadRequest, fmt.Errorf("invalid minComplexity value: %v", err)
			}
			min = uint16(minVal)
		}

		if maxStr := parameters.Get("maxComplexity"); maxStr != "" {
			maxVal, err := strconv.ParseUint(maxStr, 10, 8)
			if err != nil {
				return http.StatusBadRequest, fmt.Errorf("invalid maxComplexity value: %v", err)
			}
			max = uint16(maxVal)
		}

		if min > max {
			return http.StatusBadRequest, fmt.Errorf("minComplexity cannot be greater than maxComplexity")
		}

		*opts = append(*opts, filters.ByTaskComplexity(min, max))
	}
	return 0, nil
}
