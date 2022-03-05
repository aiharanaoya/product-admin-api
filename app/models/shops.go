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

// ショップ一覧取得レスポンス構造体
type ShopListRes struct {
	Shops	[]Shop	`json:"shops"`
	Pagination	Pagination `json:"pagination"`
}

// ショップの検索を行う
func SearchShops(page int, perPage int, name string) (shops []Shop, err error) {
	// nameは部分一致
	cmd := `
		select id, name, description, created_at, updated_at
		from shops
		where name like concat('%', ?, '%')
		limit ?
		offset ?
	`

	offset := (page * perPage) - perPage

	rows, err := Db.Query(cmd, name, perPage, offset)
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var shop Shop
		err = rows.Scan(
			&shop.Id,
			&shop.Name,
			&shop.Description,
			&shop.CreatedAt,
			&shop.UpdatedAt,
		)
		if err != nil {
			fmt.Println(err)
		}
		shops = append(shops, shop)
	}
	rows.Close()

	return shops, err
}

// ショップのトータル件数を取得する
func FetchShopsTotal() (total int, err error) {
	cmd := `
		select count(*)
		from shops
	`

	err = Db.QueryRow(cmd).Scan(&total)

	return total, err
}

// IDからショップ1件を取得する
func FetchShopById(id string) (shop Shop, err error) {
	cmd := `
		select id, name, description, created_at, updated_at
		from shops
		where id = ?
	`

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

// 指定IDのショップを更新する
func (s *Shop) UpdateShopById() (err error) {
	cmd := `
		update shops
		set name = ?, description = ?
		where id = ? 
	`

	_, err = Db.Exec(cmd, s.Name, s.Description, s.Id)

	return err
}

// 指定IDのショップを削除する
func DeleteShopById(id string) (err error) {
	cmd := `
		delete
		from shops
		where id = ? 
	`

	_, err = Db.Exec(cmd, id)

	return err
}
