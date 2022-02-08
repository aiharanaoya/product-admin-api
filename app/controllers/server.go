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

// ルーティング設定
func SetRouter() {
	http.HandleFunc("/health_check", healthCheck)

	// ショップ
	http.HandleFunc("/v1/shops", shopsHandler)
	http.HandleFunc("/v1/shops/", shopsIdHandler)
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
		// 仮エラーハンドリング
		http.Error(w, "仮エラー", http.StatusInternalServerError)
	}
}
