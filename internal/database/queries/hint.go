package queries

const (
	GetHintByID = `
	SELECT h.id, h.task_id, h.theme, h.hint_text, h.is_used
	FROM eduroam.public.hint h
	WHERE h.id = $1
	`
	GetAllHintsByTaskID = `
	SELECT h.id, h.task_id, h.theme, h.hint_text, h.is_used
	FROM eduroam.public.hint h 
	WHERE h.task_id = $1
	ORDER BY h.id
	`
	CreateNewHint = `
	INSERT INTO eduroam.public.hint (task_id, theme, hint_text)
	VALUES ($1, $2, $3)
	RETURNING id, task_id, theme, hint_text, is_used 
	`
	UpdateHintByID = `
	UPDATE eduroam.public.hint h
	SET h.theme = COALESCE($1, h.theme)
		h.task_id = COALESCE($1, h.theme)
		h.is_used = COALESCE($1, h.theme)
		h.hint_text = COALESCE($1, h.theme)
	WHERE h.id = $5
	RETURNING h.id, h.task_id, h.theme, h.hint_text, h.is_used
	`
)
