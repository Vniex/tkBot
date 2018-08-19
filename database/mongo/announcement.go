package mongo

import (
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Announcement struct {
	Id int `bson:"id" json:"id"`
	Title string `bson:"title" json:"title"`
	Abstract  string `bson:"abstract" json:"abstract"`
	Link string `bson:"link" json:"link"`
	ExchangeId int `bson:"exchange_id" json:"exchange_id"`
	Logo string `bson:"logo" json:"logo"`
	ExchangeName string `bson:"exchange_name" json:"exchange_name"`
	ExchangeAlias string `bson:"exchange_alias" json:"exchange_alias"`
	PostedAt int `bson:"posted_at" json:"posted_at"`
	LangType string `bson:"lang_type" json:"lang_type"`
	Source string `bson:"source" json:"source"`

}
func NewAnnouncement(id int ,title ,abstract,link string,exchange_id int,logo,exchange_name,exchange_alias string,posted_at int,lang_type,source string) *Announcement{

	return &Announcement{
		id,
		title,
		abstract,
		link,
		exchange_id,
		logo,
		exchange_name,
		exchange_alias,
		posted_at,
		lang_type,
		source,
	}
}

type AnnouncementDB struct {
	Database string
	Collection string

}
func NewAnnouncementDB() *AnnouncementDB{
	return &AnnouncementDB{
		Database,
		AnnouncementCollection,
	}
}

func (a *AnnouncementDB)Insert(ann *Announcement) error{
	session_clone,err:=GetSessionClone()
	if err != nil {
		log.Println(err)
		return err
	}
	defer session_clone.Clone()
	err=session_clone.DB(a.Database).C(a.Collection).Insert(ann)
	if err!=nil{
		log.Println(err)
		return err
	}
	return nil
}

func (a *AnnouncementDB)DeleteById(ann_id int) error{
	session_clone,err:=GetSessionClone()
	if err != nil {
		log.Println(err)
		return err
	}
	defer session_clone.Clone()
	err=session_clone.DB(a.Database).C(a.Collection).Remove(bson.M{"id": ann_id})
	if err!=nil{
		log.Println(err)
		return err
	}
	return nil
}


func (a *AnnouncementDB) FindOneById(ann_id int) (*Announcement,error) {
	result := &Announcement{}
	session_clone,err:=GetSessionClone()
	if err != nil {
		log.Println(err)
		return nil,err
	}
	defer session_clone.Clone()
	err = session_clone.DB(a.Database).C(a.Collection).Find(bson.M{"id": ann_id}).One(&result)
	if err != nil {
		return nil,err
	}

	return result,nil
}

func (a *AnnouncementDB) FindAnnsByExchangeId(exchange_id int)([]*Announcement,error){
	var result []*Announcement
	session_clone,err:=GetSessionClone()
	if err != nil {
		log.Println(err)
		return nil,err
	}
	defer session_clone.Clone()
	err = session_clone.DB(a.Database).C(a.Collection).Find(bson.M{"exchange_id":exchange_id}).Sort("-posted_at").All(&result)
	if err != nil {
		return nil,err
	}

	return result,nil
}

func (a *AnnouncementDB) FindAnns(limit int)([]*Announcement,error){
	var result []*Announcement
	session_clone,err:=GetSessionClone()
	if err != nil {
		log.Println(err)
		return nil,err
	}
	defer session_clone.Clone()
	err = session_clone.DB(a.Database).C(a.Collection).Find(nil).Sort("-posted_at").Limit(limit).All(&result)
	if err != nil {
		return nil,err
	}

	return result,nil
}