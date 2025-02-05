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
)
