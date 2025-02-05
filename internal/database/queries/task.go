package queries

const (
	GetAllTasks = `
	SELECT t.id, t.theme, t.is_finished, t.attempts, t.complexity
	FROM eduroam.public.task t
	ORDER BY t.id
	`
	GetTaskById = `
	SELECT t.id, t.theme, t.task_text, t.is_finished, t.attempts, t.complexity
	FROM eduroam.public.task t
	WHERE t.id = $1
	`
	GetTasksTheme = `
	SELECT DISTINCT t.theme 
	FROM eduroam.public.task t
	ORDER BY t.theme
	`
)
