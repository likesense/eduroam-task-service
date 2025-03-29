package models

type FinalTest struct {
	ID            int    `json:"id" db:"id"`
	CourseID      int    `json:"course_id" db:"course_id"`
	TestText      string `json:"test_text" db:"test_text"`
	CorrectAnswer string `json:"correct_answer" db:"correct_answer"`
	MaxScore      int    `json:"max_score" db:"max_score"`
}
