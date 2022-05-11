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
        "/auth/confirmation/": {
            "post": {
                "description": "User confirmation after registration",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Confirmation",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.AuthConfirmation"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/schemas.Error"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "$ref": "#/definitions/schemas.Error"
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
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/schemas.Error"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "$ref": "#/definitions/schemas.Error"
                        }
                    }
                }
            }
        },
        "/auth/refresh/": {
            "post": {
                "description": "Refresh exists session",
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
                        "description": "Request data",
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
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/schemas.Error"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "$ref": "#/definitions/schemas.Error"
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
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/schemas.Error"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "$ref": "#/definitions/schemas.Error"
                        }
                    }
                }
            }
        },
        "/groups/": {
            "get": {
                "security": [
                    {
                        "ApiHeaderAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "List",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schemas.GroupsListResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "$ref": "#/definitions/schemas.Error"
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
                "description": "Create a new group",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Create",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Group"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Group"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/schemas.Error"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "$ref": "#/definitions/schemas.Error"
                        }
                    }
                }
            }
        },
        "/groups/{id}/": {
            "get": {
                "security": [
                    {
                        "ApiHeaderAuth": []
                    }
                ],
                "description": "Get group detail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Retrieve",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Group identifier",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Group"
                        }
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "$ref": "#/definitions/schemas.Error"
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
                "description": "Update group",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Update",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Group identifier",
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
                            "$ref": "#/definitions/models.Group"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Group"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/schemas.Error"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "$ref": "#/definitions/schemas.Error"
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
                "description": "Delete group",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Delete",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Group identifier",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "$ref": "#/definitions/schemas.Error"
                        }
                    }
                }
            }
        },
        "/profile/": {
            "get": {
                "security": [
                    {
                        "ApiHeaderAuth": []
                    }
                ],
                "description": "Get profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "Retrieve",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "$ref": "#/definitions/schemas.Error"
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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "Update",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.ProfileUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "$ref": "#/definitions/schemas.Error"
                        }
                    }
                }
            }
        },
        "/profile/password/": {
            "put": {
                "security": [
                    {
                        "ApiHeaderAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "Password Update",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.ProfilePasswordUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "$ref": "#/definitions/schemas.Error"
                        }
                    }
                }
            }
        },
        "/projects/": {
            "get": {
                "security": [
                    {
                        "ApiHeaderAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "List",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "groupID",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schemas.ProjectsListResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "$ref": "#/definitions/schemas.Error"
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
                "description": "Create a new project",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "Create",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Project"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Project"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/schemas.Error"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "$ref": "#/definitions/schemas.Error"
                        }
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
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "$ref": "#/definitions/schemas.Error"
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
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/schemas.Error"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "$ref": "#/definitions/schemas.Error"
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
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "$ref": "#/definitions/schemas.Error"
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
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/schemas.Error"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "$ref": "#/definitions/schemas.Error"
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
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "$ref": "#/definitions/schemas.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Access": {
            "type": "object",
            "properties": {
                "groupId": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "projectId": {
                    "type": "integer"
                },
                "role": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "models.Group": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "members": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Access"
                    }
                },
                "name": {
                    "type": "string"
                },
                "projects": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Project"
                    }
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "models.Project": {
            "type": "object",
            "properties": {
                "group": {
                    "$ref": "#/definitions/models.Group"
                },
                "groupId": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "members": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Access"
                    }
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "models.ProjectStatus": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "isFinal": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "projectId": {
                    "type": "integer"
                },
                "sortOrder": {
                    "type": "integer"
                }
            }
        },
        "models.Task": {
            "type": "object",
            "required": [
                "status",
                "title"
            ],
            "properties": {
                "customStatus": {
                    "$ref": "#/definitions/models.ProjectStatus"
                },
                "customStatusId": {
                    "type": "integer"
                },
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
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "schemas.AuthConfirmation": {
            "type": "object",
            "required": [
                "token"
            ],
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "schemas.AuthLoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
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
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                }
            }
        },
        "schemas.AuthRefreshRequest": {
            "type": "object",
            "required": [
                "refreshToken"
            ],
            "properties": {
                "refreshToken": {
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
                "refreshToken": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                }
            }
        },
        "schemas.AuthRegistrationRequest": {
            "type": "object",
            "required": [
                "email",
                "firstName",
                "lastName",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "schemas.Error": {
            "type": "object",
            "properties": {
                "message": {},
                "validationErrors": {}
            }
        },
        "schemas.GroupsListResponse": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Group"
                    }
                },
                "totalCount": {
                    "type": "integer"
                }
            }
        },
        "schemas.ProfilePasswordUpdate": {
            "type": "object",
            "required": [
                "newPassword",
                "oldPassword"
            ],
            "properties": {
                "newPassword": {
                    "type": "string"
                },
                "oldPassword": {
                    "type": "string"
                }
            }
        },
        "schemas.ProfileUpdate": {
            "type": "object",
            "required": [
                "firstName",
                "lastName"
            ],
            "properties": {
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                }
            }
        },
        "schemas.ProjectsListResponse": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Project"
                    }
                },
                "totalCount": {
                    "type": "integer"
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
                "totalCount": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyHeader": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    },
    "tags": [
        {
            "description": "Auth endpoint",
            "name": "auth"
        },
        {
            "description": "Groups endpoint",
            "name": "groups"
        },
        {
            "description": "Projects endpoint",
            "name": "projects"
        },
        {
            "description": "Tasks endpoint",
            "name": "tasks"
        },
        {
            "description": "Profile endpoint",
            "name": "profile"
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
