package controllers

import (
	"fmt"
	"net/http"
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
		// TODO aiharanaoya
		// 仮で500のStatusTextを返している。今後エラーハンドリングを実装。
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// ショップIDハンドラ
func shopsIdHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getShopById()
	case http.MethodPut:
		putShopById()
	case http.MethodDelete:
		deleteShopById()
	default:
		// TODO aiharanaoya
		// 仮で500のStatusTextを返している。今後エラーハンドリングを実装。
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
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
func getShopById() {
	fmt.Println("ショップ取得処理")
}

// ショップ更新
func putShopById() {
	fmt.Println("ショップ更新処理")
}

// ショップ削除
func deleteShopById() {
	fmt.Println("ショップ削除処理")
}