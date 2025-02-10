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
	CreateNewTask = `
	INSERT INTO eduroam.public.task (theme, complexity, task_text) 
	VALUES ($1, $2, $3)
	RETURNING id, theme, is_finished, attempts, complexity, task_text
	`
	UpdateTaskByID = `
	UPDATE eduroam.public.task t 
	SET theme = COALESCE($1, theme),
	    complexity = COALESCE($2, complexity),
	    task_text = COALESCE($3, task_text),
	    attempts = COALESCE($4, attempts),
	    is_finished = COALESCE($5, is_finished)
	WHERE t.id = $6
    RETURNING id, theme, task_text, attempts, is_finished
`
)
