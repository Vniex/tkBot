package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"errors"
	"log"
)

type Asset struct {
	session *mgo.Session
	collection *mgo.Collection

	asset []AssetItem
}

type AssetItem struct {
	Name			string `json:"name"`
	Date            int64   `json:"date"`
	Hm              string  `json:"hm"`
	High            float64 `json:"high"`
	Low             float64 `json:"low"`
	Open            float64 `json:"open"`
	Close           float64 `json:"close"`
	Volume          float64 `json:"volume"`
	QuoteVolume     float64 `json:"quoteVolume"`
	WeightedAverage float64 `json:"weightedAverage"`
	Exchange		string `json:"exchange"`
}

func (t *Asset) Connect() error {
	session, err := mgo.Dial(MongoURL)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return err
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(Database).C(AssetCollectin)

	t.session = session
	t.collection = c

	return nil
}

func (t *Asset) Close() {
	if t.session != nil {
		t.session.Close()
	}
}

func (t *Asset) LoadCharts(exchange string, name string, period int) error {

	if t.collection == nil {
		return errors.New("Mongo is not connected")
	}

	t.asset = []AssetItem{}
	t.collection.Find(bson.M{"exchange": exchange, "name": name}).All(&t.asset);

	for i := 0; i < len(t.asset); i++{
		log.Printf("chart:%v",t.asset[i])
	}

	return nil;
}
