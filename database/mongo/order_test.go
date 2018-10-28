package mongo

import "testing"

func TestOrderDB_Insert(t *testing.T) {
	order:=&Order{
		"test",
		6.8,
		100,
		0,
		0,
		"123",
		0,
		"USDT_QC",
		"BUY",

	}
	db:=GetOrderDB()
	t.Log(db.Insert(order))
}

func TestOrderDB_FindAllOrderByRobot(t *testing.T) {
	db:=GetOrderDB()
	res,err:=db.FindAllOrderByRobot("test")
	t.Log(err)
	t.Log(res)
	t.Log(res[0])
}