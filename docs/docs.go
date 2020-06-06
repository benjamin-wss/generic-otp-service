// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Benjamin Wong",
            "url": "http://www.swagger.io/support",
            "email": "do-not-mail-this@gmail.com"
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
        "/": {
            "get": {
                "description": "Returns values regarding sever uptime and caller HTTP request metadata",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "server-health-check"
                ],
                "summary": "Gets status of current server",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.HealthCheckGreeting"
                        }
                    }
                }
            }
        },
        "/api/internal/v1/acquire": {
            "post": {
                "description": "Generates T.O.T.P. number. Read more: https://en.wikipedia.org/wiki/Time-based_One-time_Password_algorithm",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "internal-otp"
                ],
                "summary": "Generates T.O.T.P. number.",
                "parameters": [
                    {
                        "description": "Payload to generate T.O.T.P.",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ApiInputBasicOtp"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.OtpRepositoryTimeBasedOtpResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
                        }
                    }
                }
            }
        },
        "/api/internal/v1/validate": {
            "post": {
                "description": "Validates T.O.T.P. number. Read more: https://en.wikipedia.org/wiki/Time-based_One-time_Password_algorithm",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "internal-otp"
                ],
                "summary": "Validates T.O.T.P. number.",
                "parameters": [
                    {
                        "description": "Payload to validate T.O.T.P.",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ApiInputValidateBasicOtp"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ApiResultValidateBasicOtp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.ApiInputBasicOtp": {
            "type": "object",
            "properties": {
                "length": {
                    "type": "integer"
                },
                "otpLifespanInSeconds": {
                    "type": "integer"
                },
                "requester": {
                    "type": "string",
                    "example": "jim@starfleet.com"
                }
            }
        },
        "dto.ApiInputValidateBasicOtp": {
            "type": "object",
            "properties": {
                "length": {
                    "type": "integer"
                },
                "otp": {
                    "type": "string"
                },
                "otpLifespanInSeconds": {
                    "type": "integer"
                },
                "referenceToken": {
                    "type": "string"
                },
                "requester": {
                    "type": "string",
                    "example": "jim@starfleet.com"
                }
            }
        },
        "dto.ApiResultValidateBasicOtp": {
            "type": "object",
            "properties": {
                "input": {
                    "type": "object",
                    "$ref": "#/definitions/dto.ApiInputValidateBasicOtp"
                },
                "isValid": {
                    "type": "boolean"
                }
            }
        },
        "dto.HealthCheckGreeting": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string",
                    "example": "2020-06-04T00:00:16.2963059+08:00"
                },
                "greeting": {
                    "type": "string",
                    "example": "Ah, la vache! Ze service is working !"
                },
                "uptime": {
                    "type": "string",
                    "example": "10178631900"
                }
            }
        },
        "dto.HttpError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        },
        "dto.OtpRepositoryTimeBasedOtpResult": {
            "type": "object",
            "properties": {
                "expiryInSeconds": {
                    "type": "integer"
                },
                "otp": {
                    "type": "string"
                },
                "referenceToken": {
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
	Version:     "1.0",
	Host:        "localhost:3000",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Generic OTP Service API",
	Description: "Generic OTP API Service.",
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
