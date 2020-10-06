// @Author: 2014BDuck
// @Date: 2020/10/6

package server

import (
	"context"
	"encoding/json"
	"github.com/2014BDuck/tag-service/pkg/bapi"
	"github.com/2014BDuck/tag-service/pkg/errcode"
	pb "github.com/2014BDuck/tag-service/proto"
)

type TagServer struct{}

func NewTagServer() *TagServer {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	api := bapi.NewAPI("http://127.0.0.1:8000")
	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, errcode.TogRPCError(errcode.ErrorGetTagListFail)
	}
	tagList := pb.GetTagListReply{}
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	return &tagList, nil
}
