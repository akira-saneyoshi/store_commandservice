package adapter

import (
	"commandservice/domain/models/categories"
	"commandservice/domain/models/products"
	"commandservice/errs"

	v1 "github.com/akira-saneyoshi/store_pb/pb/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// パラメータと実行結果の変換インターフェイスの実装
type productAdapaterImpl struct{}

func NewproductAdapaterImpl() ProductAdapter {
	return &productAdapaterImpl{}
}

// ProductUpParamからProductに変換する
func (ins *productAdapaterImpl) ToEntity(param *v1.ProductUpParam) (*products.Product, error) {
	// コマンド種別別のエンティティ生成
	switch param.GetCrud() {
	case v1.CRUD_INSERT:
		name, err := products.NewProductName(param.GetName())
		if err != nil {
			return nil, err
		}
		price, err := products.NewProductPrice(uint32(param.GetPrice()))
		if err != nil {
			return nil, err
		}
		id, err := categories.NewCategoryId(param.GetCategoryId())
		if err != nil {
			return nil, err
		}
		product, err := products.NewProduct(name, price, categories.BuildCategory(id, nil))
		if err != nil {
			return nil, err
		}
		return product, nil
	case v1.CRUD_UPDATE:
		id, err := products.NewProductId(param.GetId())
		if err != nil {
			return nil, err
		}
		name, err := products.NewProductName(param.GetName())
		if err != nil {
			return nil, err
		}
		price, err := products.NewProductPrice(uint32(param.GetPrice()))
		if err != nil {
			return nil, err
		}
		cid, err := categories.NewCategoryId(param.GetCategoryId())
		if err != nil {
			return nil, err
		}
		return products.BuildProduct(id, name, price, categories.BuildCategory(cid, nil)), nil
	case v1.CRUD_DELETE:
		id, err := products.NewProductId(param.GetId())
		if err != nil {
			return nil, err
		}
		return products.BuildProduct(id, nil, nil, nil), nil
	default:
		return nil, errs.NewDomainError("不明な操作を受信しました。")
	}
}

// サービス実行結果からProductUpResultに変換する
func (ins *productAdapaterImpl) ToResult(result any) *v1.ProductUpResult {
	var up_product *v1.Product
	var up_err *v1.Error
	switch v := result.(type) {
	case *products.Product: // 実行結果がentity.Productの場合
		var c *v1.Category
		if v.Category() == nil {
			c = &v1.Category{Id: "", Name: ""}
		} else {
			c = &v1.Category{Id: v.Category().Id().Value(), Name: ""}
		}
		var name string = ""
		if v.Name() != nil {
			name = v.Name().Value()
		}
		var price int32 = 0
		if v.Price() != nil {
			price = int32(v.Price().Value())
		}
		up_product = &v1.Product{Id: v.Id().Value(), Name: name, Price: price, Category: c}
	case *errs.DomainError: // 実行結果がerrs.DomainErrorの場合
		up_err = &v1.Error{Type: "Domain Error", Message: v.Error()}
	case *errs.CRUDError: // 実行結果がerrs.CRUDErrorの場合
		up_err = &v1.Error{Type: "CRUD Error", Message: v.Error()}
	case *errs.InternalError: // 実行結果がerrs.InternalErrorの場合
		up_err = &v1.Error{Type: "Internal Error", Message: "只今、サービスを提供できません。"}
	}
	return &v1.ProductUpResult{
		Product:   up_product,
		Error:     up_err,
		Timestamp: timestamppb.Now(),
	}
}
