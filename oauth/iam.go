package oauth

import (
	"crypto/tls"
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

//iam accesstoken 数据
type IAMAccessToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	CreatedAt    int    `json:"created_at"`
}

//iam 用户数据
//用户使用iam登录的时候，直接根据iam的id获取数据
type IAMUser struct {
	Id       int    `json:"id"`                      //用户id
	MemberId int    `json:"member_id"`               //绑定的用户id
	Email    string `json:"email" orm:"size(50)"`    //电子邮箱
	Login    string `json:"username" orm:"size(50)"` //用户名
	Name     string `json:"name" orm:"size(50)"`     //昵称
}

//获取accessToken
func GetIAMAccessToken(code string) (token IAMAccessToken, err error) {
	var resp string
	Api := beego.AppConfig.String("oauth::IAMAccesstoken")
	ClientId := beego.AppConfig.String("oauth::IAMClientId")
	ClientSecret := beego.AppConfig.String("oauth::IAMClientSecret")
	Callback := beego.AppConfig.String("oauth::IAMCallback")
	req := httplib.Post(Api)
	if strings.HasPrefix(Api, "https") {
		req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	req.Param("grant_type", "authorization_code")
	req.Param("code", code)
	req.Param("client_id", ClientId)
	req.Param("redirect_uri", Callback)
	req.Param("client_secret", ClientSecret)
	if resp, err = req.String(); err == nil {
		err = json.Unmarshal([]byte(resp), &token)
	}
	return
}

//获取用户信息
func GetIAMUserInfo(accessToken string) (info IAMUser, err error) {
	beego.Debug(accessToken)
	var resp string
	Api := beego.AppConfig.String("oauth::IAMUserInfo")
	req := httplib.Get(Api)
	if strings.HasPrefix(Api, "https") {
		req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		req.Header("Authorization", "Bearer "+accessToken)
	}
	if resp, err = req.String(); err == nil {
		beego.Debug(resp)
		err = json.Unmarshal([]byte(resp), &info)
	}
	return
}
