package server

import (
	"github.com/go-oauth2-mysql"
)

// SetTokenType token type
func (s *Server) SetTokenType(tokenType string) {
	s.Config.TokenType = tokenType
}

// SetAllowGetAccessRequest to allow GET requests for the token
func (s *Server) SetAllowGetAccessRequest(allow bool) {
	s.Config.AllowGetAccessRequest = allow
}

// SetAllowedResponseType allow the authorization types
func (s *Server) SetAllowedResponseType(types ...oauth2.ResponseType) {
	s.Config.AllowedResponseTypes = types
}

// SetAllowedGrantType 允许授予类型
func (s *Server) SetAllowedGrantType(types ...oauth2.GrantType) {
	s.Config.AllowedGrantTypes = types
}

// SetClientInfoHandler 从请求中获取客户端信息
func (s *Server) SetClientInfoHandler(handler ClientInfoHandler) {
	s.ClientInfoHandler = handler
}

// SetClientAuthorizedHandler 检查允许使用此授权授予类型的客户端
func (s *Server) SetClientAuthorizedHandler(handler ClientAuthorizedHandler) {
	s.ClientAuthorizedHandler = handler
}

// SetClientScopeHandler 检查客户端允许使用范围
func (s *Server) SetClientScopeHandler(handler ClientScopeHandler) {
	s.ClientScopeHandler = handler
}

// SetUserAuthorizationHandler 从请求授权中获取用户id
func (s *Server) SetUserAuthorizationHandler(handler UserAuthorizationHandler) {
	s.UserAuthorizationHandler = handler
}

// SetPasswordAuthorizationHandler 从用户名和密码中获取用户id
func (s *Server) SetPasswordAuthorizationHandler(handler PasswordAuthorizationHandler) {
	s.PasswordAuthorizationHandler = handler
}

// SetRefreshingScopeHandler 检查刷新令牌的范围
func (s *Server) SetRefreshingScopeHandler(handler RefreshingScopeHandler) {
	s.RefreshingScopeHandler = handler
}

// SetRefreshingValidationHandler 检查refresh_token是否仍然有效。没有撤销或其他
func (s *Server) SetRefreshingValidationHandler(handler RefreshingValidationHandler) {
	s.RefreshingValidationHandler = handler
}

// SetResponseErrorHandler response error handling
func (s *Server) SetResponseErrorHandler(handler ResponseErrorHandler) {
	s.ResponseErrorHandler = handler
}

// SetInternalErrorHandler 内部错误处理
func (s *Server) SetInternalErrorHandler(handler InternalErrorHandler) {
	s.InternalErrorHandler = handler
}

// SetExtensionFieldsHandler 响应带有字段扩展的访问令牌
func (s *Server) SetExtensionFieldsHandler(handler ExtensionFieldsHandler) {
	s.ExtensionFieldsHandler = handler
}

// SetAccessTokenExpHandler 设置访问令牌的过期日期
func (s *Server) SetAccessTokenExpHandler(handler AccessTokenExpHandler) {
	s.AccessTokenExpHandler = handler
}

// SetAuthorizeScopeHandler 设置访问令牌的范围
func (s *Server) SetAuthorizeScopeHandler(handler AuthorizeScopeHandler) {
	s.AuthorizeScopeHandler = handler
}
