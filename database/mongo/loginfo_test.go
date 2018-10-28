package mongo

import (
	"testing"
	"time"
)

func TestLogInfoDB_Insert(t *testing.T) {
	logInfo:=NewLogInfo(
		"test",
		int(time.Now().Unix()),
		INFO,
		"test",
	)
	db:=GetLogInfoDB()
	t.Log(db.Insert(logInfo))
}

func TestLogInfoDB_FindlastestLogInfo(t *testing.T) {
	db:=GetLogInfoDB()
	res,err:=db.FindlastestLogInfoByRobot("test",10)
	t.Log(err)
	t.Log(res)
	t.Log(res[0].RobotName)
	t.Log(res[0].TimeStamp)
	t.Log(res[0].Level)
	t.Log(res[0].Msg)
}
