package services

import (
	"database/sql"
	"github.com/tianhphahai2/hello-grpc"
)

type Test_rgpcServiceServer struct {
	db * sql.DB
}

func (s *Test_rgpcServiceServer) checkAPI(api string) error {
	if len(api) > 0 {
		if
	}
}


//func (HelloServiceImpl) Hello(ctx context.Context, rq *hello.HelloRequest) (*hello.HelloResponse, error)  {
//	return &hello.HelloResponse{
//		Message: "Hello " + rq.Name,
//	}, nil
//}