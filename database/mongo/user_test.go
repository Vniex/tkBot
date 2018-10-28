package mongo

import (
	"gopkg.in/mgo.v2/bson"
	"testing"
	"time"
	"tkBot/utils"
)

func TestUserDB_Insert(t *testing.T) {

	u:=&User{
		bson.NewObjectId(),
		"admin",
		"admin",
		"md5",
		time.Now().Unix(),
		"secretkey",
	}
	db:=GetUserDB()
	t.Log(db.Insert(u))
}

func TestUserDB_FindUserByUserName(t *testing.T) {
	db:=GetUserDB()
	user,err:=db.FindUserByUserName("test_user")
	t.Log(err)
	t.Log(user)

	s:=user.Username+user.Password+utils.Int64ToString(user.CreatetTime)
	ps:=utils.GetSHA256(s)
	t.Log(ps)
}

func TestUserDB_FindUserById(t *testing.T) {
	db:=GetUserDB()
	user,err:=db.FindUserById("5b9384187cb15f0f61ec1e69")
	t.Log(err)
	t.Log(user)
}
