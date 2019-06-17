package wechat

import (
	"fmt"
	"net/url"

	"github.com/esap/wechat/util"
)

// WXAPIOauth2 oauth2鉴权
const (
	WXAPIOauth2           = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%v&redirect_uri=%v&response_type=code&scope=snsapi_base&state=110#wechat_redirect"
	WXAPIOauth2token      = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%v&secret=%v&code=%v&grant_type=authorization_code"
	WXAPIJscode2session   = "https://api.weixin.qq.com/sns/jscode2session?appid=%v&secret=%v&js_code=%v&grant_type=authorization_code"
	CorpAPIJscode2session = "https://qyapi.weixin.qq.com/cgi-bin/miniprogram/jscode2session?access_token=%v&js_code=%v&grant_type=authorization_code"
	//开放平台
	OpenApiOauth2 = "https://open.weixin.qq.com/connect/qrconnect?appid=%v&redirect_uri=%v&response_type=code&scope=snsapi_login&state=110#wechat_redirect"
)

// WxSession 兼容企业微信和服务号
type WxSession struct {
	WxErr
	SessionKey string `json:"session_key"`
	// corp
	CorpId string `json:"corpid"`
	UserId string `json:"userid"`
	// mp
	OpenId  string `json:"openid"`
	UnionId string `json:"unionid"`
}

// GetOauth2Url 获取鉴权页面
func GetOauth2Url(corpId, host string) string {
	return fmt.Sprintf(WXAPIOauth2, corpId, url.QueryEscape(host))
}

// 开放平台扫码登录页面
func GetOpenOauth2Url(corpId, host string) string {
	return fmt.Sprintf(OpenApiOauth2, corpId, url.QueryEscape(host))
}

// 获取OAuth AccessToken的结果 如果错误，返回结果{"errcode":40029,"errmsg":"invalid code"}
type OAuthAccessTokenResult struct {
	WxErr
	Access_token string `json:"access_token"`
	// corp
	Expires_in    int64  `json:"expires_in"`
	Refresh_token string `json:"Refresh_token"`
	// mp
	OpenId  string `json:"openid"`
	Scope   string `json:"scope"`
	UnionId string `json:"unionid"`
}

// 微信网页授权 通过code换取网页授权access_token
func (s *Server) Code2token(code string) (ws *OAuthAccessTokenResult, err error) {
	url := fmt.Sprintf(WXAPIOauth2token, s.AppId, s.Secret, code)
	ws = new(OAuthAccessTokenResult)
	err = util.GetJson(url, ws)
	if ws.Error() != nil {
		err = ws.Error()
	}
	return
}

// Jscode2Session code换session
func (s *Server) Jscode2Session(code string) (ws *WxSession, err error) {
	url := fmt.Sprintf(WXAPIJscode2session, s.AppId, s.Secret, code)
	ws = new(WxSession)
	err = util.GetJson(url, ws)

	if ws.Error() != nil {
		err = ws.Error()
	}
	return
}

// Jscode2SessionEnt code换session（企业微信）
func (s *Server) Jscode2SessionEnt(code string) (ws *WxSession, err error) {
	url := fmt.Sprintf(CorpAPIJscode2session, s.GetAccessToken(), code)
	ws = new(WxSession)
	err = util.GetJson(url, ws)

	if ws.Error() != nil {
		err = ws.Error()
	}
	return
}
