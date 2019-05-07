package db

import (
	pb "github.com/wh-servers/grpcContacterServer/contacter"
)

func OrmInit(conf string) (err error) {
	return
}

func Insert(con *pb.Contacter) (res *pb.Contacter, err error) {
	return
}

func GetInfo(con *pb.Contacter) (res *pb.Contacter, err error) {
	//get from map

	//get from db

	return
}

func Modify(con *pb.Contacter) (res *pb.Contacter, err error) {
	return
}

func Delete(id int32) (res *pb.Contacter, err error) {
	return
}

func ReloadMap(loadType pb.LoadType, typeId interface{}) (err error) {
	switch loadType {
	case pb.LoadType_ID:
		reloadMap("id", typeId)
	case pb.LoadType_NAME:
		reloadMap("name", typeId)
	case pb.LoadType_COUNTRY:
		reloadMap("country", typeId)
	case pb.LoadType_EMAIL:
		reloadMap("email", typeId)
	case pb.LoadType_STATUS:
		reloadMap("status", typeId)
	case pb.LoadType_HEADPHONE:
		reloadMap("headphone", typeId)
	}
	return
}
