package mongo

import "testing"

func TestStrategyDB_Insert(t *testing.T) {
	para:=make(map[string]string)
	para["interval"]="5"
	s:=&Strategy{
		1,
		"announctment spider",
		para,
	}
	db:=NewStrategyDB()
	t.Log(db.Insert(s))
}

func TestStrategyDB_FindStrategies(t *testing.T) {
	db:=NewStrategyDB()
	ss,_:=db.FindStrategies()
	for _,s :=range ss{
		t.Log(s.Id)
		t.Log(s.Desp)
		t.Log(s.Para)
	}
}