package utils

import "testing"

func Test_message_wx(t *testing.T) {
	SendToWechat("","test text" ,"test desp")
}

func TestToFloat64(t *testing.T) {
	res:=ToFloat64(11)
	t.Log(res)
}

func TestAdjustFloat(t *testing.T) {
	res,_:=AdjustFloat(15.23682,-1,true)
	t.Log(res)
}