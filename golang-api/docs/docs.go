// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/mahasiswa": {
            "get": {
                "description": "get string by NIM",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Mahasiswa"
                ],
                "summary": "Daftar mahasiswa",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Contact Name search",
                        "name": "q",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Mahasiswa"
                        }
                    }
                }
            },
            "post": {
                "description": "tambah mahasiswa baru, jika ada input yang salah print error",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Mahasiswa"
                ],
                "summary": "Tambah mahasiswa",
                "parameters": [
                    {
                        "description": "Mahasiswa body",
                        "name": "contact",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Mahasiswa"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Mahasiswa"
                        }
                    }
                }
            }
        },
        "/mahasiswa/{id}": {
            "put": {
                "description": "Cari data mahasiswa dari Nim, lalu update data tersebut",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Mahasiswa"
                ],
                "summary": "Ubah data mahasiswa sesuai Nim",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Mahasiswa Nim",
                        "name": "nim",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Mahasiswa body",
                        "name": "mahasiswa",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Mahasiswa"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Mahasiswa"
                        }
                    }
                }
            },
            "delete": {
                "description": "Cari data mahasiswa dari Nim, lalu hapus data tersebut",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Mahasiswa"
                ],
                "summary": "Hapus data mahasiswa",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Mahasiswa Nim",
                        "name": "nim",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Mahasiswa"
                        }
                    }
                }
            }
        },
        "/mahasiswa/{nim}": {
            "get": {
                "description": "Cari mahasiswa dari NIM",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Mahasiswa"
                ],
                "summary": "Detail mahasiswa",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Mahasiswa Nim",
                        "name": "nim",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Mahasiswa"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Mahasiswa": {
            "type": "object",
            "properties": {
                "nama": {
                    "type": "string"
                },
                "nim": {
                    "type": "integer"
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
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Mahasiswa API",
	Description: "API Mahasiswa Golang",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
