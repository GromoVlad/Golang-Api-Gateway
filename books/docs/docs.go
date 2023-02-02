// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/book": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "Создать запись о книге",
                "parameters": [
                    {
                        "description": "Тело запроса",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/createBook.DTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/createBook.Response"
                        }
                    }
                }
            }
        },
        "/book/list": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "Возвращает пагинированый список книг",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "Номер страницы",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "maximum": 20,
                        "minimum": 1,
                        "type": "integer",
                        "description": "Кол-во записей на странице",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Идентификатор книги",
                        "name": "book_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Поиск по названию книги",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Идентификатор автора",
                        "name": "author_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Категория",
                        "name": "category",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/listBook.Response"
                        }
                    }
                }
            }
        },
        "/book/{bookId}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "Найти книгу",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор пользователя",
                        "name": "bookId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/findBookResponse.Response"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "Обновить запись о книге",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор книги",
                        "name": "bookId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Тело запроса",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/updateBook.DTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/updateBook.Response"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "Удалить запись о книге",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор книги",
                        "name": "bookId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/deleteBook.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "books.Book": {
            "type": "object",
            "properties": {
                "author_id": {
                    "type": "integer",
                    "format": "int",
                    "example": 42
                },
                "book_id": {
                    "type": "integer",
                    "format": "int",
                    "example": 42
                },
                "category": {
                    "type": "string",
                    "format": "string",
                    "example": "Некрореализм"
                },
                "created_at": {
                    "type": "string",
                    "format": "string",
                    "example": "2022-01-01 00:00:00"
                },
                "description": {
                    "type": "string",
                    "format": "string",
                    "example": "Описание"
                },
                "name": {
                    "type": "string",
                    "format": "string",
                    "example": "Мрак твоих глаз"
                },
                "updated_at": {
                    "type": "string",
                    "format": "string",
                    "example": "2022-01-01 00:00:00"
                }
            }
        },
        "createBook.DTO": {
            "type": "object",
            "required": [
                "author_id",
                "category",
                "name"
            ],
            "properties": {
                "author_id": {
                    "type": "integer"
                },
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "createBook.Response": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/books.Book"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "deleteBook.Response": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "findBookResponse.Response": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/books.Book"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "listBook.ListBookResponse": {
            "type": "object",
            "properties": {
                "books": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/books.Book"
                    }
                },
                "current_page": {
                    "type": "integer",
                    "example": 1
                },
                "limit": {
                    "type": "integer",
                    "example": 10
                }
            }
        },
        "listBook.Response": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/listBook.ListBookResponse"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "updateBook.DTO": {
            "type": "object",
            "properties": {
                "author_id": {
                    "type": "integer"
                },
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "updateBook.Response": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/books.Book"
                },
                "success": {
                    "type": "boolean"
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
