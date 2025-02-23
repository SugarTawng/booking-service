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
        "/mentor": {
            "put": {
                "description": "Update mentor",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mentor"
                ],
                "summary": "Update mentor",
                "operationId": "update-mentor",
                "parameters": [
                    {
                        "description": "Mentor",
                        "name": "mentor",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Mentor"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response-model_Mentor"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create mentor",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mentor"
                ],
                "summary": "Create mentor",
                "operationId": "create-mentor",
                "parameters": [
                    {
                        "description": "Mentor",
                        "name": "mentor",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Mentor"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response-model_Mentor"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/mentor/{id}": {
            "get": {
                "description": "Get mentor by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mentor"
                ],
                "summary": "Get mentor by ID",
                "operationId": "get-mentor-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Mentor ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response-model_Mentor"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete mentor",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mentor"
                ],
                "summary": "Delete mentor",
                "operationId": "delete-mentor",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Mentor ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response-string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/mentors": {
            "get": {
                "description": "Get all mentors",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mentor"
                ],
                "summary": "Get all mentors",
                "operationId": "get-all-mentors",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response-model_Mentor"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/schedule": {
            "put": {
                "description": "Update schedule",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "schedule"
                ],
                "summary": "Update schedule",
                "operationId": "update-schedule",
                "parameters": [
                    {
                        "description": "Schedule",
                        "name": "schedule",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Schedule"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response-model_Schedule"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create schedule",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "schedule"
                ],
                "summary": "Create schedule",
                "operationId": "create-schedule",
                "parameters": [
                    {
                        "description": "Schedule",
                        "name": "schedule",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateScheduleRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response-model_Schedule"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/schedule/appointment": {
            "post": {
                "description": "Schedule appointment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Learner"
                ],
                "summary": "Schedule appointment",
                "operationId": "schedule-appointment",
                "parameters": [
                    {
                        "description": "Appointment",
                        "name": "appointment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ScheduleAppointmentRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer \u003ctoken\u003e",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response-model_Appointment"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/schedule/appointment/{schedule_id}": {
            "get": {
                "description": "Get appointment by Schedule ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "schedule"
                ],
                "summary": "Get appointment by Schedule ID",
                "operationId": "get-appointment-by-schedule-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Schedule ID",
                        "name": "schedule_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response-model_Appointment"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Unschedule appointment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Learner"
                ],
                "summary": "Unschedule appointment",
                "operationId": "unschedule-appointment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Schedule ID",
                        "name": "schedule_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Bearer \u003ctoken\u003e",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response-string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/schedule/{id}": {
            "get": {
                "description": "Get schedule by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "schedule"
                ],
                "summary": "Get schedule by ID",
                "operationId": "get-schedule-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Schedule ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response-model_Schedule"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete schedule",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "schedule"
                ],
                "summary": "Delete schedule",
                "operationId": "delete-schedule",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Schedule ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response-string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/schedules": {
            "get": {
                "description": "Get all schedules",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "schedule"
                ],
                "summary": "Get all schedules",
                "operationId": "get-all-schedules",
                "parameters": [
                    {
                        "type": "string",
                        "default": "2019-01-01 00:00:00+07",
                        "description": "Start time",
                        "name": "from",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "2222-01-01 00:00:00+07",
                        "description": "End time",
                        "name": "to",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.PaginationResponse-model_Schedule"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/schedules/appointments": {
            "get": {
                "description": "Get many appointments",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "schedule"
                ],
                "summary": "Get many appointments",
                "operationId": "get-many-appointments",
                "parameters": [
                    {
                        "type": "string",
                        "default": "2019-01-01 00:00:00+07",
                        "description": "Start time",
                        "name": "from",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "2222-01-01 00:00:00+07",
                        "description": "End time",
                        "name": "to",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.PaginationResponse-model_Appointment"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/schedules/available": {
            "get": {
                "description": "Get available schedule",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "schedule"
                ],
                "summary": "Get available schedule",
                "operationId": "get-available-schedule",
                "parameters": [
                    {
                        "type": "string",
                        "default": "2019-01-01 00:00:00+07",
                        "description": "Start time",
                        "name": "from",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "2222-01-01 00:00:00+07",
                        "description": "End time",
                        "name": "to",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response-array_model_Schedule"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/schedules/learner/{learner_id}": {
            "get": {
                "description": "Get all schedules by learner ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "schedule"
                ],
                "summary": "Get all schedules by learner ID",
                "operationId": "get-all-schedules-by-learner-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Learner ID",
                        "name": "learner_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response-array_model_Schedule"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/schedules/mentor/{mentor_id}": {
            "get": {
                "description": "Get all schedules by mentor ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "schedule"
                ],
                "summary": "Get all schedules by mentor ID",
                "operationId": "get-all-schedules-by-mentor-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Mentor ID",
                        "name": "mentor_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response-array_model_Schedule"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Appointment": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "detail": {},
                "learner_id": {
                    "type": "string"
                },
                "schedule": {
                    "description": "Omit from GORM processing",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.Schedule"
                        }
                    ]
                },
                "schedule_id": {
                    "type": "string"
                },
                "week_number": {
                    "type": "integer"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "model.CreateScheduleRequest": {
            "type": "object",
            "required": [
                "mentor_id",
                "start_at"
            ],
            "properties": {
                "mentor_id": {
                    "type": "string"
                },
                "start_at": {
                    "type": "string"
                }
            }
        },
        "model.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "error": {
                    "type": "string",
                    "example": "message"
                }
            }
        },
        "model.Mentor": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "career_level": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "model.PaginationResponse-model_Appointment": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Appointment"
                    }
                },
                "message": {
                    "type": "string"
                },
                "paging": {
                    "$ref": "#/definitions/model.Paging"
                }
            }
        },
        "model.PaginationResponse-model_Schedule": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Schedule"
                    }
                },
                "message": {
                    "type": "string"
                },
                "paging": {
                    "$ref": "#/definitions/model.Paging"
                }
            }
        },
        "model.Paging": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "model.Response-array_model_Schedule": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Schedule"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.Response-model_Appointment": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/model.Appointment"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.Response-model_Mentor": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/model.Mentor"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.Response-model_Schedule": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/model.Schedule"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.Response-string": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.Schedule": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "interval": {
                    "description": "this field does not store in db",
                    "type": "object",
                    "properties": {
                        "from": {
                            "type": "string"
                        },
                        "to": {
                            "type": "string"
                        }
                    }
                },
                "mentor_id": {
                    "type": "string"
                },
                "start_at": {
                    "description": "this field does store in db can be used to generate intervals and wday",
                    "type": "string"
                },
                "wday": {
                    "description": "this field does not store in db",
                    "type": "string"
                }
            }
        },
        "model.ScheduleAppointmentRequest": {
            "type": "object",
            "required": [
                "learner_id"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "learner_id": {
                    "type": "string"
                },
                "schedule_at": {
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
