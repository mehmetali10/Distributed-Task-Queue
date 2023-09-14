// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/SmsQueue/Enqueue": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "This endpoint enqueues an SMS for processing in the task queue.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TaskQueue"
                ],
                "summary": "Enqueue SMS",
                "parameters": [
                    {
                        "description": "Request object for enqueuing an SMS",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/shared.EnqueueSmsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response with details of the enqueued SMS",
                        "schema": {
                            "$ref": "#/definitions/shared.EnqueueSmsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/SmsQueue/ReadAll": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "This endpoint retrieves all SMS queue entries.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TaskQueue"
                ],
                "summary": "Read All SMS Queue",
                "responses": {
                    "200": {
                        "description": "List of SMS queue entries",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/shared.SmsQueue"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/SmsQueue/ReadAll/Fail": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "This endpoint retrieves all failed SMS queue entries.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TaskQueue"
                ],
                "summary": "Read All Failed SMS Queue Entries",
                "responses": {
                    "200": {
                        "description": "List of failed SMS queue entries",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/shared.SmsQueue"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/SmsQueue/TriggerWorker": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "This endpoint triggers a worker for processing tasks in the task queue.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TaskQueue"
                ],
                "summary": "Trigger Worker",
                "parameters": [
                    {
                        "description": "Request object for triggering a worker",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/shared.TriggerWorkerRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response with details of worker execution",
                        "schema": {
                            "$ref": "#/definitions/shared.TriggerWorkerResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        }
    },
    "definitions": {
        "shared.EnqueueSmsRequest": {
            "type": "object",
            "required": [
                "phoneNumber",
                "smsBody"
            ],
            "properties": {
                "phoneNumber": {
                    "type": "string"
                },
                "smsBody": {
                    "type": "string"
                }
            }
        },
        "shared.EnqueueSmsResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "shared.SmsQueue": {
            "type": "object",
            "properties": {
                "createdDate": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "smsBody": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "tryCount": {
                    "type": "integer"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "shared.TriggerWorkerRequest": {
            "type": "object"
        },
        "shared.TriggerWorkerResponse": {
            "type": "object",
            "properties": {
                "handledSmsCount": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "Please include a valid bearer token in the 'Authorization' header for authentication.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "TaskQueue API",
	Description:      "This is a microservice server called TaskQueue, which provides various endpoints for managing tasks in a queue. It allows you to enqueue SMS messages, trigger workers for processing tasks, and retrieve SMS queue entries.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
