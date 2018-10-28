package utils

import (
	"net/url"
	"net/http"

)

func SendToWechat(SERVER_SCKEY,text,desp string) {

	wechatUrl:="https://sc.ftqq.com/"+SERVER_SCKEY+".send"
	params := url.Values{}
	params.Set("text",text)
	params.Set("desp",desp)
	http.PostForm(wechatUrl,params)
}

