package adapter

import (
	"commandservice/domain/models/products"

	v1 "github.com/akira-saneyoshi/store_pb/pb/v1"
)

// パラメータと実行結果の変換インターフェス
type ProductAdapter interface {
	// ProductUpParamからProductに変換する
	ToEntity(param *v1.ProductUpParam) (*products.Product, error)
	// サービス実行結果からProductUpResultに変換する
	ToResult(result any) *v1.ProductUpResult
}
