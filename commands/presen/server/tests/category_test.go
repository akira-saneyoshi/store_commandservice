package server_test

import (
	"commandservice/application"
	"commandservice/presen/adapter"
	"commandservice/presen/server"
	"context"
	"fmt"

	v1 "github.com/akira-saneyoshi/store_pb/pb/v1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/fx"
)

var _ = Describe("categoryServer構造体", Ordered, Label("メソッドのテスト"), func() {
	var srv v1.CategoryCommandServer
	var category *v1.Category
	var ctx context.Context
	var container *fx.App
	// 前処理
	BeforeAll(func() {
		ctx = context.Background() // Contextの生成
		container = fx.New(
			application.SrvDepend,
			fx.Provide(
				adapter.NewcategoryAdapaterImpl,
				server.NewcategoryServer,
			),
			fx.Populate(&srv),
		)
		// fxを起動し、起動時にエラーがないことを確認する
		err := container.Start(ctx)
		Expect(err).NotTo(HaveOccurred())
	})
	// 後処理
	AfterEach(func() {
		err := container.Stop(context.Background())
		Expect(err).NotTo(HaveOccurred())
	})
	// Add()メソッドのテスト
	Context("Add()メソッドのテスト", Label("Add"), func() {
		It("カテゴリ登録が成功し、CategoryUpResultが返る", func() {
			param := v1.CategoryUpParam{Crud: v1.CRUD_INSERT, Id: "", Name: "飲料水"}
			result, _ := srv.Create(ctx, &param)
			category = result.Category
			Expect(result.Error).To(BeNil())
		})
		It("カテゴリ登録が失敗し、v1.Errorを保持したCategoryUpResultが返る", func() {
			param := v1.CategoryUpParam{Crud: v1.CRUD_INSERT, Id: category.GetId(), Name: category.GetName()}
			result, _ := srv.Create(ctx, &param)
			e := v1.Error{Type: "CRUD Error", Message: "飲料水は既に登録されています。"}
			Expect(result.Error).To(Equal(&e))
		})
	})
	// Update()メソッドのテスト
	Context("Update()メソッドのテスト", Label("Update"), func() {
		It("カテゴリの更新が成功し、CategoryUpResultが返る", func() {
			param := v1.CategoryUpParam{Crud: v1.CRUD_UPDATE, Id: category.GetId(), Name: "衣料品"}
			result, _ := srv.Update(ctx, &param)
			Expect(result.Error).To(BeNil())
		})
		It("カテゴリの更新が失敗し、CategoryUpResultが返る", func() {
			id := "b1524011-b6af-417e-8bf2-f449dd58b5c1"
			param := v1.CategoryUpParam{Crud: v1.CRUD_UPDATE, Id: id, Name: "衣料品"}
			result, _ := srv.Update(ctx, &param)
			e := v1.Error{Type: "CRUD Error", Message: fmt.Sprintf("カテゴリ番号:%sは存在しないため、更新できませんでした。", id)}
			Expect(result.Error).To(Equal(&e))
		})
	})
	// Delete()メソッドのテスト
	Context("Delete()メソッドのテスト", Label("Delete"), func() {
		It("カテゴリの削除が成功し、CategoryUpResultが返る", func() {
			param := v1.CategoryUpParam{Crud: v1.CRUD_DELETE, Id: category.GetId(), Name: category.GetName()}
			result, _ := srv.Delete(ctx, &param)
			Expect(result.Error).To(BeNil())
		})
		It("カテゴリの削除が失敗し、CategoryUpResultが返る", func() {
			id := "b1524011-b6af-417e-8bf2-f449dd58b5c1"
			param := v1.CategoryUpParam{Crud: v1.CRUD_DELETE, Id: id, Name: "衣料品"}
			result, _ := srv.Delete(ctx, &param)
			e := v1.Error{Type: "CRUD Error", Message: fmt.Sprintf("カテゴリ番号:%sは存在しないため、削除できませんでした。", id)}
			Expect(result.Error).To(Equal(&e))
		})
	})
})
