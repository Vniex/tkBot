package global

const (
	SUCCESS = iota
	ERROR_NO_FILE
	ERROR_INTERNAL
	ERROR_PARAMETER
	ERROR_AUTH_CHECK_TOKEN_FAIL
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT
	ERROR_AUTH_TOKEN
	ERROR_AUTH
	ERROR_TIMESTAMP


)

func GetMsg(code int)string{
	var msg = ""
	switch code {
	case SUCCESS:msg="成功"
	case ERROR_NO_FILE: msg="没有该文件"
	case ERROR_PARAMETER: msg="参数错误"
	case ERROR_AUTH_CHECK_TOKEN_FAIL: msg="token验证错误"
	case ERROR_AUTH_CHECK_TOKEN_TIMEOUT : msg="token已过期"
	case ERROR_AUTH:msg="token错误"
	case ERROR_AUTH_TOKEN: msg="Token生成失败"
	case ERROR_TIMESTAMP:msg="timestamp错误"


	}
	return msg
}