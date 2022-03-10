package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ヘルスチェック構造体
type HealthCheck struct {
	Status string `json:"status"`
}

// 共通エラー構造体
type CommonError struct {
	Message string `json:"message"`
}

// ルーティング設定
func SetRouter() {
	http.HandleFunc("/health_check", healthCheck)

	// ショップ
	http.HandleFunc("/api/shops", shopsHandler)
	http.HandleFunc("/api/shops/", shopsIdHandler)

	// 商品
	http.HandleFunc("/api/products", productsHandler)
	http.HandleFunc("/api/products/", productsIdHandler)

	// ユーザー
	http.HandleFunc("/api/users", usersHandler)
	http.HandleFunc("/api/users/login", usersLoginHandler)
	http.HandleFunc("/api/users/logout", usersLogoutHandler)
}

// サーバーを起動する
func StartServer() {
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}
}

// ヘルスチェック
func healthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		healthCheck := HealthCheck{Status: "OK"}

		healthCheckRes, err := json.Marshal(healthCheck)
		if err != nil {
				fmt.Println(err)
		}
		
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(healthCheckRes))
	} else {
		ResponseCommonError(w, http.StatusNotFound, "そのURLは存在しません。")
	}
}

// 共通エラーレスポンス
func ResponseCommonError(w http.ResponseWriter, statusCode int, message string) {
	error := CommonError{Message: message}

	errorRes, err := json.Marshal(error)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	fmt.Fprint(w, string(errorRes))
}
