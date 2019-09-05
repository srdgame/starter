// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-09-05 05:33:19.595627266 +0000 UTC m=+177.506544043

package docs

import (
	"bytes"
	"encoding/json"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "starter",
        "contact": {},
        "license": {},
        "version": "1.0"
    },
    "host": "golang-project.com",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/order": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "订单列表"
                ],
                "summary": "订单",
                "parameters": [
                    {
                        "type": "string",
                        "description": "订单id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "0": {
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/entities.Order"
                        }
                    },
                    "404": {}
                }
            }
        }
    },
    "definitions": {
        "entities.Order": {
            "type": "object",
            "properties": {
                "amount": {
                    "description": "总金额",
                    "type": "integer"
                },
                "item_id": {
                    "description": "订单id",
                    "type": "string"
                },
                "total": {
                    "description": "总数量",
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
var SwaggerInfo = swaggerInfo{Schemes: []string{}}

type s struct{}

func (s *s) ReadDoc() string {
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
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
