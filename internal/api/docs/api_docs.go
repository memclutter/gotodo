// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate_api = `{
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
        "/auth/info/": {
            "get": {
                "security": [
                    {
                        "ApiHeaderAuth": []
                    }
                ],
                "description": "Get current session info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/auth/login/": {
            "post": {
                "description": "Login in new sessions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Request data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.AuthLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.AuthLoginResponse"
                        }
                    }
                }
            }
        },
        "/auth/logout/": {
            "post": {
                "security": [
                    {
                        "ApiHeaderAuth": []
                    }
                ],
                "description": "Logout current session",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Logout",
                "responses": {
                    "204": {
                        "description": ""
                    }
                }
            }
        },
        "/auth/refresh/": {
            "post": {
                "description": "Refresh current session",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Refresh",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.AuthRefreshRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.AuthRefreshResponse"
                        }
                    }
                }
            }
        },
        "/auth/registration/": {
            "post": {
                "description": "Register new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Registration",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.AuthRegistrationRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": ""
                    }
                }
            }
        },
        "/tasks/": {
            "get": {
                "security": [
                    {
                        "ApiHeaderAuth": []
                    }
                ],
                "description": "Get tasks list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "List",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.TasksListResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiHeaderAuth": []
                    }
                ],
                "description": "Create new tasks",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Create",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Task"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Task"
                        }
                    }
                }
            }
        },
        "/tasks/{id}/": {
            "get": {
                "security": [
                    {
                        "ApiHeaderAuth": []
                    }
                ],
                "description": "Get task details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Retrieve",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task identifier",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Task"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiHeaderAuth": []
                    }
                ],
                "description": "Update task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Update",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task identifier",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Task"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Task"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiHeaderAuth": []
                    }
                ],
                "description": "Delete task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Delete",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task identifier",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Session": {
            "type": "object",
            "properties": {
                "dateCreated": {
                    "type": "string"
                },
                "dateExpired": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.Task": {
            "type": "object",
            "properties": {
                "dateCreated": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "sessions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Session"
                    }
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "schemas.AuthLoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "schemas.AuthLoginResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "refreshToken": {
                    "$ref": "#/definitions/models.Session"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                }
            }
        },
        "schemas.AuthRefreshRequest": {
            "type": "object",
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "schemas.AuthRefreshResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                }
            }
        },
        "schemas.AuthRegistrationRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "schemas.TasksListResponse": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Task"
                    }
                },
                "total_count": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyHeader": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "Header"
        }
    },
    "tags": [
        {
            "description": "Auth endpoints",
            "name": "auth"
        },
        {
            "description": "Tasks endpoint",
            "name": "tasks"
        }
    ]
}`

// SwaggerInfo_api holds exported Swagger Info so clients can modify it
var SwaggerInfo_api = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost",
	BasePath:         "/api/",
	Schemes:          []string{},
	Title:            "gotodo",
	Description:      "gotodo api",
	InfoInstanceName: "api",
	SwaggerTemplate:  docTemplate_api,
}

func init() {
	swag.Register(SwaggerInfo_api.InstanceName(), SwaggerInfo_api)
}
