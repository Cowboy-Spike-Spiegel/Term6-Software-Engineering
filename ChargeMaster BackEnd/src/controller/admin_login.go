package controller

import (
	"demo/src/consts"
	"demo/src/model"
	"demo/src/output"
	"demo/src/service"
	"demo/src/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminLoginGet(c *gin.Context) {

}

func AdminLoginPost(c *gin.Context) {
	// get body
	var body model.ObjectAdminLoginPost
	if err := c.BindJSON(&body); err != nil {

		c.JSON(http.StatusOK, gin.H{
			"code": consts.SUCCESS,
			"msg":  "Login body invalid",
			"data": gin.H{
				"token": nil,
			},
		})
		return
	}

	// get parameters
	account := body.Account
	password := body.Password

	// query for account and password
	output.Printer("Controller", "Login with account="+account+", password="+password)
	user_id, user_auth, ok := service.AdminLogin(account, password)
	if ok {
		output.Printer("Controller", "Login success")
		// release token
		token, err := utils.ReleaseToken(user_id, user_auth)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"code": consts.SUCCESS,
			"msg":  "Login success",
			"data": gin.H{
				"token": token,
			},
		})
	} else {
		output.Printer("Controller", "Login fail")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "Login fail",
			"data": gin.H{
				"token": nil,
			},
		})
	}
}
