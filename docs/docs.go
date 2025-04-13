// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "https://github.com/blackhorseya/todolist",
            "email": "support@example.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/blackhorseya/todolist/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/categories": {
            "get": {
                "description": "取得所有分類列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分類"
                ],
                "summary": "列出分類",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Category"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "建立新的分類",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分類"
                ],
                "summary": "建立分類",
                "parameters": [
                    {
                        "description": "分類資訊",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.CreateCategoryRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.Category"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/v1/categories/{id}": {
            "get": {
                "description": "透過 ID 取得特定分類",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分類"
                ],
                "summary": "取得分類",
                "parameters": [
                    {
                        "type": "string",
                        "description": "分類 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Category"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "更新特定分類的資訊",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分類"
                ],
                "summary": "更新分類",
                "parameters": [
                    {
                        "type": "string",
                        "description": "分類 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新資訊",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.UpdateCategoryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Category"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "刪除特定分類",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "分類"
                ],
                "summary": "刪除分類",
                "parameters": [
                    {
                        "type": "string",
                        "description": "分類 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/v1/todos": {
            "get": {
                "description": "取得所有待辦事項列表，可依照分類、狀態和優先級別進行過濾",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "待辦事項"
                ],
                "summary": "列出待辦事項",
                "parameters": [
                    {
                        "type": "string",
                        "description": "分類 ID",
                        "name": "categoryId",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "狀態",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "優先級別",
                        "name": "priority",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Todo"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "建立新的待辦事項",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "待辦事項"
                ],
                "summary": "建立待辦事項",
                "parameters": [
                    {
                        "description": "待辦事項資訊",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.CreateTodoRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.Todo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/v1/todos/{id}": {
            "get": {
                "description": "透過 ID 取得特定待辦事項",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "待辦事項"
                ],
                "summary": "取得待辦事項",
                "parameters": [
                    {
                        "type": "string",
                        "description": "待辦事項 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Todo"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "更新特定待辦事項的資訊",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "待辦事項"
                ],
                "summary": "更新待辦事項",
                "parameters": [
                    {
                        "type": "string",
                        "description": "待辦事項 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新資訊",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.UpdateTodoRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Todo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "刪除特定待辦事項",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "待辦事項"
                ],
                "summary": "刪除待辦事項",
                "parameters": [
                    {
                        "type": "string",
                        "description": "待辦事項 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Category": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "entity.Priority": {
            "type": "integer",
            "enum": [
                1,
                2,
                3
            ],
            "x-enum-varnames": [
                "Low",
                "Medium",
                "High"
            ]
        },
        "entity.Status": {
            "type": "integer",
            "enum": [
                1,
                2,
                3
            ],
            "x-enum-varnames": [
                "TodoStatus",
                "InProgress",
                "Done"
            ]
        },
        "entity.Todo": {
            "type": "object",
            "properties": {
                "categoryID": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "dueDate": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "priority": {
                    "$ref": "#/definitions/entity.Priority"
                },
                "status": {
                    "$ref": "#/definitions/entity.Status"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "handler.CreateCategoryRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "handler.CreateTodoRequest": {
            "type": "object",
            "required": [
                "categoryId",
                "dueDate",
                "priority",
                "title"
            ],
            "properties": {
                "categoryId": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "dueDate": {
                    "type": "string"
                },
                "priority": {
                    "$ref": "#/definitions/entity.Priority"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "handler.UpdateCategoryRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "handler.UpdateTodoRequest": {
            "type": "object",
            "properties": {
                "categoryId": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "dueDate": {
                    "type": "string"
                },
                "priority": {
                    "$ref": "#/definitions/entity.Priority"
                },
                "status": {
                    "$ref": "#/definitions/entity.Status"
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
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "待辦事項清單 API",
	Description:      "此為使用 Clean Architecture 和 DDD 實作的待辦事項清單 API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
