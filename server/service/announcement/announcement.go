package announcement

import (
	"context"
	"tkBot/server/service/pb"
	Mongo "tkBot/database/mongo"
	"tkBot/server/service"
	"log"
)

type Announcement struct {
	Id int32 `bson:"id" json:"id"`
	Title string `bson:"title" json:"title"`
	Abstract  string `bson:"abstract" json:"abstract"`
	Link string `bson:"link" json:"link"`
	ExchangeId int32 `bson:"exchange_id" json:"exchange_id"`
	Logo string `bson:"logo" json:"logo"`
	ExchangeName string `bson:"exchange_name" json:"exchange_name"`
	ExchangeAlias string `bson:"exchange_alias" json:"exchange_alias"`
	PostedAt int64 `bson:"posted_at" json:"posted_at"`
	LangType string `bson:"lang_type" json:"lang_type"`
	Source string `bson:"source" json:"source"`

}


func Req2Announcement(req *pb.CreateAnnouncementRequest) *Announcement{
	return &Announcement{
		req.Id,
		req.Title,
		req.Abstract,
		req.Link,
		req.ExchangeId,
		req.Logo,
		req.ExchangeName,
		req.ExchangeAlias,
		req.PostedAt,
		req.LangType,
		req.Source,
	}

}


func (this *Announcement)GetAnnouncementList(ctx context.Context, req *pb.GetAnnouncementListRequest) (*pb.GetAnnouncementListResponse, error){
	db:=Mongo.GetAnnouncementDB()
	anns,err:=db.FindAnnsBeforeTime(req.ExchangeId,req.LastTime,int(req.Limit))
	if err!=nil{
		log.Println(err)
		resp:=&pb.GetAnnouncementListResponse{
			service.ERROR_INTERNAL,
			err.Error(),
			nil,
			nil,nil,nil,
		}
		return resp,err
	}
	data:=make([]*pb.AnnouncementResponse,0,20)
	for _,ann:= range anns{
		data=append(data, &pb.AnnouncementResponse{
			ann.Id,
			ann.Title,
			ann.Link,
			ann.Logo,
			ann.PostedAt,
			nil,
			nil,
			nil,

		})
	}
	resp:=&pb.GetAnnouncementListResponse{
		service.SUCCESS,
		"",
		data,
		nil,
		nil,
		nil,
	}

	return resp,nil
}

func (this *Announcement) CreateAnnouncement(ctx context.Context,req  *pb.CreateAnnouncementRequest) (*pb.CreateAnnouncementResponse, error){
	ann:=Req2Announcement(req)
	db:=Mongo.GetAnnouncementDB()
	if err:=db.Insert(ann);err!=nil{
		log.Println(err)
		resp:=&pb.CreateAnnouncementResponse{
			service.ERROR_INTERNAL,
			err.Error(),
			nil,
			nil,
			nil,
		}
		return resp,err
	}

	resp:=&pb.CreateAnnouncementResponse{
		service.SUCCESS,
		"",
		nil,
		nil,
		nil,
	}
	return resp,nil
}