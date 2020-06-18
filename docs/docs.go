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
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/auth/captcha": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "授权"
                ],
                "summary": "创建验证码",
                "responses": {
                    "100": {
                        "description": "Continue",
                        "schema": {
                            "$ref": "#/definitions/util.Result"
                        }
                    },
                    "104": {
                        "schema": {
                            "$ref": "#/definitions/util.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "授权"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "description": "登录表单",
                        "name": "loginForm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/form.LoginForm"
                        }
                    }
                ],
                "responses": {
                    "100": {
                        "description": "Continue",
                        "schema": {
                            "$ref": "#/definitions/util.Result"
                        }
                    },
                    "104": {
                        "schema": {
                            "$ref": "#/definitions/util.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "授权"
                ],
                "summary": "注册",
                "parameters": [
                    {
                        "description": "注册表单",
                        "name": "regForm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/form.RegisterForm"
                        }
                    }
                ],
                "responses": {
                    "100": {
                        "description": "Continue",
                        "schema": {
                            "$ref": "#/definitions/util.Result"
                        }
                    },
                    "104": {
                        "schema": {
                            "$ref": "#/definitions/util.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/categories": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "分类"
                ],
                "summary": "分页获取分类数据",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页条数",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "关键词",
                        "name": "key",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "分类类型",
                        "name": "category_type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "100": {
                        "description": "Continue",
                        "schema": {
                            "$ref": "#/definitions/util.Result"
                        }
                    },
                    "104": {
                        "schema": {
                            "$ref": "#/definitions/util.Result"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "分类"
                ],
                "summary": "添加分类",
                "parameters": [
                    {
                        "description": "添加分类表单",
                        "name": "addForm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/form.CategoryAddForm"
                        }
                    }
                ],
                "responses": {
                    "100": {
                        "description": "Continue",
                        "schema": {
                            "$ref": "#/definitions/util.Result"
                        }
                    },
                    "104": {
                        "schema": {
                            "$ref": "#/definitions/util.Result"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "分类"
                ],
                "summary": "批量删除分类",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ids",
                        "name": "ids",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "100": {
                        "description": "Continue",
                        "schema": {
                            "$ref": "#/definitions/util.Result"
                        }
                    },
                    "104": {
                        "schema": {
                            "$ref": "#/definitions/util.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/categories/all": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "分类"
                ],
                "summary": "获取所有分类",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "分类类型",
                        "name": "category_type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "100": {
                        "description": "Continue",
                        "schema": {
                            "$ref": "#/definitions/util.Result"
                        }
                    },
                    "104": {
                        "schema": {
                            "$ref": "#/definitions/util.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/categories/parent": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "分类"
                ],
                "summary": "获取所有父级分类",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "分类类型",
                        "name": "category_type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "100": {
                        "description": "Continue",
                        "schema": {
                            "$ref": "#/definitions/util.Result"
                        }
                    },
                    "104": {
                        "schema": {
                            "$ref": "#/definitions/util.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/categories/{id}": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "分类"
                ],
                "summary": "修改分类",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "修改分类表单",
                        "name": "editForm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/form.CategoryEditForm"
                        }
                    }
                ],
                "responses": {
                    "100": {
                        "description": "Continue",
                        "schema": {
                            "$ref": "#/definitions/util.Result"
                        }
                    },
                    "104": {
                        "schema": {
                            "$ref": "#/definitions/util.Result"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "分类"
                ],
                "summary": "删除分类",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "100": {
                        "description": "Continue",
                        "schema": {
                            "$ref": "#/definitions/util.Result"
                        }
                    },
                    "104": {
                        "schema": {
                            "$ref": "#/definitions/util.Result"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "form.CategoryAddForm": {
            "type": "object",
            "required": [
                "name",
                "url"
            ],
            "properties": {
                "name": {
                    "description": "分类名称",
                    "type": "string"
                },
                "parent_id": {
                    "description": "父级分类 ID",
                    "type": "integer"
                },
                "type": {
                    "description": "分类类型，默认值为 0 表文章；1 表友链",
                    "type": "integer"
                },
                "url": {
                    "description": "访问 URL",
                    "type": "string"
                }
            }
        },
        "form.CategoryEditForm": {
            "type": "object",
            "required": [
                "name",
                "url"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "description": "分类名称",
                    "type": "string"
                },
                "parent_id": {
                    "description": "父级分类 ID",
                    "type": "integer"
                },
                "type": {
                    "description": "分类类型，默认值为 0 表文章；1 表友链",
                    "type": "integer"
                },
                "url": {
                    "description": "访问 URL",
                    "type": "string"
                }
            }
        },
        "form.LoginForm": {
            "type": "object",
            "required": [
                "captcha_id",
                "captcha_val",
                "pwd",
                "username"
            ],
            "properties": {
                "captcha_id": {
                    "description": "验证码 ID",
                    "type": "string"
                },
                "captcha_val": {
                    "description": "验证码",
                    "type": "string"
                },
                "pwd": {
                    "description": "密码",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "form.RegisterForm": {
            "type": "object",
            "required": [
                "email",
                "pwd",
                "second_pwd",
                "username"
            ],
            "properties": {
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "pwd": {
                    "description": "密码",
                    "type": "string"
                },
                "second_pwd": {
                    "description": "确认密码",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "util.Result": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "状态码",
                    "type": "integer",
                    "example": 100
                },
                "data": {
                    "description": "数据",
                    "type": "object"
                },
                "msg": {
                    "description": "提示",
                    "type": "string",
                    "example": "密码错误"
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
	Host:        "localhost:8088",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Gin Swagger",
	Description: "Aries 开源博客项目 API 接口文档",
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
