package mongo

import (
	"time"
	"testing"
)

func TestAssetDB_Insert(t *testing.T) {

	asset:=&Asset{
		"test",
		100,
		int(time.Now().Unix()),
	}
	db:=GetAssetDB()
	t.Log(db.Insert(asset))
}

func TestAssetDB_FindAssetsInTime(t *testing.T) {
	db:=GetAssetDB()
	assets,err:=db.FindAssetsInTimeByRobot("test",int(time.Now().Unix()-2*60*60))
	t.Log(err)
	t.Log(assets[0].RobotName)
	t.Log(assets[0].NetAsset)
	t.Log(assets[0].TimeStamp)
}