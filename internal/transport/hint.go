package http

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	service "github.com/likesense/task-service/internal/services"
)

type HintHandler struct {
	service service.Hint
}

func NewHintHandler(service *service.Services) *HintHandler {
	return &HintHandler{
		service: service.Hint,
	}
}

func (hh *HintHandler) RegisterHintRoutes(grp *gin.RouterGroup) {
	hint := grp.Group("/hint")
	{
		hint.GET("/task/:task_id", hh.getAllHints)
		hint.GET("/:ID", hh.getHintByID)
	}
}

// getAllHints returns all hints for a specific task
//
// @Summary Get all hints by task ID
// @Description Retrieves all hints associated with a specific task
// @Tags hints
// @Accept json
// @Produce json
// @Param task_id path string true "Task ID"
// @Success 200 {array} models.Hint "List of hints"
// @Failure 400 "Invalid task ID"
// @Failure 500 "Internal server error"
// @Router /api/task-service/hint/task/{task_id} [get]
func (hh *HintHandler) getAllHints(ctx *gin.Context) {
	taskIDStr := ctx.Param("task_id")
	if strings.TrimSpace(taskIDStr) == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "taskID is required"})
		return
	}
	taskID, err := strconv.ParseUint(taskIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID format"})
		return
	}
	hint, err := hh.service.GetAllHints(taskID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "can not give a list of hints"})
		return
	}
	ctx.JSON(http.StatusOK, hint)

}

// getHintByID returns hint by ID
//
// @Summary Get hint by ID
// @Description Retrieves a hint by its ID
// @Tags hints
// @Accept json
// @Produce json
// @Param ID query string true "Hint ID"
// @Success 200 {object} models.Hint "Hint details"
// @Failure 400 "Invalid hint ID"
// @Failure 500 "Internal server error"
// @Router /api/task-service/hint/{ID} [get]
func (hh *HintHandler) getHintByID(ctx *gin.Context) {
	hintIDStr := ctx.Query("ID")
	if strings.TrimSpace(hintIDStr) == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "hintID is required"})
		return
	}
	hintID, err := strconv.ParseUint(hintIDStr, 10, 64)
	if err != nil {
		log.Println(err.Error())
		return
	}
	hint, err := hh.service.GetHintByID(hintID)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "can not get hint by ID"})
		return
	}
	ctx.JSON(http.StatusOK, hint)
}
