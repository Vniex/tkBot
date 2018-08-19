package mongo

import "testing"

func TestExchangeDB_Insert(t *testing.T) {
	ex:=NewExchange(
		108,
		"Huobi Pro",
		"火币全球站",
		"https://cdn.mytoken.org/Fts8YtPInuwPmWxySgOlo1T55E_s",
		16187,

	)
	db:=NewExchangeDB()
	t.Log(db.Insert(ex))
}

func TestExchangeDB_FindOneByName(t *testing.T) {
	db := NewExchangeDB()
	ex, err := db.FindOneByName("Huobi Pro")
	t.Log(err)
	t.Log(ex)
}

func TestExchangeDB_FindOneById(t *testing.T) {
	db := NewExchangeDB()
	ex, err := db.FindOneById(108)
	t.Log(err)
	t.Log(ex)
}

func TestExchangeDB_DeleteById(t *testing.T) {
	db:=NewExchangeDB()
	err:=db.DeleteById(108)
	t.Log(err)
}

func TestExchangeDB_AddAnnouncement(t *testing.T) {
	db:=NewExchangeDB()
	err:=db.AddAnnouncement(108,126)
	t.Log(err)
}

