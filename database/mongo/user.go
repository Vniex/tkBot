package mongo

import (
	"gopkg.in/mgo.v2/bson"
	log "github.com/sirupsen/logrus"
	"time"
	Utils "tkBot/utils"

	"errors"
)

var userDB *UserDB

type User struct {
	Id bson.ObjectId `bson:"_id" json:"id"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
	SecretPassword string `bson:"secret_password" json:"secret_password"`
	CreatetTime int64 `bson:"createt_time" json:"createt_time"`
	SecretKey string `bson:"secret_key" json:"secret_key"`
}


func NewUser(username,password string)*User{
	createTimeStamp:=time.Now().Unix()
	secretPassword:=Utils.GetSHA256(username+password+string(createTimeStamp))
	secretKey:=Utils.GetSHA256(secretPassword)
	return &User{
		bson.NewObjectId(),
		username,
		password,
		secretPassword,
		createTimeStamp,
		secretKey,
	}
}

type UserDB struct {
	Database string
	Collection string

}

func GetUserDB() *UserDB{
	if userDB==nil{
		userDB=&UserDB{
			Database,
			UserCollection,
		}
	}
	return userDB
}

func (u *UserDB)Insert(user *User) error{
	session_clone,err:=GetSessionClone()
	if err != nil {
		log.Error(err)
		return err
	}
	defer session_clone.Clone()
	err=session_clone.DB(u.Database).C(u.Collection).Insert(user)
	if err!=nil{
		log.Error(err)
		return err
	}
	return nil
}

func (u *UserDB)FindUserByUserName(username string) (*User,error) {
	result := &User{}
	session_clone,err:=GetSessionClone()
	if err != nil {
		log.Error(err)
		return nil,err
	}
	defer session_clone.Clone()
	err = session_clone.DB(u.Database).C(u.Collection).Find(bson.M{"username": username}).One(&result)
	if err != nil {
		log.Error(err)
		return nil,err
	}
	result.Password="******"
	return result,nil
}

func (u *UserDB)FindUserById(id string) (*User,error) {
	if !bson.IsObjectIdHex(id){
		return nil,errors.New("mongo id error")
	}
	result := &User{}
	session_clone,err:=GetSessionClone()
	if err != nil {
		log.Error(err)
		return nil,err
	}
	defer session_clone.Clone()
	err = session_clone.DB(u.Database).C(u.Collection).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
	if err != nil {
		log.Error(err)
		return nil,err
	}
	result.Password="******"
	return result,nil
}