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
            "name": "Rodolfo do Nascimento Azevedo",
            "email": "rof20004@gmail.com"
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
        "/consultas": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Consulta"
                ],
                "summary": "Cria uma consulta para um paciente",
                "parameters": [
                    {
                        "description": "Informações do paciente, do profissional e a data da consulta",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/consulta.CreateConsultaRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
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
        "/pacientes": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Paciente"
                ],
                "summary": "Retorna a lista de pacientes",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Paciente"
                ],
                "summary": "Cria um novo paciente",
                "parameters": [
                    {
                        "description": "Informações do paciente",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/paciente.CreatePacienteRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
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
        "/profissionais": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profissional"
                ],
                "summary": "Retorna a lista de todos os profissionais",
                "responses": {
                    "200": {
                        "description": "OK"
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
        "/profissional-agendas/profissionais/{profissionalId}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profissional Agenda"
                ],
                "summary": "Retorna a agenda do profissional pelo id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id do profissional",
                        "name": "profissionalId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "consulta.CreateConsultaRequest": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string",
                    "format": "01/02/2003 14:00"
                },
                "pacienteId": {
                    "type": "string"
                },
                "profissionalId": {
                    "type": "string"
                }
            }
        },
        "paciente.CreatePacienteRequest": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "idade": {
                    "type": "integer"
                },
                "nome": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Ajuda API",
	Description:      "Sistema de agendamento de consultas médicas para a prefeitura de Embu-Guaçu",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
