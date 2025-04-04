// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/task-service/course": {
            "get": {
                "description": "Retrieves all available courses",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courses"
                ],
                "summary": "Get all courses",
                "responses": {
                    "200": {
                        "description": "List of all courses",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Course"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            },
            "post": {
                "description": "Creates a new course",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courses"
                ],
                "summary": "Create new course",
                "parameters": [
                    {
                        "description": "Course object that needs to be created",
                        "name": "course",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Course"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Course successfully created",
                        "schema": {
                            "$ref": "#/definitions/models.Course"
                        }
                    },
                    "201": {
                        "description": "Course successfully created",
                        "schema": {
                            "$ref": "#/definitions/models.Course"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/api/task-service/course/fill/{ID}": {
            "post": {
                "description": "Retrieves content for a specified course",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courses"
                ],
                "summary": "Get course content",
                "parameters": [
                    {
                        "type": "string",
                        "example": "1",
                        "description": "Course ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Course content retrieved successfully",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.CourseContentResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid course ID"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/api/task-service/course/{ID}": {
            "get": {
                "description": "Retrieves a course by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courses"
                ],
                "summary": "Get course by ID",
                "parameters": [
                    {
                        "type": "string",
                        "example": "1",
                        "description": "Course ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Course details",
                        "schema": {
                            "$ref": "#/definitions/models.Course"
                        }
                    },
                    "400": {
                        "description": "Invalid course ID"
                    },
                    "404": {
                        "description": "Course not found"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            },
            "patch": {
                "description": "Updates an existing course by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courses"
                ],
                "summary": "Update course by ID",
                "parameters": [
                    {
                        "type": "string",
                        "example": "1",
                        "description": "Course ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated course data",
                        "name": "course",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Course"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Course successfully updated",
                        "schema": {
                            "$ref": "#/definitions/models.Course"
                        }
                    },
                    "400": {
                        "description": "Invalid course ID or request body"
                    },
                    "404": {
                        "description": "Course not found"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/api/task-service/hint": {
            "post": {
                "description": "Creates a new hint for a specific task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hints"
                ],
                "summary": "Create new hint",
                "parameters": [
                    {
                        "description": "Hint object that needs to be created",
                        "name": "hint",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Hint"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Hint successfully created",
                        "schema": {
                            "$ref": "#/definitions/models.Hint"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/api/task-service/hint/byTask/{taskID}": {
            "get": {
                "description": "Retrieves all hints associated with a specific task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hints"
                ],
                "summary": "Get all hints by task ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Task ID",
                        "name": "task_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of hints",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Hint"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid task ID"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/api/task-service/hint/{ID}": {
            "get": {
                "description": "Retrieves a hint by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hints"
                ],
                "summary": "Get hint by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Hint ID",
                        "name": "ID",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Hint details",
                        "schema": {
                            "$ref": "#/definitions/models.Hint"
                        }
                    },
                    "400": {
                        "description": "Invalid hint ID"
                    },
                    "404": {
                        "description": "Status Not Found"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            },
            "patch": {
                "description": "Updates an existing hint in the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hints"
                ],
                "summary": "Update hint by ID",
                "parameters": [
                    {
                        "type": "string",
                        "example": "1",
                        "description": "Hint ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Hint object with fields to update",
                        "name": "hint",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Hint"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Hint successfully updated",
                        "schema": {
                            "$ref": "#/definitions/models.Hint"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Hint not found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/task-service/task": {
            "post": {
                "description": "Creates a new task in the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Create new task",
                "parameters": [
                    {
                        "description": "Task object that needs to be created",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Task"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Task successfully created",
                        "schema": {
                            "$ref": "#/definitions/models.Task"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/task-service/task/": {
            "get": {
                "description": "Retrieves a list of tasks based on the applied filters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Get tasks by filters",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Task theme",
                        "name": "theme",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "Is task finished",
                        "name": "isFinished",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Minimum complexity",
                        "name": "minComplexity",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Maximum complexity",
                        "name": "maxComplexity",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of filtered tasks",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Task"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/task-service/task/themes": {
            "get": {
                "description": "Retrieves a list of themes for tasks filters",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Get themes for filters",
                "responses": {
                    "200": {
                        "description": "List of themes"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/task-service/task/{ID}": {
            "get": {
                "description": "Retrieves a specific task using its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Get task by ID",
                "parameters": [
                    {
                        "type": "string",
                        "example": "1",
                        "description": "Task ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task object",
                        "schema": {
                            "$ref": "#/definitions/models.Task"
                        }
                    },
                    "400": {
                        "description": "Invalid task ID format"
                    },
                    "404": {
                        "description": "Status Not Found"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            },
            "patch": {
                "description": "Updates an existing task in the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Update task by ID",
                "parameters": [
                    {
                        "type": "string",
                        "example": "1",
                        "description": "Task ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Task object with fields to update",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Task"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task successfully updated",
                        "schema": {
                            "$ref": "#/definitions/models.Task"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Task not found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CourseContentResponse": {
            "type": "object",
            "properties": {
                "content_type": {
                    "type": "string"
                },
                "course_id": {
                    "type": "integer"
                },
                "hints": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Hint"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "order_number": {
                    "type": "integer"
                },
                "task": {
                    "$ref": "#/definitions/models.Task"
                },
                "theory": {
                    "$ref": "#/definitions/models.Theory"
                }
            }
        },
        "models.Course": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_active": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.Hint": {
            "type": "object",
            "properties": {
                "hint_text": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_used": {
                    "type": "boolean"
                },
                "task_id": {
                    "type": "integer"
                },
                "theme": {
                    "type": "string"
                }
            }
        },
        "models.Task": {
            "type": "object",
            "properties": {
                "attempts": {
                    "description": "попытки",
                    "type": "integer"
                },
                "complexity": {
                    "description": "сложность",
                    "type": "integer"
                },
                "course_task_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "is_finished": {
                    "type": "boolean"
                },
                "task_text": {
                    "type": "string"
                },
                "theme": {
                    "description": "тема задачи",
                    "type": "string"
                }
            }
        },
        "models.Theory": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "course_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
