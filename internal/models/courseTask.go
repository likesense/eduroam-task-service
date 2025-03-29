package models

type CourseTask struct {
	ID          int `json:"id" db:"id"`
	CourseID    int `json:"course_id" db:"course_id"`
	TaskID      int `json:"task_id" db:"task_id"`
	OrderNumber int `json:"order_number" db:"order_number"`
}
