package service

import (
	"fmt"

	"github.com/likesense/task-service/internal/models"
	repository "github.com/likesense/task-service/internal/repositories"
)

type HintService struct {
	repo repository.Hint
}

func NewHintService(repo *repository.Repositories) *HintService {
	return &HintService{
		repo: repo.Hint,
	}
}

func (hs *HintService) GetAllHints(taskID uint64) (hints []*models.Hint, err error) {
	hints, err = hs.repo.GetAllByTaskID(taskID)
	if err != nil {
		return nil, fmt.Errorf("error getting hints by taskID: %v", err)
	}
	return hints, nil
}
func (hs *HintService) GetHintByID(hintID uint64) (*models.Hint, error) {
	hint, err := hs.repo.GetByID(hintID)
	if err != nil {
		return nil, fmt.Errorf("error getting hint by ID: %v", err)
	}
	return hint, nil
}
