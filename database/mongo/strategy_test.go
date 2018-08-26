package mongo

import (
	"testing"
	"gopkg.in/mgo.v2/bson"
)

func TestStrategyDB_Insert(t *testing.T) {
	para:=make(map[string]interface{})
	para["interval"]="5"
	s:=&Strategy{
		bson.NewObjectId(),
		"annSipder",

		para,
	}
	db:=GetStrategyDB()
	t.Log(db.Insert(s))
}

func TestStrategyDB_FindStrategies(t *testing.T) {
	db:=GetStrategyDB()
	ss,_:=db.FindStrategies()
	for _,s :=range ss{
		t.Log(s.Id.Hex())
		t.Log(s.StrategyName)
		t.Log(s.Para)
	}
}