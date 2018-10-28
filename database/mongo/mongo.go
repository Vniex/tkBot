package mongo

import (
	"gopkg.in/mgo.v2"
	"log"
	Config "tkBot/config"
)

const MongoURL = "mongodb://"+Config.MongoUser+":"+Config.MongoPwd+"@"+Config.MongoIP+":"+Config.MongoPort



const Database = "tkBot"

const AssetCollection = "Asset"
const ExchangeCollection = "Exchange"
const AnnouncementCollection = "Announcement"
const StrategyCollection="Strategy"
const UserCollection="User"
const LogInfoCollection="LogInfo"
const OrderCollection="Order"


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

