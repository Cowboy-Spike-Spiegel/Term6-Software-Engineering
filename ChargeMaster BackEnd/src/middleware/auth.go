package middleware

import (
	"demo/src/consts"
	"demo/src/output"
	"demo/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AUTHMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")

		// for login, go through middleware without authentic
		url := c.Request.URL.String()
		if strings.Contains(url, "/client/login") ||
			strings.Contains(url, "/admin/login") ||
			strings.Contains(url, "/client/register") {
			c.Next()
			return
		}

		// Parse token and authentic
		token_parse, claims, err := utils.ParseToken(token)
		if err != nil || !token_parse.Valid {
			output.Printer("Middleware", "Authentic fail. token invalid")
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": consts.FAIL,
				"msg":  "Authentic fail. No authority",
				"data": nil,
			})
			c.Abort()
		} else {
			user_id := claims.UserId
			auth := claims.Auth

			if strings.Contains(url, "admin") && auth < 1 {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code": consts.FAIL,
					"msg":  "Authentic fail. No Auth to enter admin page",
					"data": nil,
				})
				c.Abort()
			}
			// get userId in claims and judge
			if user_id == consts.NotExistId {
				output.Printer("Middleware", "Authentic fail. empty ID")
				c.JSON(http.StatusUnauthorized, gin.H{
					"code": consts.SUCCESS,
					"msg":  "Authentic fail. User not exist",
					"data": nil,
				})
				c.Abort()
			} else {
				output.Printer("Middleware", "Authentic success. ID = "+string(user_id))
				c.Set("UserId", user_id) // 用户存在将user的信息写入上下文，方便读取
				c.Next()
			}
		}
	}
}
