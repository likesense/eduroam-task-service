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

func (hr *HintRepository) Create(hint *models.Hint) (newHint *models.Hint, err error) {
	newHint = new(models.Hint)
	err = hr.db.Get(newHint, queries.CreateNewHint, hint.TaskID, hint.Theme, hint.HintText)
	if err != nil {
		return nil, fmt.Errorf("failed to create a hint: %v", err)
	}
	return newHint, nil
}

func (hr *HintRepository) Update(hint *models.Hint) (patchedHint *models.Hint, err error) {
	patchedHint = new(models.Hint)
	err = hr.db.Get(patchedHint, queries.UpdateHintByID, hint.Theme, hint.IsUsed, hint.HintText, hint.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update a hint: %v", err)
	}
	return patchedHint, nil
}
