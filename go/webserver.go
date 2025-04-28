package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	routes.RegisterAPIRoutes(r)

	// todos エンドポイント
	r.GET("/todos", func(c *gin.Context) {
		todos := []Todo{
			{Id: 1, Title: "チャーハン！", Completed: true},
			{Id: 2, Title: "豚肉も入れるよ！", Completed: false},
		}
		c.JSON(http.StatusOK, todos)
	})

	// users エンドポイント
	r.GET("/users", func(c *gin.Context) {
		// サンプルのユーザーデータ
		users := []UserList{
			{UserID: "1", UserName: "ユーザーA", Biography: "これは例です", Status: "active"},
			{UserID: "2", UserName: "ユーザーB", Biography: "自己紹介文です", Status: "inactive"},
		}
		c.JSON(http.StatusOK, users)
	})

	// サーバーを指定ポートで起動
	return r.Run(":8080")
}
