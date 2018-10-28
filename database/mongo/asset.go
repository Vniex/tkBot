package mongo

import (
	"gopkg.in/mgo.v2/bson"

	log "github.com/sirupsen/logrus"
)


var assetDB *AssetDB

type Asset struct {
	RobotName string `bson:"robot_name" json:"robot_name"`
	NetAsset float64 `bson:"net_asset" json:"net_asset"`
	TimeStamp int `bson:"time_stamp" json:"time_stamp"`

}


func NewAsset(robot_name string ,net_asset float64,timestamp int)*Asset{

	return &Asset{
		robot_name,
		net_asset,
		timestamp,
	}
}

type AssetDB struct {
	Database string
	Collection string

}

func GetAssetDB() *AssetDB{
	if assetDB==nil{
		assetDB=&AssetDB{
			Database,
			AssetCollection,
		}
	}
	return assetDB
}

func (db *AssetDB)Insert(asset *Asset) error{
	session_clone,err:=GetSessionClone()
	if err != nil {
		log.Error(err)
		return err
	}
	defer session_clone.Clone()
	err=session_clone.DB(db.Database).C(db.Collection).Insert(asset)
	if err!=nil{
		log.Error(err)
		return err
	}
	return nil
}


func (db *AssetDB)FindAssetsInTimeByRobot(robot_name string,timestamp int) ([]*Asset,error) {
	var result []*Asset
	session_clone,err:=GetSessionClone()
	if err != nil {

		log.Error(err)
		return nil,err
	}
	defer session_clone.Clone()
	err = session_clone.DB(db.Database).C(db.Collection).Find(bson.M{"robot_name":robot_name,"time_stamp": bson.M{"$gte": timestamp}}).Sort("time_stamp").All(&result)
	if err != nil {
		log.Error(err)
		return nil,err
	}

	return result,nil
}