package mongo

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

var logInfoDB *LogInfoDB


const (
	INFO =iota
	WARN
	ERROR
	FATAL
)

type LogInfo struct {
	RobotName string `bson:"robot_name" json:"robot_name"`
	TimeStamp int `bson:"time_stamp" json:"time_stamp"`
	Level int 	`bson:"level" json:"level"`
	Msg string `bson:"msg" json:"msg"`

}


func NewLogInfo(robot_name string ,timestamp int,level int,msg string)*LogInfo{

	return &LogInfo{
		robot_name,
		timestamp,
		level,
		msg,
	}
}

type LogInfoDB struct {
	Database string
	Collection string

}

func GetLogInfoDB() *LogInfoDB{
	if logInfoDB==nil{
		logInfoDB=&LogInfoDB{
			Database,
			LogInfoCollection,
		}
	}
	return logInfoDB
}

func (db *LogInfoDB)Insert(logInfo *LogInfo) error{
	session_clone,err:=GetSessionClone()
	if err != nil {
		log.Error(err)
		return err
	}
	defer session_clone.Clone()
	err=session_clone.DB(db.Database).C(db.Collection).Insert(logInfo)
	if err!=nil{
		log.Error(err)
		return err
	}
	return nil
}


func (db *LogInfoDB)FindlastestLogInfoByRobot(robot_name string,limit int) ([]*LogInfo,error) {
	var result []*LogInfo
	session_clone,err:=GetSessionClone()
	if err != nil {
		log.Error(err)
		return nil,err
	}
	defer session_clone.Clone()
	err = session_clone.DB(db.Database).C(db.Collection).Find(bson.M{"robot_name":robot_name}).Sort("time_stamp").Limit(limit).All(&result)
	if err != nil {
		log.Error(err)
		return nil,err
	}

	return result,nil
}
