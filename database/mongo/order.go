package mongo

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

var orderDB *OrderDB


type Order struct {
	RobotName string `bson:"robot_name" json:"robot_name"`
	Price float64 `bson:"price"json:"price"`
	Amount float64 `bson:"amount" json:"amount"`
	AvgPrice float64 `bson:"avg_price" json:"avg_price"`
	Fee float64 `bson:"fee" json:"fee"`
	OrderID  string `bson:"order_id" json:"order_id"`
	OrderTime int `bson:"order_time" json:"order_time"`
	Pair  string	`bson:"pair" json:"pair"`
	Side  string `bson:"side" json:"side"`

}

func NewOrderRough(robot_name string, price,amount,fee float64,orderId,side,pair string,timestamp int) *Order{
	return &Order{
		robot_name,
		price,
		amount,
		price,
		fee,
		orderId,
		timestamp,
		pair,
		side,

	}
}



type OrderDB struct {
	Database string
	Collection string

}

func GetOrderDB() *OrderDB{
	if orderDB==nil{
		orderDB=&OrderDB{
			Database,
			OrderCollection,
		}
	}
	return orderDB
}

func (db *OrderDB)Insert(order *Order) error{
	session_clone,err:=GetSessionClone()
	if err != nil {
		log.Error(err)
		return err
	}
	defer session_clone.Clone()
	err=session_clone.DB(db.Database).C(db.Collection).Insert(order)
	if err!=nil{
		log.Error(err)
		return err
	}
	return nil
}


func (db *OrderDB)FindAllOrderByRobot(robot_name string) ([]*Order,error) {
	var result []*Order
	session_clone,err:=GetSessionClone()
	if err != nil {
		log.Error(err)
		return nil,err
	}
	defer session_clone.Clone()
	err = session_clone.DB(db.Database).C(db.Collection).Find(bson.M{"robot_name":robot_name}).All(&result)
	if err != nil {
		log.Error(err)
		return nil,err
	}

	return result,nil
}

