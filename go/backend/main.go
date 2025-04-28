package main

import (
	"log"

	"react-go-api/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // PostgreSQLドライバ（適宜変更）
)

// Todo の構造体
type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// UserList の構造体
type UserList struct {
	UserID    string `json:"user_id"`
	UserName  string `json:"user_name"`
	Biography string `json:"biography"`
	Status    string `json:"status"`
}

// StartWebServer は Web サーバーを起動します
func StartWebServer() error {
	r := gin.Default()

	// CORS設定の追加
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 必要に応じて特定のオリジンのみ許可
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	api.RegisterAPIRoutes(r)

	return r.Run(":8080")
}

func main() {
	// controllers パッケージから StartWebServer を呼び出す
	err := StartWebServer()
	if err != nil {
		log.Fatal(err)
	}
}
