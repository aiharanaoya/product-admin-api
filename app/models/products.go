package models

import (
	"fmt"
	"time"
)

// モデル
// 構造体定義、DBにアクセスする関数、レシーバメソッドを記載する。

// 商品構造体
type Product struct {
	Id	string	`json:"id"`
	Title	string	`json:"title"`
	Price	int	`json:"price"`
	Description	string	`json:"description"`
	ShopId	string	`json:"shopId"`
	CreatedAt	time.Time	`json:"createdAt"`
	UpdatedAt	time.Time	`json:"updatedAt"`
}

// 商品一覧取得レスポンス構造体
type ProductListRes struct {
	Products	[]Product	`json:"products"`
	Pagination	Pagination `json:"pagination"`
}

// 商品の検索を行う
func SearchProducts(page int, perPage int, title string) (products []Product, err error) {
	// titleは部分一致
	cmd := `
		select id, title, price, description, shopId, created_at, updated_at
		from products
		where title like concat('%', ?, '%')
		limit ?
		offset ?
	`

	offset := (page * perPage) - perPage

	rows, err := Db.Query(cmd, title, perPage, offset)
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var product Product
		err = rows.Scan(
			&product.Id,
			&product.Title,
			&product.Price,
			&product.Description,
			&product.ShopId,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			fmt.Println(err)
		}
		products = append(products, product)
	}
	rows.Close()

	return products, err
}

// 商品のトータル件数を取得する
func FetchProductsTotal() (total int, err error) {
	cmd := `
		select count(*)
		from products
	`

	err = Db.QueryRow(cmd).Scan(&total)

	return total, err
}

// IDから商品1件を取得する
func FetchProductById(id string) (product Product, err error) {
	cmd := `
		select id, title, price, description, shopId, created_at, updated_at
		from products
		where id = ?
	`

	err = Db.QueryRow(cmd, id).Scan(
		&product.Id,
			&product.Title,
			&product.Price,
			&product.Description,
			&product.ShopId,
			&product.CreatedAt,
			&product.UpdatedAt,
	)

	return product, err
}

// 商品を作成する
func (s *Product) CreateProduct() (id string, err error) {
	cmd := `
		insert into products (title, price, description, shopId)
		values (?, ?, ?, ?)
	`

	_, err = Db.Exec(cmd, s.Title, s.Price, s.Description, s.ShopId)
	if err != nil {
		fmt.Println(err)
	}

	err = Db.QueryRow(`select last_insert_id()`).Scan(
		&id,
	)

	return id, err
}

// 指定IDの商品を更新する
func (s *Product) UpdateProductById() (err error) {
	cmd := `
		update products
		set title = ?, price = ?, description = ?, shopId = ?,
		where id = ?
	`

	_, err = Db.Exec(cmd, s.Name, s.Description, s.Id)

	return err
}

// 指定IDの商品を削除する
func DeleteProductById(id string) (err error) {
	cmd := `
		delete
		from products
		where id = ? 
	`

	_, err = Db.Exec(cmd, id)

	return err
}
