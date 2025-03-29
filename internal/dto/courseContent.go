package dto

import "github.com/likesense/task-service/internal/models"

type CourseContent struct {
	ID          uint64 `json:"id" db:"id"`
	CourseID    uint64 `json:"course_id" db:"course_id"`
	ContentType string `json:"content_type" db:"content_type"`
	ContentID   uint64 `json:"content_id" db:"content_id"`
	OrderNumber uint64 `json:"order_number" db:"order_number"`
	Content     any    `json:"content" db:"content"`
}

type CourseContentResponse struct {
	ID          uint64         `json:"id"`
	CourseID    uint64         `json:"course_id"`
	ContentType string         `json:"content_type"`
	OrderNumber uint64         `json:"order_number"`
	Theory      *models.Theory `json:"theory,omitempty"`
	Task        *models.Task   `json:"task,omitempty"`
	Hints       []*models.Hint `json:"hints,omitempty"`
}
