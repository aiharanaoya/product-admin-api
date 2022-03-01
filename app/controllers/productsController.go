package controllers

import (
	"fmt"
	"net/http"
)

// ハンドラ関数
// URL、HTTPメソッドから呼び出す関数をハンドリングする。
// 基本的にコントローラ関数を呼び出すのみで処理はコントローラ関数に記載する。

// 商品ハンドラ
func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		xxxx(w, r)
	case http.MethodPost:
		xxxx(w, r)
	default:
		ResponseError(w, http.StatusNotFound)
	}
}

// 商品IDハンドラ
func productsIdHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		xxxx(w, r)
	case http.MethodPut:
		xxxx(w, r)
	case http.MethodDelete:
		xxxx(w, r)
	default:
		ResponseError(w, http.StatusNotFound)
	}
}

func xxxx(w http.ResponseWriter, r *http.Request) {
	fmt.Println("xxxx")
}