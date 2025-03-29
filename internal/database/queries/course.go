package queries

const (
	CreateNewCourse = `
	INSERT INTO eduroam.public.course (title, description, is_active)
	VALUES ($1, $2, $3)
	RETURNING id, title, description, is_active 
	`
	GetCourseByID = `
	SELECT c.id, c.title, c.description, c.is_active 
	FROM eduroam.public.course c
	WHERE id = $1
	`
	GetAllCourses = `
	SELECT c.id, c.title, c.description, c.is_active 
	FROM eduroam.public.course c
	ORDER BY c.id
	`
	UpdateCourseByID = `
	UPDATE eduroam.public.course c
	SET title = COALESCE($1, title),
	    description = COALESCE($2, description),
	    is_active = COALESCE($3, is_active)
	WHERE c.id = $4
    RETURNING id, title, description, is_active
	`

	CreateCourseContent = `
	INSERT INTO eduroam.public.course_content (course_id, content_type, content_id, order_number)
	VALUES ($1, $2, $3, $4)
	RETURNING id, course_id, content_type, content_id, order_number
	`
	GetCourseContent = `
	SELECT cc.*, 
	    CASE 
	        WHEN cc.content_type = 'theory' THEN t.content
	        WHEN cc.content_type = 'task' THEN task.task_text
	        WHEN cc.content_type = 'final_test' THEN ft.test_text
	    END as content
	FROM eduroam.public.course_content cc
	LEFT JOIN eduroam.public.theory t ON cc.content_type = 'theory' AND cc.content_id = t.id
	LEFT JOIN eduroam.public.task ON cc.content_type = 'task' AND cc.content_id = task.id
	LEFT JOIN eduroam.public.final_test ft ON cc.content_type = 'final_test' AND cc.content_id = ft.id
	WHERE cc.course_id = $1
	ORDER BY cc.order_number
	`
)
