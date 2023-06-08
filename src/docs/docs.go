// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Oswaldo Rafael Zamora Ramirez",
            "url": "https://github.com/reitmas32",
            "email": "rafa.zamora.rals@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/computer": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Computers"
                ],
                "summary": "get a item of the computers",
                "operationId": "get-computer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of Computer",
                        "name": "id-computer",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Computers"
                ],
                "summary": "update a item of the computers",
                "operationId": "put-computer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of Computer",
                        "name": "id-computer",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Schema by Update New Computer",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.ComputerUpdateSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Computers"
                ],
                "summary": "add a new item of the computers",
                "operationId": "create-computer",
                "parameters": [
                    {
                        "description": "Schema by Create New Computer",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.ComputerCreateSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Computers"
                ],
                "summary": "delete a item of the computers",
                "operationId": "delete-computer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of Computer",
                        "name": "id-computer",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/computer-lab": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Computer Lab"
                ],
                "summary": "get a item of the computer-lab",
                "operationId": "get-computer-lab",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of ComputerLab",
                        "name": "id-computer-lab",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Computer Lab"
                ],
                "summary": "update a item of the computer-lab",
                "operationId": "put-computer-lab",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID del item",
                        "name": "id-computer-lab",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Schema by Update New Computer Lab",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.ComputerLabUpdateSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Computer Lab"
                ],
                "summary": "add a new item of the computer-lab",
                "operationId": "create-computer-lab",
                "parameters": [
                    {
                        "description": "Schema by Create New Computer Lab",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.ComputerLabCreateSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Computer Lab"
                ],
                "summary": "delete a item of the computer-lab",
                "operationId": "delete-computer-lab",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID del item",
                        "name": "id-computer-lab",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/computer-labs": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Computer Lab"
                ],
                "summary": "get all items of the computer-lab",
                "operationId": "get-computer-labs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/computer-labs-limit/{length}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Computer Lab"
                ],
                "summary": "get N items of the computer-lab",
                "operationId": "get-computer-labs_limit",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Length of result",
                        "name": "length",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/link-account": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "LinkAccount User whit ComputerLab",
                "operationId": "put-link-account",
                "parameters": [
                    {
                        "description": "Schema by LinkAccount User whit ComputerLab",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.LinkAccountConfirmationSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "LinkAccount User whit ComputerLab",
                "operationId": "post-link-account",
                "parameters": [
                    {
                        "description": "Schema by LinkAccount User whit ComputerLab",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.LinkAccountRequisitionSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/room": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rooms"
                ],
                "summary": "get a item of the rooms",
                "operationId": "get-room",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of Room",
                        "name": "id-room",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rooms"
                ],
                "summary": "update a item of the rooms",
                "operationId": "put-room",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of Room",
                        "name": "id-room",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Schema by Update New Room",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.RoomUpdateSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rooms"
                ],
                "summary": "add a new item of the rooms",
                "operationId": "create-room",
                "parameters": [
                    {
                        "description": "Schema by Create New Room",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.RoomCreateSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rooms"
                ],
                "summary": "delete a item of the rooms",
                "operationId": "delete-room",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of Room",
                        "name": "id-room",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/signin": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "SignIn User",
                "operationId": "signin-user",
                "parameters": [
                    {
                        "description": "Schema by SignIn User",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.UserSignInSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/signup": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "SignUp User",
                "operationId": "signup-user",
                "parameters": [
                    {
                        "description": "Schema by SignUp User",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.UserSignUpSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/student": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Students"
                ],
                "summary": "get a item of the students",
                "operationId": "get-student",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of Student",
                        "name": "id-student",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Students"
                ],
                "summary": "update a item of the students",
                "operationId": "put-student",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of Students",
                        "name": "id-student",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Schema by Update New Student",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.StudentUpdateSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Students"
                ],
                "summary": "add a new item of the students",
                "operationId": "create-student",
                "parameters": [
                    {
                        "description": "Schema by Create New Student",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.StudentCreateSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Students"
                ],
                "summary": "delete a item of the students",
                "operationId": "delete-student",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of Student",
                        "name": "id-student",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Response": {
            "type": "object",
            "properties": {
                "Data": {},
                "Message": {
                    "type": "string",
                    "example": "Equipo 2"
                },
                "Success": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "schemas.ComputerCreateSchema": {
            "description": "Descripción de la estructura A",
            "type": "object",
            "properties": {
                "id_room": {
                    "type": "string",
                    "example": "62dcb6fc-0f68-4a50-9d8c-8fe352b0f7f3"
                },
                "name": {
                    "type": "string",
                    "example": "Equipo 2"
                },
                "type": {
                    "type": "string",
                    "example": "Impresora"
                }
            }
        },
        "schemas.ComputerLabCreateSchema": {
            "type": "object",
            "required": [
                "description",
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
        "schemas.ComputerLabUpdateSchema": {
            "type": "object",
            "required": [
                "description",
                "name"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "description": "Data",
                    "type": "string"
                }
            }
        },
        "schemas.ComputerUpdateSchema": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Solo falta instalar DevC++"
                },
                "name": {
                    "description": "Data",
                    "type": "string",
                    "example": "Equipo 3"
                },
                "pos_x": {
                    "description": "UI",
                    "type": "number",
                    "example": 156.29
                },
                "pos_y": {
                    "type": "number",
                    "example": 157.3
                },
                "state": {
                    "type": "string",
                    "example": "Disponible"
                },
                "type": {
                    "type": "string",
                    "example": "Prestamo"
                }
            }
        },
        "schemas.LinkAccountConfirmationSchema": {
            "type": "object",
            "required": [
                "code",
                "user_name"
            ],
            "properties": {
                "code": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "schemas.LinkAccountRequisitionSchema": {
            "type": "object",
            "required": [
                "idComputerLab",
                "user_name"
            ],
            "properties": {
                "idComputerLab": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "schemas.RoomCreateSchema": {
            "type": "object",
            "properties": {
                "id_computer_lab": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "schemas.RoomUpdateSchema": {
            "type": "object",
            "properties": {
                "index": {
                    "type": "integer"
                },
                "name": {
                    "description": "Data",
                    "type": "string"
                }
            }
        },
        "schemas.StudentCreateSchema": {
            "type": "object",
            "required": [
                "account_number",
                "last_name",
                "name",
                "semester",
                "university_program"
            ],
            "properties": {
                "account_number": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "name": {
                    "description": "Info Personal",
                    "type": "string"
                },
                "semester": {
                    "type": "integer"
                },
                "university_program": {
                    "description": "Info Academic",
                    "type": "string"
                }
            }
        },
        "schemas.StudentUpdateSchema": {
            "type": "object",
            "properties": {
                "account_number": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "name": {
                    "description": "Info Personal",
                    "type": "string"
                },
                "semester": {
                    "type": "integer"
                },
                "university_program": {
                    "description": "Info Academic",
                    "type": "string"
                }
            }
        },
        "schemas.UserSignInSchema": {
            "type": "object",
            "required": [
                "password",
                "user_name"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "schemas.UserSignUpSchema": {
            "type": "object",
            "required": [
                "date_of_birth",
                "email",
                "last_name",
                "name",
                "password",
                "phone_number",
                "role",
                "user_name"
            ],
            "properties": {
                "date_of_birth": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "UNICyber-API",
	Description:      "This is a API by System UNICyber|SISEC https://github.com/reitmas32/UNICyber-BackEnd",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
