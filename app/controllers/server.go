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

// エラー構造体
type Error struct {
	Status string `json:"status"`
}

// ルーティング設定
func SetRouter() {
	http.HandleFunc("/health_check", healthCheck)

	// ショップ
	http.HandleFunc("/api/shops", shopsHandler)
	http.HandleFunc("/api/shops/", shopsIdHandler)

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
		ResponseError(w, http.StatusNotFound)
	}
}

// エラーレスポンス
func ResponseError(w http.ResponseWriter, statusCode int) {
	error := Error{Status: http.StatusText(statusCode)}

	errorRes, err := json.Marshal(error)
	if err != nil {
			fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	fmt.Fprint(w, string(errorRes))
}