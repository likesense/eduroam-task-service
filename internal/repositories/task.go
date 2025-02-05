package repository

import (
	"fmt"
	"log"

	sqrl "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/likesense/task-service/internal/database/queries"
	"github.com/likesense/task-service/internal/models"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (tr *TaskRepository) GetById(id uint64) (*models.Task, error) {
	task := new(models.Task)
	err := tr.db.Get(task, queries.GetTaskById, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get task from db: %v", err)
	}
	return task, nil
}

func (tr *TaskRepository) GetAll() (tasks []*models.Task, err error) {

	err = tr.db.Select(&tasks, queries.GetAllTasks)
	if err != nil {
		return nil, fmt.Errorf("failed to get all tasks: %v", err)
	}
	return tasks, nil
}

func (tr *TaskRepository) GetAllThemes() ([]string, error) {
	var themes []string
	err := tr.db.Select(&themes, queries.GetTasksTheme)
	if err != nil {
		return nil, fmt.Errorf("failed to get themes from db: %v", err)
	}
	return themes, nil
}

func (tr *TaskRepository) GetByFilterList(filters ...func(any) any) (tasks []*models.Task, err error) {
	const taskSelectQuery = `t.id, t.theme, t.task_text,
	 t.attempts, t.is_finished, t.complexity
	`
	sb := sqrl.Select(taskSelectQuery).
		From("eduroam.public.task t").
		PlaceholderFormat(sqrl.Dollar)
	for _, filter := range filters {
		sb = filter(sb).(sqrl.SelectBuilder)
	}

	query, args, err := sb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build task query: %v", err)
	}
	log.Println("запрос", query, "аргументы", args)
	err = tr.db.Select(&tasks, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error reading tasks: %v", err)
	}
	return tasks, nil
}
