package prepare

import (
	v1 "github.com/akira-saneyoshi/store_pb/pb/v1"
	"google.golang.org/grpc"
)

// gRPCサーバの生成とServer機能の登録
type CommandServer struct {
	Server *grpc.Server // gRPCServer
}

// コンストラクタ 平文を利用する
func NewCommandServer(category v1.CategoryCommandServer, product v1.ProductCommandServer) *CommandServer {
	// gRPCサーバを生成する
	server := grpc.NewServer()
	// CategoryCommandServerを登録する
	v1.RegisterCategoryCommandServer(server, category)
	// ProductCommandServerを登録する
	v1.RegisterProductCommandServer(server, product)
	return &CommandServer{Server: server}
}
