// Package wechat provides constants for using OAuth2 to access Wechat.
package wechat

import (
	"golang.org/x/oauth2"
)

// Endpoint is Github's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://open.weixin.qq.com/connect/oauth2/authorize",
	TokenURL: "https://api.weixin.qq.com/sns/oauth2/access_token",
}
