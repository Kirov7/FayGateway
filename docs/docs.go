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
        "/admin/admin_info": {
            "get": {
                "description": "管理员信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员接口"
                ],
                "summary": "管理员信息",
                "operationId": "/admin/admin_info",
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.AdminInfoOutput"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/admin/change_pwd": {
            "post": {
                "description": "管理员信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员接口"
                ],
                "summary": "管理员信息",
                "operationId": "/admin/change_pwd",
                "parameters": [
                    {
                        "description": "body",
                        "name": "polygon",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ChangePwdInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/admin_login/login": {
            "post": {
                "description": "管理员登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员接口"
                ],
                "summary": "管理员登录",
                "operationId": "/admin_login/login",
                "parameters": [
                    {
                        "description": "body",
                        "name": "polygon",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AdminLoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.AdminLoginOutput"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/admin_login/logout": {
            "get": {
                "description": "管理员退出",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员接口"
                ],
                "summary": "管理员退出",
                "operationId": "/admin_login/logout",
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/demo/bind": {
            "post": {
                "description": "测试数据绑定",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "测试数据绑定",
                "operationId": "/demo/bind",
                "parameters": [
                    {
                        "description": "body",
                        "name": "polygon",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.DemoInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.DemoInput"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/service/service_list": {
            "get": {
                "description": "服务列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "服务管理"
                ],
                "summary": "服务列表",
                "operationId": "/service/service_list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "关键词",
                        "name": "info",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页个数",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "当前页数",
                        "name": "page_no",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.ServiceListOutput"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.AdminInfoOutput": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "introduction": {
                    "type": "string"
                },
                "login_time": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "dto.AdminLoginInput": {
            "type": "object",
            "required": [
                "password",
                "user_name"
            ],
            "properties": {
                "password": {
                    "description": "管理员密码",
                    "type": "string",
                    "example": "123456"
                },
                "user_name": {
                    "description": "管理员用户名",
                    "type": "string",
                    "example": "admin"
                }
            }
        },
        "dto.AdminLoginOutput": {
            "type": "object",
            "properties": {
                "token": {
                    "description": "token",
                    "type": "string",
                    "example": "token"
                }
            }
        },
        "dto.ChangePwdInput": {
            "type": "object",
            "required": [
                "password"
            ],
            "properties": {
                "password": {
                    "description": "管理员密码",
                    "type": "string",
                    "example": "123456"
                }
            }
        },
        "dto.DemoInput": {
            "type": "object",
            "required": [
                "age",
                "name",
                "passwd"
            ],
            "properties": {
                "age": {
                    "type": "integer",
                    "example": 20
                },
                "name": {
                    "type": "string",
                    "example": "姓名"
                },
                "passwd": {
                    "type": "string",
                    "example": "123456"
                }
            }
        },
        "dto.ServiceListItemOutput": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "id",
                    "type": "integer"
                },
                "load_type": {
                    "description": "服务类型",
                    "type": "integer"
                },
                "qpd": {
                    "description": "qpd",
                    "type": "integer"
                },
                "qps": {
                    "description": "qps",
                    "type": "integer"
                },
                "service_addr": {
                    "description": "服务地址",
                    "type": "string"
                },
                "service_desc": {
                    "description": "服务描述",
                    "type": "string"
                },
                "service_name": {
                    "description": "服务名称",
                    "type": "string"
                },
                "total_node": {
                    "description": "节点数",
                    "type": "integer"
                }
            }
        },
        "dto.ServiceListOutput": {
            "type": "object",
            "properties": {
                "list": {
                    "description": "列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.ServiceListItemOutput"
                    }
                },
                "total": {
                    "description": "总数",
                    "type": "integer"
                }
            }
        },
        "middleware.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "errmsg": {
                    "type": "string"
                },
                "errno": {
                    "type": "integer"
                },
                "stack": {},
                "trace_id": {}
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
