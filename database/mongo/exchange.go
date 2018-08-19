package mongo

import (
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Exchange struct {
	ExchangeId int `bson:"exchange_id"`
	ExchangeName string `bson:"exchange_name"`
	ExchangeAlias string `bson:"exchange_alias"`
	Logo string `bson:"logo"`
	AnnouncementList []int `bson:"announcement_list"`
}

func NewExchange(id int ,name ,alias,logo string,anns ...int) *Exchange{
	var ann_list []int
	for _,ann:= range anns{
		ann_list=append(ann_list, ann)
	}
	return &Exchange{
		id,
		name,
		alias,
		logo,
		ann_list,
	}
}

type ExchangeDB struct {
	Database string
	Collection string

}
 func NewExchangeDB() *ExchangeDB{
 	return &ExchangeDB{
 		Database,
 		ExchangeCollection,
	}
 }

func (e *ExchangeDB)Insert(exchange *Exchange) error{
	session_clone,err:=GetSessionClone()
	if err != nil {
		log.Println(err)
		return err
	}
	defer session_clone.Clone()
	err=session_clone.DB(e.Database).C(e.Collection).Insert(exchange)
	if err!=nil{
		log.Println(err)
		return err
	}
	return nil
}

func (e *ExchangeDB)DeleteById(exchange_id int) error{
	session_clone,err:=GetSessionClone()
	if err != nil {
		log.Println(err)
		return err
	}
	defer session_clone.Clone()
	err=session_clone.DB(e.Database).C(e.Collection).Remove(bson.M{"exchange_id": exchange_id})
	if err!=nil{
		log.Println(err)
		return err
	}
	return nil
}



func (e  *ExchangeDB) AddAnnouncement(exchange_id int, ann_id int) (err error) {
	session_clone,err:=GetSessionClone()
	if err != nil {
		log.Println(err)
		return err
	}
	defer session_clone.Clone()
	query := bson.M{"exchange_id": exchange_id}
	update := bson.M{"$push": bson.M{"announcement_list": ann_id}}
	err =session_clone.DB(e.Database).C(e.Collection).Update(query, update)
	if err!=nil{
		log.Println(err)
		return err
	}
	return nil
}


func (e *ExchangeDB) FindOneByName(exchange_name string) (*Exchange,error) {
	result := &Exchange{}
	session_clone,err:=GetSessionClone()
	if err != nil {
		log.Println(err)
		return nil,err
	}
	defer session_clone.Clone()
	err = session_clone.DB(e.Database).C(e.Collection).Find(bson.M{"exchange_name": exchange_name}).One(&result)
	if err != nil {
		return nil,err
	}

	return result,nil
}

func (e *ExchangeDB) FindOneById(exchange_id int) (*Exchange,error) {
	result := &Exchange{}
	session_clone,err:=GetSessionClone()
	if err != nil {
		log.Println(err)
		return nil,err
	}
	defer session_clone.Clone()
	err = session_clone.DB(e.Database).C(e.Collection).Find(bson.M{"exchange_id": exchange_id}).One(&result)
	if err != nil {
		return nil,err
	}

	return result,nil
}
