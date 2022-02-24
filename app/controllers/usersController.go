package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nao11aihara/product-admin-api/app/models"
)

// ハンドラ関数
// URL、HTTPメソッドから呼び出す関数をハンドリングする。
// 基本的にコントローラ関数を呼び出すのみで処理はコントローラ関数に記載する。

// ユーザーハンドラ
func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		postUser(w, r)
	default:
		ResponseError(w, http.StatusNotFound)
	}
}

// ユーザーログインハンドラ
func usersLoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		postUserLogin(w, r)
	default:
		ResponseError(w, http.StatusNotFound)
	}
}

// ユーザーログアウトハンドラ
func usersLogoutHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		deleteUserLogout(w, r)
	default:
		ResponseError(w, http.StatusNotFound)
	}
}

// コントローラ関数
// それぞれのAPIに対応した関数。
// モデル関数で定義した構造体の呼び出し、JSONの変換処理等を行う。
// DBのアクセス関数、レシーバメソッド、複雑になるロジックはモデル関数に定義する。

// ユーザー登録
func postUser(w http.ResponseWriter, r *http.Request) {
	userReq := models.UserReq{}

	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		fmt.Println(err)
	}

	// ユーザー登録する
	userRes, err := userReq.RegisterUser()
	if err != nil {
		fmt.Println(err)
	}

	// ユーザー登録後、そのままログインする
	sessionId, err := models.Login(userRes.UserId)
	if err != nil {
		fmt.Println(err)
	}

	// ログイン時生成したセッションIDをレスポンスに加える
	userRes.SessionId = sessionId

	// JSON変換
	userResJson, err := json.Marshal(userRes)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(userResJson))
}

// ログイン
func postUserLogin(w http.ResponseWriter, r *http.Request) {
	userReq := models.UserReq{}

	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		fmt.Println(err)
	}

	isOk, userRes, err := userReq.CheckLogin()
	if err != nil {
		fmt.Println(err)
	}

	if !isOk {
		// 仮
		http.Error(w, "ログイン失敗", http.StatusUnauthorized)
		return
	}

	sessionId, err := models.Login(userReq.UserId)
	if err != nil {
		fmt.Println(err)
	}

	// ログイン時生成したセッションIDをレスポンスに加える
	userRes.SessionId = sessionId

	// JSON変換
	userResJson, err := json.Marshal(userRes)
	if err != nil {
			fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(userResJson))
}

// ログアウト
func deleteUserLogout(w http.ResponseWriter, r *http.Request) {
	sessionId := r.Header.Get("sessionId")
	if sessionId == "" {
		ResponseError(w, http.StatusBadRequest)
	}

	err := 	models.Logout(sessionId)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusNoContent)
}
