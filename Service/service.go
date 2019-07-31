package service

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/proto"
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

func (s *ContacterServer) ModifyContacter(ctx context.Context, req *pb.ModifyRequest) (res *pb.ModifyResponse, err error) {
	contacter := req.GetContacter()
	result, _ := db.Modify(contacter)
	res.Contacter = result
	return
}

func (s *ContacterServer) Delete(ctx context.Context, req *pb.DeleteRequest) (res *pb.DeleteResponse, err error) {
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

func (s *ContacterServer) DemoAPI(ctx context.Context, r *pb.DemoRequest) (*pb.DemoResponse, error) {
	res := &pb.DemoResponse{}
	if r == nil {
		return nil, fmt.Errorf("nil req")
	}
	fmt.Println("the input is: ", r.DemoInput)
	res = &pb.DemoResponse{}
	res.DemoOutput = proto.String("this is a demo outout")
	return res, nil
}
