package models

type Theory struct {
    ID       int    `json:"id" db:"id"`
    CourseID int    `json:"course_id" db:"course_id"`
    Title    string `json:"title" db:"title"`
    Content  string `json:"content" db:"content"`
}