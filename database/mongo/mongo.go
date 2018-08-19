package mongo

import (
	"gopkg.in/mgo.v2"
	"log"
)

const MongoURL = "mongodb://192.168.8.104:27017"
//mongodb://myuser:mypass@localhost:40001,


const Database = "tkBot"

const AssetCollection = "Asset"
const ExchangeCollection = "Exchange"
const AnnouncementCollection = "Announcement"

const TradeRecordCollection = "TradeRecord"
const BalancesCollection = "Balances"
const FundCollection = "Funds"
const OkexDiffHistory = "OkexDiffHistory"
const ErrorNotConnected = "Mongo is not connected"


var GlobalSession *mgo.Session


func GetSessionClone()(*mgo.Session,error){
	var (
		err error
	)
	if GlobalSession==nil{
		GlobalSession, err = mgo.Dial(MongoURL)
		if err!=nil{
			log.Println(err)
			return nil,err
		}
	}
	return GlobalSession.Clone(),nil

}

