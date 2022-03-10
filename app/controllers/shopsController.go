package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
		getshops(w, r)
	case http.MethodPost:
		postShop(w, r)
	default:
		ResponseCommonError(w, http.StatusNotFound, "そのURLは存在しません。")
	}
}

// ショップIDハンドラ
func shopsIdHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getShopById(w, r)
	case http.MethodPut:
		putShopById(w, r)
	case http.MethodDelete:
		deleteShopById(w, r)
	default:
		ResponseCommonError(w, http.StatusNotFound, "そのURLは存在しません。")
	}
}

// コントローラ関数
// それぞれのAPIに対応した関数。
// モデル関数で定義した構造体の呼び出し、JSONの変換処理等を行う。
// DBのアクセス関数、レシーバメソッド、複雑になるロジックはモデル関数に定義する。

// ショップ一覧取得
func getshops(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータを取得
	pageStr := r.FormValue("page")
	perPageStr := r.FormValue("perPage")
	name := r.FormValue("name")

	// ページ、1ページあたりの件数の初期値
	page := 1
	perPage := 20
	var err error

	// ページ整数変換
	if(pageStr != "") {
		page, err = strconv.Atoi(pageStr)
	}
	if err != nil {
		fmt.Println(err)
		ResponseCommonError(w, http.StatusBadRequest, "不正なリクエストです。")
		return
	}

	// 1ページあたりの件数整数変換
	if(perPageStr != "") {
		perPage, err = strconv.Atoi(perPageStr)
	}
	if err != nil {
		fmt.Println(err)
		ResponseCommonError(w, http.StatusBadRequest, "不正なリクエストです。")
		return
	}

	// ショップ検索
	shops, err := models.SearchShops(page, perPage, name)
	if err != nil {
		fmt.Println(err)
		ResponseCommonError(w, http.StatusInternalServerError, "サーバーエラー")
		return
	}

	// ショップトータル件数取得
	total, err := models.FetchShopsTotal()
	if err != nil {
		fmt.Println(err)
		ResponseCommonError(w, http.StatusInternalServerError, "サーバーエラー")
		return
	}

	// ショップ一覧取得レスポンスの形に変換
	shopListRes := models.ShopListRes{
		Shops: shops,
		Pagination: models.Pagination{
			Page: page,
			PerPage: perPage,
			Total: total,
		},
	}

	shopListResJson, err := json.Marshal(shopListRes)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(shopListResJson))
}

// ショップ取得
func getShopById(w http.ResponseWriter, r *http.Request) {
	id := getShopPathParameter(r)
	if(id == "") {
		ResponseCommonError(w, http.StatusBadRequest, "不正なリクエストです。")
		return
	}

	shop, err := models.FetchShopById(id)
	if err != nil {
		fmt.Println(err)
		ResponseCommonError(w, http.StatusInternalServerError, "サーバーエラー")
		return
	}

	shopRes, err := json.Marshal(shop)
	if err != nil {
			fmt.Println(err)
			ResponseCommonError(w, http.StatusInternalServerError, "サーバーエラー")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(shopRes))
}

// ショップ作成
func postShop(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		Name	string	`json:"name"`
		Description	string	`json:"description"`
	}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		fmt.Println(err)
		ResponseCommonError(w, http.StatusInternalServerError, "サーバーエラー")
		return
	}

	// TODO 仮簡易バリデーション
	if reqBody.Name == "" || reqBody.Description == "" {
		fmt.Println("必須パラメータ")
		ResponseCommonError(w, http.StatusUnprocessableEntity, "パラメータが不正です。")
		return
	}
	if len(reqBody.Name) > 100 || len(reqBody.Description) == 2000 {
		fmt.Println("文字上限")
		ResponseCommonError(w, http.StatusUnprocessableEntity, "パラメータが不正です。")
		return
	}

	shop := models.Shop{
		Name: reqBody.Name,
		Description: reqBody.Description,
	}

	id, err := shop.CreateShop()
	if err != nil {
		fmt.Println(err)
		ResponseCommonError(w, http.StatusInternalServerError, "サーバーエラー")
		return
	}

	// 作成したIDのショップを取得
	shop, err = models.FetchShopById(id)
	if err != nil {
		fmt.Println(err)
		ResponseCommonError(w, http.StatusInternalServerError, "サーバーエラー")
		return
	}

	shopRes, err := json.Marshal(shop)
	if err != nil {
		fmt.Println(err)
		ResponseCommonError(w, http.StatusInternalServerError, "サーバーエラー")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(shopRes))
}

// ショップ更新
func putShopById(w http.ResponseWriter, r *http.Request) {
	id := getShopPathParameter(r)
	if(id == "") {
		ResponseCommonError(w, http.StatusBadRequest, "不正なリクエストです。")
		return
	}

	var reqBody struct {
		Name	string	`json:"name"`
		Description	string	`json:"description"`
	}

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		fmt.Println(err)
	}

	shop := models.Shop{
		Id: id,
		Name: reqBody.Name,
		Description: reqBody.Description,
	}

	err = shop.UpdateShopById()
	if err != nil {
		fmt.Println(err)
		ResponseCommonError(w, http.StatusInternalServerError, "サーバーエラー")
		return
	}

	// 更新したIDのショップを取得
	shop, err = models.FetchShopById(id)
	if err != nil {
		fmt.Println(err)
		ResponseCommonError(w, http.StatusInternalServerError, "サーバーエラー")
		return
	}

	shopRes, err := json.Marshal(shop)
	if err != nil {
		fmt.Println(err)
		ResponseCommonError(w, http.StatusInternalServerError, "サーバーエラー")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(shopRes))
}

// ショップ削除
func deleteShopById(w http.ResponseWriter, r *http.Request) {
	id := getShopPathParameter(r)
	if(id == "") {
		ResponseCommonError(w, http.StatusBadRequest, "不正なリクエストです。")
		return
	}

	err := models.DeleteShopById(id)
	if err != nil {
		fmt.Println(err)
		ResponseCommonError(w, http.StatusInternalServerError, "サーバーエラー")
		return
	}

	w.WriteHeader(http.StatusNoContent)
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
