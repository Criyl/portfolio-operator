// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Christopher Carroll",
            "url": "https://carroll.codes",
            "email": "chris@carroll.codes"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/portfolio": {
            "get": {
                "description": "return list of all entries",
                "tags": [
                    "Portfolio"
                ],
                "summary": "return list of all entries",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/carroll_codes_portfolio-operator_api_v1.PortfolioSpec"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/portfolio/tag/{tag}": {
            "get": {
                "description": "return list of all entries with a specified tag",
                "tags": [
                    "Portfolio"
                ],
                "summary": "return list of all entries with a specified tag",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tag",
                        "name": "tag",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/carroll_codes_portfolio-operator_api_v1.PortfolioSpec"
                            }
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "return return ok status if service is healthy",
                "tags": [
                    "Health"
                ],
                "summary": "return return ok status if service is healthy",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "carroll_codes_portfolio-operator_api_v1.PortfolioSpec": {
            "type": "object",
            "properties": {
                "blog": {
                    "type": "string"
                },
                "healthcheck": {
                    "type": "string"
                },
                "icon": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "url": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Portfolio Operator",
	Description:      "Manage your portfolio natively in your kubernetes cluster.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
