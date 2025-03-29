package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/likesense/task-service/internal/models"
	service "github.com/likesense/task-service/internal/services"
)

type TheoryHandler struct {
	service service.Theory
}

func NewTheoryHandler(service *service.Services) *TheoryHandler {
	return &TheoryHandler{
		service: service.Theory,
	}
}

func RegisterTheoryRoutes(grp *gin.RouterGroup) {
	theory := grp.Group("/theory")
	{
		theory.GET("/:ID")
		theory.POST("/")
	}
}

func (th *TheoryHandler) handleGetTheoryByID(ctx *gin.Context) {
	theoryID, err := strconv.ParseUint(ctx.Param("ID"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	theory, err := th.service.GetTheoryByID(theoryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, theory)
}

func (th *TheoryHandler) handleCreateNewTheory(ctx *gin.Context) {
	theory := new(models.Theory)
	if err := ctx.ShouldBindJSON(&theory); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTheory, err := th.service.CreatenewTheory(theory)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, newTheory)
}

