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
        "/INN": {
            "get": {
                "description": "Get company info by INN",
                "summary": "Get",
                "parameters": [
                    {
                        "description": "INN",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/proto.Request"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/proto.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "proto.Request": {
            "type": "object",
            "properties": {
                "INN": {
                    "type": "string"
                }
            }
        },
        "proto.Response": {
            "type": "object",
            "properties": {
                "CompanyName": {
                    "type": "string"
                },
                "Director": {
                    "type": "string"
                },
                "INN": {
                    "type": "string"
                },
                "KPP": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:8081",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Rusprofile",
	Description:      "Get company info by INN",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
