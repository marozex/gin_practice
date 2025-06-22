package repositories

import "gin_practice/models"

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
}

type ItemMemoryRepository struct {
	items []models.Item
}

// リポジトリ作成時には以下のようにして使う
// repo := NewItemMemoryRepository(items)
func NewItemMemoryRepository(items []models.Item) IItemRepository {
	return &ItemMemoryRepository{items: items}
}

// ItemMemoryRepository型のすべてのインスタンスにFindAllメソッドを紐づけている
func (r *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}
