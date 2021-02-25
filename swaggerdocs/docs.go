// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package swaggerdocs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "Tells if the chi-swagger APIs are working or not.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "This API can be used as health check for this application.",
                "responses": {
                    "200": {
                        "description": "api response",
                        "schema": {
                            "$ref": "#/definitions/cmd.response"
                        }
                    }
                }
            }
        },
        "/error": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "error"
                ],
                "summary": "This API always returns an error, and sends an error report to Sentry.io",
                "responses": {
                    "200": {
                        "description": "api response",
                        "schema": {
                            "$ref": "#/definitions/cmd.response"
                        }
                    },
                    "500": {
                        "description": "api response",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponseModel"
                        }
                    }
                }
            }
        },
        "/panic": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "error"
                ],
                "summary": "This API always panics, and sends a stack trace to Sentry.io",
                "responses": {
                    "200": {
                        "description": "api response",
                        "schema": {
                            "$ref": "#/definitions/cmd.response"
                        }
                    },
                    "500": {
                        "description": "api response",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponseModel"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "cmd.response": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "All is good"
                },
                "statusCode": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "utils.ErrorResponseModel": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
