package models

type Task struct {
	ID           uint64 `json:"id" db:"id"`
	Theme        string `json:"theme" db:"theme"` //тема задачи
	IsFinished   bool   `json:"is_finished" db:"is_finished"`
	TaskText     string `json:"task_text" db:"task_text"`
	Attempts     uint16 `json:"attempts" db:"attempts"`     //попытки
	Complexity   uint16 `json:"complexity" db:"complexity"` //сложность
	CourseTaskID uint64 `json:"course_task_id" db:"course_task_id"`
}
