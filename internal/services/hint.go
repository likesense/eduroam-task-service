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

func (hs *HintService) CreateNewHint(hint models.Hint) (newHint *models.Hint, err error) {
	if hint.TaskID == 0 {
		return nil, fmt.Errorf("field 'taskID' is required")
	}
	if hint.Theme == "" {
		return nil, fmt.Errorf("field 'theme' is required")
	}
	if hint.HintText == "" {
		return nil, fmt.Errorf("field 'hint_text' is required")
	}
	newHint, err = hs.repo.Create(&hint)
	if err != nil {
		return nil, fmt.Errorf("error creating hint: %v", err)
	}
	return newHint, nil
}
