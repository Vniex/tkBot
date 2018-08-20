package mongo

import "log"

type Strategy struct {
	Id int `bson:"id" json:"id"`
	StrategyName string `bson:"strategy_name" json:"strategy_name"`
	Desp string `bson:"desp"  json:"desp"`
	Para map[string]interface{} `bson:"para" json:"para"`
}



type StrategyDB struct {
	Database string
	Collection string

}
func NewStrategyDB() *StrategyDB{
	return &StrategyDB{
		Database,
		StrategyCollection,
	}
}

func (db *StrategyDB)Insert(s *Strategy) error{
	session_clone,err:=GetSessionClone()
	if err != nil {
		log.Println(err)
		return err
	}
	defer session_clone.Clone()
	err=session_clone.DB(db.Database).C(db.Collection).Insert(s)
	if err!=nil{
		log.Println(err)
		return err
	}
	return nil
}

func (db *StrategyDB) FindStrategies()([]*Strategy,error){
	var result []*Strategy
	session_clone,err:=GetSessionClone()
	if err != nil {
		log.Println(err)
		return nil,err
	}
	defer session_clone.Clone()
	err = session_clone.DB(db.Database).C(db.Collection).Find(nil).All(&result)
	if err != nil {
		return nil,err
	}

	return result,nil
}