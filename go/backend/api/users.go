package api

import (
	// "fmt"
	"net/http"
	"react-go-api/controllers"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Id       string `json:"inputId" binding:"required"`
	Password string `json:"inputPassword" binding:"required"`
}

var jwtKey = []byte("your-secret-key") // 本番では環境変数などで管理

// JWTトークン生成
func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // 有効期限: 1時間
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// 暗号(Hash)化
// func PasswordEncrypt(password string) (string, error) {
// 	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	return string(hash), err
// }

func RegisterUserRoutes(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.POST("/login", func(c *gin.Context) {
			var req LoginRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
				return
			}

			// SQL文生成
			var query string = `SELECT 
									u_pwd, u_name, bio, u_img_path
								FROM 
									users 
								WHERE 
									u_status = 'active' AND 
									u_id = '` + req.Id + "'"

			results, err := controllers.ExecuteQuery(query, true)

			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err})
				return
			} else if len(results.([]map[string]interface{})) == 0 {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "ユーザーが見つかりません"})
				return
			} else if len(results.([]map[string]interface{})) == 1 {
				pwdFromDB := results.([]map[string]interface{})[0]["u_pwd"]
				err = bcrypt.CompareHashAndPassword([]byte(pwdFromDB.(string)), ([]byte(req.Password)))
				if err != nil {
					// パスワード不一致
					c.JSON(http.StatusUnauthorized, gin.H{"error": "ユーザー名またはパスワードが一致していません"})
				} else {
					// ログイン成功
					token, err := GenerateJWT(req.Id)
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": "トークン生成に失敗しました"})
						return
					}
					c.JSON(http.StatusOK, gin.H{
						"message":  "ログインに成功しました",
						"token":    token,
						"userName": results.([]map[string]interface{})[0]["u_name"].(string),
						"bio":      results.([]map[string]interface{})[0]["bio"].(string),
					})
				}
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "複数ユーザーが存在しています"})
				return
			}
		})

		// users.GET("", func(c *gin.Context) {
		// 	c.JSON(http.StatusOK, gin.H{"message": "List of users"})
		// })

		// users.GET("/:userid", func(c *gin.Context) {
		// 	userId := c.Param("userId")
		// 	c.JSON(http.StatusOK, gin.H{"message": "User details", "userId": userId})
		// })
	}
}
