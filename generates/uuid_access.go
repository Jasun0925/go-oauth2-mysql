package generates

import (
	"bytes"
	"context"
	"encoding/base64"
	"github.com/Jasun0925/go-oauth2-mysql"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// NewAccessGenerate create to generate the access token instance
func NewAccessGenerate() *UUIDAccessGenerate {
	return &UUIDAccessGenerate{}
}

// AccessGenerate generate the access token
type UUIDAccessGenerate struct {
}

// Token based on the UUID generated token，基于UUID生成的令牌
func (ag *UUIDAccessGenerate) Token(ctx context.Context, data *oauth2.GenerateBasic, isGenRefresh bool) (string, string, error) {
	buf := bytes.NewBufferString(data.Client.GetID())
	buf.WriteString(data.UserID)
	buf.WriteString(strconv.FormatInt(data.CreateAt.UnixNano(), 10))

	access := base64.URLEncoding.EncodeToString([]byte(uuid.NewMD5(uuid.Must(uuid.NewRandom()), buf.Bytes()).String()))
	access = strings.ToUpper(strings.TrimRight(access, "="))
	refresh := ""
	if isGenRefresh {
		refresh = base64.URLEncoding.EncodeToString([]byte(uuid.NewSHA1(uuid.Must(uuid.NewRandom()), buf.Bytes()).String()))
		refresh = strings.ToUpper(strings.TrimRight(refresh, "="))
	}

	return access, refresh, nil
}
