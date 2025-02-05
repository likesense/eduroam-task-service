package models

type Hint struct {
	ID       uint64 `json:"id" db:"id"`
	TaskID   uint64 `json:"task_id" db:"task_id"`
	Theme    string `json:"theme" db:"theme"`
	HintText string `json:"hint_text" db:"hint_text"`
	IsUsed   bool   `json:"is_used" db:"is_used"`
}
