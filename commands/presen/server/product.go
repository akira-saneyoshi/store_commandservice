package server

import (
	"commandservice/application/service"
	"commandservice/presen/adapter"
	"context"

	v1 "github.com/akira-saneyoshi/store_pb/pb/v1"
)

// 商品更新サーバの実装
type productServer struct {
	adapter adapter.ProductAdapter // 商品変換
	service service.ProductService // 商品更新サービス
	// 生成されたUnimplementedProductCommandServerをエンベデッド
	v1.UnimplementedProductCommandServer
}

// コンストラクタ
func NewprductServer(adapter adapter.ProductAdapter, service service.ProductService) v1.ProductCommandServer {
	return &productServer{adapter: adapter, service: service}
}

// 商品の追加 v1.ProductCommandServerインターフェイスのメソッド実装
func (ins *productServer) Create(ctx context.Context, param *v1.ProductUpParam) (*v1.ProductUpResult, error) {
	product, err := ins.adapter.ToEntity(param)
	if err != nil {
		return ins.adapter.ToResult(err), nil
	}
	if err := ins.service.Add(ctx, product); err != nil {
		return ins.adapter.ToResult(err), nil
	}
	return ins.adapter.ToResult(product), nil
}

// 商品の変更 v1.ProductCommandServerインターフェイスのメソッド実装
func (ins *productServer) Update(ctx context.Context, param *v1.ProductUpParam) (*v1.ProductUpResult, error) {
	product, err := ins.adapter.ToEntity(param)
	if err != nil {
		return ins.adapter.ToResult(err), nil
	}
	if err := ins.service.Update(ctx, product); err != nil {
		return ins.adapter.ToResult(err), nil
	}
	return ins.adapter.ToResult(product), nil
}

// 商品の削除 v1.ProductCommandServerインターフェイスのメソッド実装
func (ins *productServer) Delete(ctx context.Context, param *v1.ProductUpParam) (*v1.ProductUpResult, error) {
	product, err := ins.adapter.ToEntity(param)
	if err != nil {
		return ins.adapter.ToResult(err), nil
	}
	if err := ins.service.Delete(ctx, product); err != nil {
		return ins.adapter.ToResult(err), nil
	}
	return ins.adapter.ToResult(product), nil
}
