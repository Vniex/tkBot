package service

import (
	"tkBot/server/service/pb"
	"google.golang.org/grpc"
	"tkBot/server/service/announcement"
)

func ServiceReg(server *grpc.Server)  {

	pb.RegisterAnnouncementServiceServer(server,&announcement.Announcement{})

}