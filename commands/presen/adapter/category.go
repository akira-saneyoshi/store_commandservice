package adapter

import (
	"commandservice/domain/models/categories"

	v1 "github.com/akira-saneyoshi/store_pb/pb/v1"
)

// パラメータと実行結果の変換インターフェス
type CategoryAdapter interface {
	// CategoryUpParamからCategoryに変換する
	ToEntity(param *v1.CategoryUpParam) (*categories.Category, error)
	// 実行結果からCategoryUpResultに変換する
	ToResult(result any) *v1.CategoryUpResult
}
