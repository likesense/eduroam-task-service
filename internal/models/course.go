package models

type Course struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	IsActive    bool   `json:"is_active" db:"is_active"`
}
