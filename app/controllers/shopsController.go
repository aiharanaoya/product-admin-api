package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/nao11aihara/product-admin-api/app/models"
)

// ハンドラ関数
// URL、HTTPメソッドから呼び出す関数をハンドリングする。
// 基本的にコントローラ関数を呼び出すのみで処理はコントローラ関数に記載する。

// ショップハンドラ
func shopsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getshops()
	case http.MethodPost:
		postShop()
	default:
		// 仮エラーハンドリング
		http.Error(w, "仮エラー", http.StatusInternalServerError)
	}
}

// ショップIDハンドラ
func shopsIdHandler(w http.ResponseWriter, r *http.Request) {
	id := getShopPathParameter(r)

	if(id == "") {
		// 仮エラーハンドリング
		http.Error(w, "仮エラー", http.StatusInternalServerError)
	}

	switch r.Method {
	case http.MethodGet:
		getShopById(w, id)
	case http.MethodPut:
		putShopById()
	case http.MethodDelete:
		deleteShopById()
	default:
		// 仮エラーハンドリング
		http.Error(w, "仮エラー", http.StatusInternalServerError)
	}
}

// コントローラ関数
// それぞれのAPIに対応した関数。
// モデル関数で定義した構造体の呼び出し、JSONの変換処理等を行う。
// DBのアクセス関数、レシーバメソッド、複雑になるロジックはモデル関数に定義する。

// ショップ一覧取得
func getshops() {
	fmt.Println("ショップ一覧取得処理")
}

// ショップ作成
func postShop() {
	fmt.Println("ショップ作成処理")
}

// ショップ取得
func getShopById(w http.ResponseWriter, id string) {
	shop, err := models.FetchShopById(id)
	if err != nil {
		fmt.Println(err)
	}

	shopRes, err := json.Marshal(shop)
	if err != nil {
			fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(shopRes))
}

// ショップ更新
func putShopById() {
	fmt.Println("ショップ更新処理")
}

// ショップ削除
func deleteShopById() {
	fmt.Println("ショップ削除処理")
}

// ロジック切り出し

// ショップのパスパラメータを取得する
func getShopPathParameter(r *http.Request) string {
	urls := strings.Split(r.RequestURI, "/")

	if(len(urls) < 4) {
		return ""
	}

	return urls[3]
}