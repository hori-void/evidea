package api

import (
	// "fmt"
	"errors"
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
func GenerateJWT(userId string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 1).Unix(), // 有効期限: 1時間
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ParseJWT(tokenString string) (string, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	userId, ok := claims["userId"].(string)
	if !ok {
		return "", errors.New("userId not found in token")
	}

	return userId, nil
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

		users.GET("/organizations", func(c *gin.Context) {
			token := c.Query("token")
			if token == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "token is required"})
				return
			}

			// JWTトークン解析（userId取得）
			claims := jwt.MapClaims{}
			parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

			if err != nil || !parsedToken.Valid {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
				return
			}

			userId, ok := claims["userId"].(string)
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "userId not found in token"})
				return
			}

			query := `
					SELECT 
						jo.j_org_id,
						jo.org_id,
						org.org_name,
						jo.admin_flg,
						CASE WHEN us.now_open_org = jo.org_id THEN true ELSE false END now_open
					FROM 
						join_org jo
					JOIN organizations org ON jo.org_id = org.org_id
					JOIN users us ON us.u_id = jo.u_id
					WHERE 
						jo.u_id = '` + userId + `' 
						AND org.status = 'active'
				`

			// クエリ実行（isSelect = true）
			result, err := controllers.ExecuteQuery(query, true)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "query execution failed"})
				return
			}

			// 返却
			c.JSON(http.StatusOK, result)
		})
	}

	// users.GET("/user/organizations", func(c *gin.Context) {
	// 	token := c.Query("token")
	// 	if token == "" {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": "token is required"})
	// 		return
	// 	}

	// 	// JWTトークンを解析してuserIdを取得
	// 	claims := jwt.MapClaims{}
	// 	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
	// 		return jwtKey, nil
	// 	})

	// 	if err != nil || !parsedToken.Valid {
	// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
	// 		return
	// 	}

	// 	userId, ok := claims["userId"].(string)
	// 	if !ok {
	// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "userId not found in token"})
	// 		return
	// 	}

	// 	// 組織情報をDBから取得
	// 	query := `
	// 		SELECT
	// 			jo.j_org_id,
	// 			jo.org_id,
	// 			org.org_name,
	// 			jo.admin_flg,
	// 			CASE WHEN us.now_open_org = jo.org_id THEN true ELSE false END now_open
	// 		FROM
	// 			join_org jo
	// 		JOIN organizations org ON jo.org_id = org.org_id
	// 		JOIN users us ON us.u_id = jo.u_id
	// 		WHERE
	// 			jo.u_id = ' ` + userId + `'
	// 			AND org.status = 'active'
	// 	`

	// 	results, err := controllers.ExecuteQuery(query, true)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "database query failed"})
	// 		return
	// 	}
	// 	// defer rows.Close()

	// 	type OrganizationInfo struct {
	// 		OrgID    string `json:"org_id"`
	// 		OrgName  string `json:"org_name"`
	// 		AdminFlg bool   `json:"admin_flg"`
	// 		NowOpen  bool   `json:"now_open"`
	// 	}

	// 	var orgs []OrganizationInfo

	// 	for i := 0; i < len(results.([]map[string]interface{})); i++ {
	// 		var org OrganizationInfo
	// 		if err := results.Scan(&org.OrgID, &org.OrgName, &org.AdminFlg, &org.NowOpen); err != nil {
	// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "scan failed"})
	// 			return
	// 		}
	// 		orgs = append(orgs, org)
	// 	}

	// 	for results.Next() {
	// 		var org OrganizationInfo
	// 		if err := results.Scan(&org.OrgID, &org.OrgName, &org.AdminFlg, &org.NowOpen); err != nil {
	// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "scan failed"})
	// 			return
	// 		}
	// 		orgs = append(orgs, org)
	// 	}

	// 	c.JSON(http.StatusOK, orgs)
	// })
}
