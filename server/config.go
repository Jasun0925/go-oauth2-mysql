package server

import (
	"github.com/Jasun0925/go-oauth2-mysql"
	"net/http"
	"time"
)

// Config configuration parameters 配置oauth请求的参数
type Config struct {
	TokenType                   string                // token type
	AllowGetAccessRequest       bool                  // to allow GET requests for the token
	AllowedResponseTypes        []oauth2.ResponseType // allow the authorization type
	AllowedGrantTypes           []oauth2.GrantType    // allow the grant type
	AllowedCodeChallengeMethods []oauth2.CodeChallengeMethod
	ForcePKCE                   bool
}

// NewConfig create to configuration instance 创建到配置实例
func NewConfig() *Config {
	return &Config{
		TokenType:            "Bearer",
		AllowedResponseTypes: []oauth2.ResponseType{oauth2.Code, oauth2.Token},
		AllowedGrantTypes: []oauth2.GrantType{
			oauth2.AuthorizationCode,
			oauth2.PasswordCredentials,
			oauth2.ClientCredentials,
			oauth2.Refreshing,
		},
		AllowedCodeChallengeMethods: []oauth2.CodeChallengeMethod{
			oauth2.CodeChallengePlain,
			oauth2.CodeChallengeS256,
		},
	}
}

// AuthorizeRequest authorization request 授权请求
type AuthorizeRequest struct {
	ResponseType        oauth2.ResponseType
	ClientID            string
	Scope               string
	RedirectURI         string
	State               string
	UserID              string
	CodeChallenge       string
	CodeChallengeMethod oauth2.CodeChallengeMethod
	AccessTokenExp      time.Duration
	Request             *http.Request
}
