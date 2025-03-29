package queries

const (
	CreateTheory = `
	INSERT INTO theory (course_id, title, content)
	VALUES ($1, $2, $3)
	RETURNING id, course_id, title, content
	`
	GetTheoryByID = `
	SELECT t.id, t.course_id, t.title, t.content 
	FROM eduroam.public.theory t
	WHERE t.id = $1 
	`
)
