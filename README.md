# Golang OAuth 2.0 Server

> 一种开放协议，允许web、移动和桌面应用程序以简单和标准的方法进行安全授权。

[![Build][build-status-image]][build-status-url] [![Codecov][codecov-image]][codecov-url] [![ReportCard][reportcard-image]][reportcard-url] [![GoDoc][godoc-image]][godoc-url] [![License][license-image]][license-url]

## 声明

- 这个是国外大佬的OAuth2修改版，主要是加入一些没有的功能，遵循MIT开源协议
- 1. client加密（已实现）
- 2. client存储在mysql，（已实现）
- 3. 已申请token的client，将直接返回申请的token，而不是再去申请（已实现）
- 4. 持久化login的数据（需自己实现）
- ps：还缺少，登陆时图形验证码（需自己实现）
- 改的不是很好，可以作为参考

## Protocol Flow

```text
     +--------+                               +---------------+
     |        |--(A)- Authorization Request ->|   Resource    |
     |        |                               |     Owner     |
     |        |<-(B)-- Authorization Grant ---|               |
     |        |                               +---------------+
     |        |
     |        |                               +---------------+
     |        |--(C)-- Authorization Grant -->| Authorization |
     | Client |                               |     Server    |
     |        |<-(D)----- Access Token -------|               |
     |        |                               +---------------+
     |        |
     |        |                               +---------------+
     |        |--(E)----- Access Token ------>|    Resource   |
     |        |                               |     Server    |
     |        |<-(F)--- Protected Resource ---|               |
     +--------+                               +---------------+
```

## Quick Start

### Download and install

```bash
go get gitee.com/w1134407270/go-oauth2-mysql-v4
```

### Create file `server.go`

```go
package main

func main() {
	manager := manage.NewDefaultManager()
	// 设置，默认的认证code配置
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	// 设置code过期时间
	manager.SetAuthorizeCodeExp(time.Second * 10)

	// 使用jwt，或者其它
	// manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodHS512))
	manager.MapAccessGenerate(generates.NewAccessGenerate())

	// token使用redis存储
	redisTokenStore := oredis.NewRedisStore(&redis.Options{
		Addr:     "xxx:6379",
		Password: "xxx",
		DB:       0,
	})
	manager.MapTokenStorage(redisTokenStore)

	// client使用mysql存储，初始化clientStore的Mysql存储对象
	dbs, err := sql.Open("mysql", "xxx:xxx@tcp(xxx:3306)/xxx?charset=utf8&parseTime=true")
	if err != nil {
		log.Fatal("mysql连接失败")
	}
	clientMySqlStore := store.NewClientMySqlStore(dbs, "custom_table_name")
	// 新的改动。只实现ClientStore的GetById接口，MapClientBcryptSecretStorage是新增的，第二个参数代表使用了bcrypt加密client_Secret
	manager.MapClientBcryptSecretStorage(clientMySqlStore, true)

	// 默认的server配置，和添加管理器
	srv := server.NewServer(server.NewConfig(), manager)
	
	// 设置用户名和密码的处理，这里可以使用MySql，取出username和password
	srv.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {
		if username == "test" && password == "test" {
			userID = "test"
		}
		return
	})

	// 设置允许获取访问请求
	srv.SetAllowGetAccessRequest(true)
	// 设置客户端信息处理
	srv.SetClientInfoHandler(server.ClientFormHandler)

	// 错误的处理
	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})
	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	// 授权处理HandleAuthorizeRequest
	http.HandleFunc("/oauth/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	// token处理HandleTokenRequest
	http.HandleFunc("/oauth/token", func(w http.ResponseWriter, r *http.Request) {
		srv.HandleTokenRequest(w, r)
	})
	
	// 开启服务
	log.Fatal(http.ListenAndServe(":9096", nil))
}

```

### Build and run

```bash
go build server.go

./server
```

### Open in your web browser

[http://localhost:9096/token?grant_type=client_credentials&client_id=000000&client_secret=999999&scope=read](http://localhost:9096/token?grant_type=client_credentials&client_id=000000&client_secret=999999&scope=read)

```json
{
  "access_token": "J86XVRYSNFCFI233KXDL0Q",
  "expires_in": 7200,
  "scope": "read",
  "token_type": "Bearer"
}
```

## Features

- Easy to use
- Based on the [RFC 6749](https://tools.ietf.org/html/rfc6749) implementation
- Token storage support TTL
- Support custom expiration time of the access token
- Support custom extension field
- Support custom scope
- Support jwt to generate access tokens

## Example

> A complete example of simulation authorization code model

Simulation examples of authorization code model, please check [example](/example)

### Use jwt to generate access tokens

```go

import (
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/dgrijalva/jwt-go"
)

// ...
manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodHS512))

// Parse and verify jwt access token
token, err := jwt.ParseWithClaims(access, &generates.JWTAccessClaims{}, func(t *jwt.Token) (interface{}, error) {
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("parse error")
	}
	return []byte("00000000"), nil
})
if err != nil {
	// panic(err)
}

claims, ok := token.Claims.(*generates.JWTAccessClaims)
if !ok || !token.Valid {
	// panic("invalid token")
}
```

## Store Implements

- [BuntDB](https://github.com/tidwall/buntdb)(default store)
- [Redis](https://github.com/go-oauth2/redis)
- [MongoDB](https://github.com/go-oauth2/mongo)
- [MySQL](https://github.com/go-oauth2/mysql)
- [MySQL (Provides both client and token store)](https://github.com/imrenagi/go-oauth2-mysql)
- [PostgreSQL](https://github.com/vgarvardt/go-oauth2-pg)
- [DynamoDB](https://github.com/contamobi/go-oauth2-dynamodb)
- [XORM](https://github.com/techknowlogick/go-oauth2-xorm)
- [XORM (MySQL, client and token store)](https://github.com/rainlay/go-oauth2-xorm)
- [GORM](https://github.com/techknowlogick/go-oauth2-gorm)
- [Firestore](https://github.com/tslamic/go-oauth2-firestore)

## Handy Utilities

- [OAuth2 Proxy Logger (Debug utility that proxies interfaces and logs)](https://github.com/aubelsb2/oauth2-logger-proxy)

## MIT License

Copyright (c) 2021 Lyric

[build-status-url]: https://travis-ci.org/go-oauth2/oauth2
[build-status-image]: https://travis-ci.org/go-oauth2/oauth2.svg?branch=master
[codecov-url]: https://codecov.io/gh/go-oauth2/oauth2
[codecov-image]: https://codecov.io/gh/go-oauth2/oauth2/branch/master/graph/badge.svg
[reportcard-url]: https://goreportcard.com/report/github.com/go-oauth2/oauth2/v4
[reportcard-image]: https://goreportcard.com/badge/github.com/go-oauth2/oauth2/v4
[godoc-url]: https://godoc.org/github.com/go-oauth2/oauth2/v4
[godoc-image]: https://godoc.org/github.com/go-oauth2/oauth2/v4?status.svg
[license-url]: http://opensource.org/licenses/MIT
[license-image]: https://img.shields.io/npm/l/express.svg
