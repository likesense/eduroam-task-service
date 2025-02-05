package http

import (
	"github.com/gin-gonic/gin"
	service "github.com/likesense/task-service/internal/services"
)

type Handler struct {
	TaskHandler *TaskHandler
	HintHandler *HintHandler
}

func NewHandler(service *service.Services) *Handler {
	return &Handler{
		HintHandler: NewHintHandler(service),
		TaskHandler: NewTaskHandler(service),
	}
}

func (h *Handler) RegisterAPI(router *gin.RouterGroup) {
	api := router.Group("/api")
	h.registerTaskAPI(api)
}

func (h *Handler) registerTaskAPI(grp *gin.RouterGroup) {
	task := grp.Group("/task-service")
	{
		h.HintHandler.RegisterHintRoutes(task)
		h.TaskHandler.RegisterTaskRoutes(task)
	}
}
