// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-05-08 21:51:07.9300943 +0800 CST m=+0.053856601

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
        "termsOfService": "http://zhimiao.org",
        "contact": {
            "name": "API Support",
            "url": "http://tools.zhimiao.org",
            "email": "mail@xiaoliu.org"
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
        "/miniprogram/Config": {
            "post": {
                "description": "小程序配置",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "小程序控制"
                ],
                "summary": "小程序配置",
                "parameters": [
                    {
                        "description": "body参数",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.MiniAppConfigParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 1,\"msg\": \"操作成功\",\"data\": null}",
                        "schema": {
                            "$ref": "#/definitions/resp.ApiResult"
                        }
                    }
                }
            }
        },
        "/miniprogram/GetWXACodeUnlimit": {
            "get": {
                "tags": [
                    "小程序信息获取"
                ],
                "summary": "获取线上小程序码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "小程序APPID",
                        "name": "MiniProgramID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "页面",
                        "name": "Page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "场景，不得超过32位，不得含有特殊符号",
                        "name": "Scene",
                        "in": "query"
                    }
                ]
            }
        },
        "/miniprogram/Lists": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "小程序信息获取"
                ],
                "summary": "小程序列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "排序列",
                        "name": "orderColumn",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "小程序appid",
                        "name": "appID",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "开放平台ID",
                        "name": "platformID",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "-1-授权失效 1授权成功，2审核中，3审核通过，4审核失败，5已发布",
                        "name": "state",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "排序方式，desc倒序， asc正序",
                        "name": "orderClase",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "lastId",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "版本号",
                        "name": "version",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/resp.MiniAppListsVO"
                            }
                        }
                    }
                }
            }
        },
        "/oplatform/manage/Lists": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "第三方平台管理"
                ],
                "summary": "获取平台列表",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "lastId",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Platform"
                            }
                        }
                    }
                }
            }
        },
        "/oplatform/manage/Set": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "第三方平台管理"
                ],
                "summary": "设施平台信息",
                "parameters": [
                    {
                        "description": "入参集合",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.PlatformParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 1,\"msg\": \"操作成功\",\"data\": null}",
                        "schema": {
                            "$ref": "#/definitions/resp.ApiResult"
                        }
                    }
                }
            }
        },
        "/oplatform/{PlatformID}/CommitCode": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "小程序控制"
                ],
                "summary": "提交代码",
                "parameters": [
                    {
                        "description": "入参集合",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.CommitCodeParam"
                        }
                    },
                    {
                        "type": "string",
                        "description": "第三方平台APPID",
                        "name": "PlatformID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.ApiResult"
                        }
                    }
                }
            }
        },
        "/oplatform/{PlatformID}/GetAuthorizerInfo": {
            "get": {
                "tags": [
                    "小程序信息获取"
                ],
                "summary": "获取小程序授权信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "小程序APPID",
                        "name": "MiniProgramID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "第三方平台APPID",
                        "name": "PlatformID",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        },
        "/oplatform/{PlatformID}/GetCodeCategory": {
            "get": {
                "tags": [
                    "小程序信息获取"
                ],
                "summary": "获取已设置类目",
                "parameters": [
                    {
                        "type": "string",
                        "description": "小程序APPID",
                        "name": "MiniProgramID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "第三方平台APPID",
                        "name": "PlatformID",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        },
        "/oplatform/{PlatformID}/GetCodePageList": {
            "get": {
                "tags": [
                    "小程序信息获取"
                ],
                "summary": "获取小程序线上页面列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "小程序APPID",
                        "name": "MiniProgramID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "第三方平台APPID",
                        "name": "PlatformID",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        },
        "/oplatform/{PlatformID}/GetLatestAuditStatus": {
            "get": {
                "tags": [
                    "小程序信息获取"
                ],
                "summary": "获取小程序最后一次审核状态",
                "parameters": [
                    {
                        "type": "string",
                        "description": "小程序APPID",
                        "name": "MiniProgramID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "第三方平台APPID",
                        "name": "PlatformID",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        },
        "/oplatform/{PlatformID}/GetTestQrcode": {
            "get": {
                "tags": [
                    "小程序信息获取"
                ],
                "summary": "获取体验二维码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "小程序APPID",
                        "name": "MiniProgramID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "第三方平台APPID",
                        "name": "PlatformID",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        },
        "/oplatform/{PlatformID}/Release": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "小程序控制"
                ],
                "summary": "发布已审核通过的小程序",
                "parameters": [
                    {
                        "type": "string",
                        "description": "小程序APPID",
                        "name": "MiniProgramID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "第三方平台APPID",
                        "name": "PlatformID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.ApiResult"
                        }
                    }
                }
            }
        },
        "/oplatform/{PlatformID}/SetDomain": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "小程序控制"
                ],
                "summary": "刷新服务器、业务域名",
                "parameters": [
                    {
                        "type": "string",
                        "description": "小程序APPID",
                        "name": "MiniProgramID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "第三方平台APPID",
                        "name": "PlatformID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.ApiResult"
                        }
                    }
                }
            }
        },
        "/oplatform/{PlatformID}/SpeedUpAudit": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "小程序控制"
                ],
                "summary": "审核加急",
                "parameters": [
                    {
                        "description": "入参集合",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.SpeedUpAuditParam"
                        }
                    },
                    {
                        "type": "string",
                        "description": "第三方平台APPID",
                        "name": "PlatformID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.ApiResult"
                        }
                    }
                }
            }
        },
        "/oplatform/{PlatformID}/SubmitAudit": {
            "post": {
                "tags": [
                    "小程序控制"
                ],
                "summary": "提交审核",
                "parameters": [
                    {
                        "type": "string",
                        "description": "小程序APPID",
                        "name": "MiniProgramID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "第三方平台APPID",
                        "name": "PlatformID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.ApiResult"
                        }
                    }
                }
            }
        },
        "/oplatform/{PlatformID}/UndoCodeAudit": {
            "post": {
                "description": "撤销审核 每天1次每月10次",
                "tags": [
                    "小程序控制"
                ],
                "summary": "撤销审核",
                "parameters": [
                    {
                        "type": "string",
                        "description": "小程序APPID",
                        "name": "MiniProgramID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "第三方平台APPID",
                        "name": "PlatformID",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        },
        "/oplatform/{PlatformID}/auth": {
            "get": {
                "tags": [
                    "授权回调"
                ],
                "summary": "授权",
                "parameters": [
                    {
                        "type": "string",
                        "description": "第三方平台APPID",
                        "name": "PlatformID",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        },
        "/oplatform/{PlatformID}/notify": {
            "post": {
                "tags": [
                    "授权回调"
                ],
                "summary": "异步回调",
                "parameters": [
                    {
                        "type": "string",
                        "description": "第三方平台APPID",
                        "name": "PlatformID",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        },
        "/oplatform/{PlatformID}/redirect": {
            "get": {
                "tags": [
                    "授权回调"
                ],
                "summary": "同步回调",
                "parameters": [
                    {
                        "type": "string",
                        "description": "第三方平台APPID",
                        "name": "PlatformID",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        },
        "/oplatform/{PlatformID}/tpl/add": {
            "post": {
                "tags": [
                    "第三方平台模板"
                ],
                "summary": "添加草稿到模板",
                "parameters": [
                    {
                        "type": "string",
                        "description": "草稿编号",
                        "name": "draft_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "第三方平台APPID",
                        "name": "PlatformID",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        },
        "/oplatform/{PlatformID}/tpl/del": {
            "delete": {
                "tags": [
                    "第三方平台模板"
                ],
                "summary": "删除模板",
                "parameters": [
                    {
                        "type": "string",
                        "description": "模板编号",
                        "name": "template_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "第三方平台APPID",
                        "name": "PlatformID",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        },
        "/oplatform/{PlatformID}/tpl/draft": {
            "get": {
                "tags": [
                    "第三方平台模板"
                ],
                "summary": "获取草稿列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "第三方平台APPID",
                        "name": "PlatformID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/string"
                        }
                    }
                }
            }
        },
        "/oplatform/{PlatformID}/tpl/list": {
            "get": {
                "tags": [
                    "第三方平台模板"
                ],
                "summary": "获取模板列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "第三方平台APPID",
                        "name": "PlatformID",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        },
        "/oplatform/{PlatformID}/tpl/pushToAuto": {
            "post": {
                "tags": [
                    "第三方平台模板"
                ],
                "summary": "将模板推送到所有自动升级小程序",
                "parameters": [
                    {
                        "type": "string",
                        "description": "模板编号",
                        "name": "template_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "第三方平台APPID",
                        "name": "PlatformID",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        },
        "/pay/{PayMchID}/notify": {
            "post": {
                "summary": "异步回调",
                "parameters": [
                    {
                        "type": "string",
                        "description": "微信支付商户号",
                        "name": "PayMchID",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        }
    },
    "definitions": {
        "models.Platform": {
            "type": "object",
            "properties": {
                "authRedirectURL": {
                    "description": "用户授权成功回跳地址",
                    "type": "string"
                },
                "bizDomain": {
                    "description": "业务域名",
                    "type": "string"
                },
                "createTime": {
                    "type": "string"
                },
                "platformID": {
                    "description": "平台 appid",
                    "type": "string"
                },
                "platformKey": {
                    "description": "平台 消息解密key",
                    "type": "string"
                },
                "platformSecret": {
                    "description": "平台 appsecret",
                    "type": "string"
                },
                "platformToken": {
                    "description": "平台 token",
                    "type": "string"
                },
                "serverDomain": {
                    "description": "服务器域名",
                    "type": "string"
                },
                "updateTime": {
                    "type": "string"
                }
            }
        },
        "req.CommitCodeParam": {
            "type": "object",
            "required": [
                "miniProgramID",
                "templateID"
            ],
            "properties": {
                "miniProgramID": {
                    "description": "小程序ID",
                    "type": "string"
                },
                "templateID": {
                    "description": "模板编号",
                    "type": "integer"
                }
            }
        },
        "req.ListMiniAppParam": {
            "type": "object",
            "properties": {
                "appID": {
                    "description": "小程序appid",
                    "type": "string"
                },
                "lastId": {
                    "type": "integer"
                },
                "orderClase": {
                    "description": "排序方式，desc倒序， asc正序",
                    "type": "string"
                },
                "orderColumn": {
                    "description": "排序列",
                    "type": "string"
                },
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "platformID": {
                    "description": "开放平台ID",
                    "type": "string"
                },
                "state": {
                    "description": "-1-授权失效 1授权成功，2审核中，3审核通过，4审核失败，5已发布",
                    "type": "integer"
                },
                "version": {
                    "description": "版本号",
                    "type": "string"
                }
            }
        },
        "req.MiniAppConfigParam": {
            "type": "object",
            "required": [
                "appID"
            ],
            "properties": {
                "appID": {
                    "description": "小程序appid",
                    "type": "string"
                },
                "autoAudit": {
                    "description": "自动提审(升级) -1否 1是",
                    "type": "integer"
                },
                "autoRelease": {
                    "description": "自动发布-1否 1是",
                    "type": "integer"
                },
                "extConfig": {
                    "description": "小程序扩展配置，发布时会注入至ext.json",
                    "type": "string"
                },
                "mchID": {
                    "description": "支付商户号id",
                    "type": "string"
                },
                "secret": {
                    "description": "小程序secret",
                    "type": "string"
                }
            }
        },
        "req.PageParam": {
            "type": "object",
            "properties": {
                "lastId": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                }
            }
        },
        "req.PlatformParam": {
            "type": "object",
            "required": [
                "platformID"
            ],
            "properties": {
                "authRedirectURL": {
                    "description": "用户授权成功回跳地址",
                    "type": "string"
                },
                "bizDomain": {
                    "description": "业务域名",
                    "type": "string"
                },
                "platformID": {
                    "description": "平台 appid",
                    "type": "string"
                },
                "platformKey": {
                    "description": "平台 消息解密key",
                    "type": "string"
                },
                "platformSecret": {
                    "description": "平台 appsecret",
                    "type": "string"
                },
                "platformToken": {
                    "description": "平台 token",
                    "type": "string"
                },
                "serverDomain": {
                    "description": "服务器域名",
                    "type": "string"
                }
            }
        },
        "req.SpeedUpAuditParam": {
            "type": "object",
            "required": [
                "auditID",
                "miniProgramID"
            ],
            "properties": {
                "auditID": {
                    "description": "要加急的审核单号",
                    "type": "integer"
                },
                "miniProgramID": {
                    "description": "小程序ID",
                    "type": "string"
                }
            }
        },
        "resp.ApiResult": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "resp.MiniAppListsVO": {
            "type": "object",
            "properties": {
                "appID": {
                    "description": "小程序appid",
                    "type": "string"
                },
                "auditID": {
                    "description": "审核编号",
                    "type": "integer"
                },
                "autoAudit": {
                    "description": "自动提审(升级) -1否 1是",
                    "type": "integer"
                },
                "autoRelease": {
                    "description": "自动发布-1否 1是",
                    "type": "integer"
                },
                "createTime": {
                    "type": "string"
                },
                "extConfig": {
                    "description": "小程序扩展配置，发布时会注入至ext.json",
                    "type": "string"
                },
                "mchID": {
                    "description": "支付商户号id",
                    "type": "string"
                },
                "originalID": {
                    "description": "原始ID",
                    "type": "string"
                },
                "platformID": {
                    "description": "开放平台ID",
                    "type": "string"
                },
                "refreshToken": {
                    "description": "接口调用凭据刷新令牌",
                    "type": "string"
                },
                "secret": {
                    "description": "小程序secret",
                    "type": "string"
                },
                "state": {
                    "description": "-1-授权失效 1授权成功，2审核中，3审核通过，4审核失败，5已发布",
                    "type": "integer"
                },
                "templateListen": {
                    "description": "模板监听开发小程序(appid)",
                    "type": "string"
                },
                "updateTime": {
                    "type": "string"
                },
                "version": {
                    "description": "当前版本",
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
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "纸喵 wechat API",
	Description: "纸喵软件系列之服务端",
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
