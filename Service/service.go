package service

import (
	"context"

	proto "github.com/golang/protobuf/proto"
	pb "github.com/wh-servers/grpcContacterServer/contacter"
	db "github.com/wh-servers/grpcContacterServer/db/service"
	redis "github.com/wh-servers/grpcContacterServer/redis/service"
)

type ContacterServer struct{}

func (s *ContacterServer) Insert(ctx context.Context, req *pb.InsertRequest) (res *pb.InsertResponse, err error) {
	result, _ := db.Insert(req.GetContacter())
	res.Contacter = result
	return
}

func (s *ContacterServer) GetInfo(ctx context.Context, req *pb.GetInfoRequest) (res *pb.GetInfoResponse, err error) {
	contacter := req.GetContacter()
	result := &pb.Contacter{}
	if req.GetSource() == pb.SourceType_REDIS {
		result, _ = redis.GetInfo(contacter)
	} else {
		result, _ = db.GetInfo(contacter) //including get from map and get from db directly.
	}
	res.Contacter = result
	return
}

func (s *ContacterServer) Modify(ctx context.Context, req *pb.ModifyRequest) (res *pb.ModifyResponse, err error) {
	contacter := req.GetContacter()
	result, _ := db.Modify(contacter)
	res.Contacter = result
	return
}

func (s *ContacterServer) Delete(cctx context.Context, req *pb.DeleteRequest) (res *pb.DeleteResponse, err error) {
	id := req.GetId()
	result, err := db.Delete(id)
	res.Contacter = result
	res.ErrMsg = proto.String("Delete success")
	return
}

func (s *ContacterServer) ReloadMap(ctx context.Context, req *pb.ReloadMapRequest) (res *pb.ReloadMapResponse, err error) {
	loadType := req.GetLdType()
	typeId := req.GetTypeId()
	db.ReloadMap(loadType, typeId)
	return
}
