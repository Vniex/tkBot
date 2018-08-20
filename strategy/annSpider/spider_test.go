package annSpider

import (
	"testing"
	 Utils "tkBot/utils"
)

var para=&Parameters{"test",5*300,5}

func TestStrategy_Fetch(t *testing.T) {
	s:=NewStrategy(para)
	s.Run()
}



func TestRegister(t *testing.T) {

	t.Log(Register())
}


func TestStrategy_Run(t *testing.T) {
	ppp:=make(map[string]interface{})
	ppp["robot_name"]="test_ann_spider"
	ppp["interval"]=5*60
	ppp["retry"]=5
	var taskPara Parameters
	Utils.Map2Struct(ppp,&taskPara,"json")
	task:=NewStrategy(&taskPara)
	t.Log(task.RobotDetect)
	task.Run()

}