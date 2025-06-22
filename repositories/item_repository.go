package repositories

import (
	"errors"
	"gin_practice/models"
)

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
}

type ItemMemoryRepository struct {
	items []models.Item
}

// 普通の関数（notメソッド）
// リポジトリ作成時には以下のようにして使う
// repo := NewItemMemoryRepository(items)
func NewItemMemoryRepository(items []models.Item) IItemRepository {
	// 値コピーしてそのアドレスを返す
	// メソッドがポインタレシーバで定義されているため、値型で返すとコンパイルエラーになる
	// そのため & が必要
	return &ItemMemoryRepository{items: items}

	// instance := ItemMemoryRepository{items: items}
	// return &instance
	// これを1行でまとめて書くと return &ItemMemoryRepository{items: items}
}

// メソッド（特定の型に属する関数）
// ItemMemoryRepository型のすべてのインスタンスにFindAllメソッドを紐づけている
// *なしの値レシーバーにするとコピーを作成してしまってNG
func (r *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}

func (r *ItemMemoryRepository) FindById(itemId uint) (*models.Item, error) {
	for _, v := range r.items {
		if v.ID == itemId {
			return &v, nil
		}
	}

	return nil, errors.New("item not found")
}
