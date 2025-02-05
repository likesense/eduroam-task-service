package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/likesense/task-service/internal/database/queries"
	"github.com/likesense/task-service/internal/models"
)

type HintRepository struct {
	db *sqlx.DB
}

func NewHintRepository(db *sqlx.DB) *HintRepository {
	return &HintRepository{
		db: db,
	}
}

func (hr *HintRepository) GetByID(ID uint64) (*models.Hint, error) {
	hint := new(models.Hint)
	err := hr.db.Get(hint, queries.GetHintByID, ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get hint by id: %v", err)
	}
	return hint, nil
}

func (hr *HintRepository) GetAllByTaskID(taskID uint64) (hints []*models.Hint, err error) {
	err = hr.db.Select(&hints, queries.GetAllHintsByTaskID, taskID)
	if err != nil {
		return nil, fmt.Errorf("failed to get hints: %v for taskID: %d", err, taskID)
	}
	return hints, nil
}
