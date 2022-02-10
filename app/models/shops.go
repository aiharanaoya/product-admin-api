package models

import (
	"fmt"
	"time"
)

// モデル
// 構造体定義、DBにアクセスする関数、レシーバメソッドを記載する。

// ショップ構造体
type Shop struct {
	Id	string	`json:"id"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	CreatedAt	time.Time	`json:"createdAt"`
	UpdatedAt	time.Time	`json:"updatedAt"`
}

// IDからショップ1件を取得する
func FetchShopById(id string) (shop Shop, err error){
	cmd := `
		select id, name, description, created_at, updated_at
		from shops
		where id = ?
	`

	shop = Shop{}

	err = Db.QueryRow(cmd, id).Scan(
		&shop.Id,
		&shop.Name,
		&shop.Description,
		&shop.CreatedAt,
		&shop.UpdatedAt,
	)

	return shop, err
}

// ショップを作成する
func (s *Shop) CreateShop() (id string, err error) {
	cmd := `
		insert into shops (name, description)
		values (?, ?)
	`

	_, err = Db.Exec(cmd, s.Name, s.Description)
	if err != nil {
		fmt.Println(err)
	}

	err = Db.QueryRow(`select last_insert_id()`).Scan(
		&id,
	)

	return id, err
}
