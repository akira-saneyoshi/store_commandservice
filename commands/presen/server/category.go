package server

import (
	"commandservice/application/service"
	"commandservice/presen/adapter"
	"context"

	v1 "github.com/akira-saneyoshi/store_pb/pb/v1"
)

// カテゴリ更新サーバの実装
type categoryServer struct {
	adapter adapter.CategoryAdapter // カテゴリ変換
	service service.CategoryService // カテゴリ更新サービス
	// 生成されたUnimplementedCategoryCommandServerをエンベデッド
	v1.UnimplementedCategoryCommandServer
}

// コンストラクタ
func NewcategoryServer(adapter adapter.CategoryAdapter, service service.CategoryService) v1.CategoryCommandServer {
	return &categoryServer{adapter: adapter, service: service}
}

// カテゴリの追加 v1.CategoryCommandServerインターフェイスのメソッド実装
func (ins *categoryServer) Create(ctx context.Context, param *v1.CategoryUpParam) (*v1.CategoryUpResult, error) {
	// v1.CategoryUpParamをentity.Categoryに変換する
	if category, err := ins.adapter.ToEntity(param); err != nil {
		return ins.adapter.ToResult(err), nil // CategoryUpResultにエラーを設定
	} else {
		// サービスのAdd()メソッドを実行する
		if err := ins.service.Add(ctx, category); err != nil {
			return ins.adapter.ToResult(err), nil // CategoryUpResultにエラーを設定
		}
		return ins.adapter.ToResult(category), nil // CategoryUpResultにCategoryを設定
	}
}

// カテゴリの変更 v1.CategoryCommandServerインターフェイスのメソッド実装
func (ins *categoryServer) Update(ctx context.Context, param *v1.CategoryUpParam) (*v1.CategoryUpResult, error) {
	// v1.CategoryUpParamをentity.Categoryに変換する
	if category, err := ins.adapter.ToEntity(param); err != nil {
		return ins.adapter.ToResult(err), nil // CategoryUpResultにエラーを設定
	} else {
		// サービスのUpdate()メソッドを実行する
		if err := ins.service.Update(ctx, category); err != nil {
			return ins.adapter.ToResult(err), nil // CategoryUpResultにエラーを設定
		}
		return ins.adapter.ToResult(category), nil // CategoryUpResultにCategoryを設定
	}
}

// カテゴリの削除 v1.CategoryCommandServerインターフェイスのメソッド実装
func (ins *categoryServer) Delete(ctx context.Context, param *v1.CategoryUpParam) (*v1.CategoryUpResult, error) {
	// v1.CategoryUpParamをentity.Categoryに変換する
	if category, err := ins.adapter.ToEntity(param); err != nil {
		return ins.adapter.ToResult(err), nil // CategoryUpResultにエラーを設定
	} else {
		// サービスのDelete()メソッドを実行する
		if err := ins.service.Delete(ctx, category); err != nil {
			return ins.adapter.ToResult(err), nil // CategoryUpResultにエラーを設定
		}
		// CategoryUpResultにCategoryを設定して返す
		return ins.adapter.ToResult(category), nil
	}
}
