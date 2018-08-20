package annSpider
import (

	"time"
	"log"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"errors"
	"bytes"
	"github.com/fatih/structs"


	Utils "tkBot/utils"
	DB "tkBot/database/mongo"
	Robot "tkBot/strategy"
	WebSocket "tkBot/server/websocket"
)


const Ann_URL="http://localhost:8888/api/v1/announcement/"
const Strategy_URL="http://localhost:8888/api/v1/strategy/"



type Parameters struct{
	RobotName string `json:"robot_name"`
	Interval int `json:"interval" `
	Retry int `json:"retry" `

}


func Register() error{
	p:=&Parameters{
		"annSpider-0",
		5*60,
		5,

	}
	structs.DefaultTagName="json"
	para_map:=structs.Map(p)
	st:=&DB.Strategy{
		1,
		"annSpider",
		"announcement spider",
		para_map,
	}
	data, _ := json.Marshal(st)

	return posData(Strategy_URL,data)


}


type Strategy struct {
	*Robot.Robot
	AnnChan chan *DB.Announcement
	LastAnnId int
	AnnUrl string
	StrategyUrl string
	Para *Parameters


}


func NewStrategy(parameters *Parameters) *Strategy{

	return &Strategy{
		Robot.NewRobot(),
		make(chan *DB.Announcement,300),
		0,
		Ann_URL,
		Strategy_URL,
		parameters,
	}

}



func (s *Strategy)Run(){
	live_msg:=WebSocket.NewRobotMsg(s.Para.RobotName,WebSocket.CmdType_HEARTBEAT,"heartbeat from "+s.Para.RobotName)
	go s.LiveDetect(live_msg)
	for {
		go s.Parse()
		time.Sleep(time.Duration(s.Para.Interval)*time.Second)
	}
}



func posData(url string,data []byte)error{
	resp,err:=http.Post(url,"application/x-www-form-urlencoded",bytes.NewBuffer(data))
	if err!=nil{
		log.Println(err)
		return err
	}
	body, err:= ioutil.ReadAll(resp.Body)
	respmap := make(map[string]interface{}, 1)

	err = json.Unmarshal(body, &respmap)
	if err != nil {
		log.Println(err)
		return err
	}
	if !respmap["success"].(bool){
		log.Println(respmap["message"].(string))
		return errors.New(respmap["message"].(string))
	}
	return nil
}


func(s *Strategy) Parse()error{
	recover()
	var (
		resp []byte
		err error
		)
	resp=Utils.RE(s.Para.Retry,s.Fetch).([]byte)
	respmap := make(map[string]interface{}, 1)
	err = json.Unmarshal(resp, &respmap)
	if err != nil {
		log.Println(err)
		return err
	}
	if Utils.ToInt(respmap["code"]) != 0 {
		log.Println(respmap["message"].(string))
		return errors.New(respmap["message"].(string))
	}
	datamap := respmap["data"].(map[string]interface{})
	ann_list := datamap["list"].([]interface{})
	max_ann_id:=s.LastAnnId
	for _, item := range ann_list {
		ann := item.(map[string]interface{})
		ann_id:=Utils.ToInt(ann["id"])
		if ann_id<s.LastAnnId{
			continue
		}

		max_ann_id=Utils.Max(max_ann_id,ann_id)
		new_ann:=DB.NewAnnouncement(
			ann_id,
			ann["title"].(string),
			ann["abstract"].(string),
			ann["link"].(string),
			Utils.ToInt(ann["exchange_id"]),
			ann["logo"].(string),
			ann["exchange_name"].(string),
			ann["exchange_alias"].(string),
			Utils.ToInt(ann["posted_at"]),
			ann["lang_type"].(string),
			ann["source"].(string),
		)
		//log.Println(ann)
		a, _ := json.Marshal(new_ann)
		go Utils.RE(s.Para.Retry,posData,s.AnnUrl,a)
		//s.AnnChan<-new_ann
	}
	return nil

}

func (s *Strategy)Fetch()([]byte,error){
	var (
		url string
		err error
		resp *http.Response
		body []byte
	)

	url = "https://mytoken.io/api/media/medialist?type=6&size=1&timestamp=1534161443394&code=b0ba0cd1f9065fbb8ebf8a3de493ad45&platform=web_pc"
	resp,err=http.Get(url)
	if err!=nil{
		log.Println(err)
		return nil,err
	}
	body, err= ioutil.ReadAll(resp.Body)
	if err!=nil{
		log.Println(err)
		return nil,err
	}
	return body,nil

}
