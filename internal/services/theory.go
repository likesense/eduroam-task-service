package service

import (
	"fmt"

	"github.com/likesense/task-service/internal/models"
	repository "github.com/likesense/task-service/internal/repositories"
)

type TheoryService struct {
	repo repository.Theory
}

func NewTheoryService(repo *repository.Repositories) *TheoryService {
	return &TheoryService{
		repo: repo.Theory,
	}
}

func (ts *TheoryService) CreatenewTheory(theory *models.Theory) (newTheory *models.Theory, err error) {
	newTheory, err = ts.repo.Create(theory)
	if err != nil {
		return nil, fmt.Errorf("error creating theory in service layer: %v", err)
	}
	return newTheory, nil
}

func (ts *TheoryService) GetTheoryByID(id uint64) (theory *models.Theory, err error) {
	theory, err = ts.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("error getting theory in service layer: %v", err)
	}
	return theory, nil
}
