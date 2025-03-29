package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/likesense/task-service/internal/database/queries"
	"github.com/likesense/task-service/internal/models"
)

type TheoryRepository struct {
	db *sqlx.DB
}

func NewTheoryRepository(db *sqlx.DB) *TheoryRepository {
	return &TheoryRepository{
		db: db,
	}
}

func (tr *TheoryRepository) GetByID(id uint64) (theory *models.Theory, err error) {
	theory = new(models.Theory)
	err = tr.db.Get(theory, queries.GetTheoryByID, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get theory id: %v", err)
	}
	return theory, nil
}

func (tr *TheoryRepository) Create(theory *models.Theory) (newTheory *models.Theory, err error) {
	newTheory = new(models.Theory)
	err = tr.db.Get(newTheory, queries.CreateTheory, theory.Title, theory.Content)
	if err != nil {
		fmt.Printf("failed to create new theory: %v", err)
	}
	return newTheory, nil
}
