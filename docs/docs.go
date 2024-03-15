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
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/event/join": {
            "post": {
                "description": "Join an event by event ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event"
                ],
                "summary": "Join an event",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Event ID",
                        "name": "event_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/exchanges": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "名刺交換の一覧を取得するためのエンドポイント",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Exchange"
                ],
                "summary": "交換した名刺一覧取得",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.GetExchangesOutput"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "名刺交換のためのエンドポイント",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Exchange"
                ],
                "summary": "名刺交換",
                "parameters": [
                    {
                        "description": "Exchange request body",
                        "name": "exchange",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.CreateExchangeInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/me": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "自分の情報を取得する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "自分の情報",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.GetMeOutput"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "ユーザー情報を更新する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "ユーザー情報更新",
                "parameters": [
                    {
                        "description": "更新するユーザー情報(更新後の情報を全て含む必要があります)",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.UpdateUserInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "指定したユーザーの情報を取得する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "ユーザー情報取得",
                "parameters": [
                    {
                        "description": "ユーザーID",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.GetByIDInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.GetByIDOutput"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "schema.CreateExchangeInput": {
            "type": "object",
            "properties": {
                "event_id": {
                    "type": "string"
                },
                "user_id_1": {
                    "type": "string"
                },
                "user_id_2": {
                    "type": "string"
                }
            }
        },
        "schema.Event": {
            "type": "object",
            "properties": {
                "event_id": {
                    "type": "string"
                },
                "finished_at": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "owner_id": {
                    "type": "string"
                },
                "started_at": {
                    "type": "string"
                }
            }
        },
        "schema.Exchange": {
            "type": "object",
            "properties": {
                "event_id": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "schema.GetByIDInput": {
            "type": "object",
            "properties": {
                "user_id": {
                    "type": "string"
                }
            }
        },
        "schema.GetByIDOutput": {
            "type": "object",
            "properties": {
                "events": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.Event"
                    }
                },
                "image_url": {
                    "type": "string"
                },
                "is_organizer": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "skills": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "urls": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "schema.GetExchangesOutput": {
            "type": "object",
            "properties": {
                "exchanges": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.Exchange"
                    }
                }
            }
        },
        "schema.GetMeOutput": {
            "type": "object",
            "properties": {
                "events": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.Event"
                    }
                },
                "image_url": {
                    "type": "string"
                },
                "is_organizer": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "skills": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "urls": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "schema.UpdateUserInput": {
            "type": "object",
            "properties": {
                "image_url": {
                    "type": "string"
                },
                "is_organizer": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "skills": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "urls": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "user_id": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "server-u7kyixk36q-an.a.run.app",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server Petstore server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
