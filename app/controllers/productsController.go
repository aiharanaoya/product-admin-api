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

// 商品ハンドラ
func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getProducts(w, r)
	case http.MethodPost:
		postProduct(w, r)
	default:
		ResponseError(w, http.StatusNotFound)
	}
}

// 商品IDハンドラ
func productsIdHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getProductById(w, r)
	case http.MethodPut:
		putProductById(w, r)
	case http.MethodDelete:
		deleteProductById(w, r)
	default:
		ResponseError(w, http.StatusNotFound)
	}
}

// コントローラ関数
// それぞれのAPIに対応した関数。
// モデル関数で定義した構造体の呼び出し、JSONの変換処理等を行う。
// DBのアクセス関数、レシーバメソッド、複雑になるロジックはモデル関数に定義する。

// 商品一覧取得
func getProducts(w http.ResponseWriter, r *http.Request) {
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
		ResponseError(w, http.StatusBadRequest)
		return
	}

	// 1ページあたりの件数整数変換
	if(perPageStr != "") {
		perPage, err = strconv.Atoi(perPageStr)
	}
	if err != nil {
		fmt.Println(err)
		ResponseError(w, http.StatusBadRequest)
		return
	}

	// 商品検索
	prodcuts, err := models.SearchProducts(page, perPage, name)
	if err != nil {
		fmt.Println(err)
		ResponseError(w, http.StatusInternalServerError)
		return
	}

	// 商品トータル件数取得
	total, err := models.FetchProductsTotal()
	if err != nil {
		fmt.Println(err)
		ResponseError(w, http.StatusInternalServerError)
		return
	}

	// 商品一覧取得レスポンスの形に変換
	prodcutListRes := models.ProductListRes{
		Products: prodcuts,
		Pagination: models.Pagination{
			Page: page,
			PerPage: perPage,
			Total: total,
		},
	}

	prodcutListResJson, err := json.Marshal(prodcutListRes)
	if err != nil {
			fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(prodcutListResJson))
}

// 商品取得
func getProductById(w http.ResponseWriter, r *http.Request) {
	id := getProductPathParameter(r)
	if(id == "") {
		ResponseError(w, http.StatusBadRequest)
		return
	}

	prodcut, err := models.FetchProductById(id)
	if err != nil {
		fmt.Println(err)
		ResponseError(w, http.StatusInternalServerError)
		return
	}

	prodcutRes, err := json.Marshal(prodcut)
	if err != nil {
			fmt.Println(err)
			ResponseError(w, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(prodcutRes))
}

// 商品作成
func postProduct(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		Name	string	`json:"name"`
		Description	string	`json:"description"`
	}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		fmt.Println(err)
		ResponseError(w, http.StatusInternalServerError)
		return
	}

	// TODO 仮簡易バリデーション
	if reqBody.Name == "" || reqBody.Description == "" {
		fmt.Println("必須パラメータ")
		ResponseError(w, http.StatusUnprocessableEntity)
		return
	}
	if len(reqBody.Name) > 100 || len(reqBody.Description) == 2000 {
		fmt.Println("文字上限")
		ResponseError(w, http.StatusUnprocessableEntity)
		return
	}

	prodcut := models.Product{
		Name: reqBody.Name,
		Description: reqBody.Description,
	}

	id, err := prodcut.CreateProduct()
	if err != nil {
		fmt.Println(err)
		ResponseError(w, http.StatusInternalServerError)
		return
	}

	// 作成したIDの商品を取得
	prodcut, err = models.FetchProductById(id)
	if err != nil {
		fmt.Println(err)
		ResponseError(w, http.StatusInternalServerError)
		return
	}

	prodcutRes, err := json.Marshal(prodcut)
	if err != nil {
		fmt.Println(err)
		ResponseError(w, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(prodcutRes))
}

// 商品更新
func putProductById(w http.ResponseWriter, r *http.Request) {
	id := getProductPathParameter(r)
	if(id == "") {
		ResponseError(w, http.StatusBadRequest)
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

	prodcut := models.Product{
		Id: id,
		Name: reqBody.Name,
		Description: reqBody.Description,
	}

	err = prodcut.UpdateProductById()
	if err != nil {
		fmt.Println(err)
		ResponseError(w, http.StatusInternalServerError)
		return
	}

	// 更新したIDの商品を取得
	prodcut, err = models.FetchProductById(id)
	if err != nil {
		fmt.Println(err)
		ResponseError(w, http.StatusInternalServerError)
		return
	}

	prodcutRes, err := json.Marshal(prodcut)
	if err != nil {
		fmt.Println(err)
		ResponseError(w, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(prodcutRes))
}

// 商品削除
func deleteProductById(w http.ResponseWriter, r *http.Request) {
	id := getProductPathParameter(r)
	if(id == "") {
		ResponseError(w, http.StatusBadRequest)
		return
	}

	err := models.DeleteProductById(id)
	if err != nil {
		fmt.Println(err)
		ResponseError(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ロジック切り出し

// 商品のパスパラメータを取得する
func getProductPathParameter(r *http.Request) string {
	urls := strings.Split(r.RequestURI, "/")

	if(len(urls) < 4) {
		return ""
	}

	return urls[3]
}
