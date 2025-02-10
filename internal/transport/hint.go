package http

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/likesense/task-service/internal/models"
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
		hint.GET("/byTask/:taskID", hh.handleGetAllHintsByTaskID)
		hint.GET("/:ID", hh.handleGetHintByID)
		hint.PATCH("/:ID", hh.handlePatchHintByID)
		hint.POST("", hh.handleCreateNewHint)
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
// @Router /api/task-service/hint/byTask/{taskID} [get]
func (hh *HintHandler) handleGetAllHintsByTaskID(ctx *gin.Context) {
	taskIDStr := ctx.Param("taskID")
	if strings.TrimSpace(taskIDStr) == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "taskID is required"})
		return
	}
	taskID, err := strconv.ParseUint(taskIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid taskID format"})
		return
	}
	hint, err := hh.service.GetAllHints(taskID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "can not give a list of hints"})
		return
	}

	ctx.JSON(http.StatusOK, hint)

}

// handleUpdateHintByID updates an existing hint
//
// @Summary Update hint by ID
// @Description Updates an existing hint in the system
// @Tags hints
// @Accept json
// @Produce json
// @Param ID path string true "Hint ID" example(1)
// @Param hint body models.Hint true "Hint object with fields to update"
// @Success 200 {object} models.Hint "Hint successfully updated"
// @Failure 400 "Bad Request"
// @Failure 404 "Hint not found"
// @Failure 500 "Internal Server Error"
// @Router /api/task-service/hint/{ID} [patch]
func (hh *HintHandler) handlePatchHintByID(ctx *gin.Context) {
	HintIDStr := ctx.Param("ID")
	if len(strings.TrimSpace(HintIDStr)) == 0{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}
	HintID, err := strconv.ParseUint(HintIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid Hint ID format"})
		return
	}
	var hint models.Hint
	if err = ctx.ShouldBindJSON(&hint); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "invalid input format"})
		return
	}
	patchedHint, err := hh.service.UpdateHintByID(HintID, hint)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, patchedHint)
}

// handleCreateNewHint creates a new hint
//
// @Summary Create new hint
// @Description Creates a new hint for a specific task
// @Tags hints
// @Accept json
// @Produce json
// @Param hint body models.Hint true "Hint object that needs to be created"
// @Success 201 {object} models.Hint "Hint successfully created"
// @Failure 400 "Bad Request"
// @Failure 500 "Internal server error"
// @Router /api/task-service/hint [post]
func (hh *HintHandler) handleCreateNewHint(ctx *gin.Context) {
	var hint models.Hint
	if err := ctx.ShouldBindJSON(&hint); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input format"})
		return
	}
	newHint, err := hh.service.CreateNewHint(hint)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, newHint)
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
// @Failure 404 "Status Not Found"
// @Failure 500 "Internal server error"
// @Router /api/task-service/hint/{ID} [get]
func (hh *HintHandler) handleGetHintByID(ctx *gin.Context) {
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if hint == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("hint with ID: %v, not found", hintID)})
		return
	}

	ctx.JSON(http.StatusOK, hint)
}
