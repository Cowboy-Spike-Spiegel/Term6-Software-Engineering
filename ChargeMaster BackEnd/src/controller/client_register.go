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

func ClientRegisterGet(c *gin.Context) {

}

func ClientRegisterPost(c *gin.Context) {
	// get body
	var body model.ObjectClientRegisterPost
	if err := c.BindJSON(&body); err != nil {

		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
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
	// name := body.Name

	// repeat
	exist := service.UserSearch(account)
	if exist == true {
		output.Printer("Controller", "Register with invalid account or password")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "Register account already exist",
			"data": gin.H{
				"token": nil,
			},
		})
		return
	}

	// format check
	if len(account) < consts.AccountMinLength || len(password) < consts.PasswordMinLength {
		output.Printer("Controller", "Register with invalid account or password")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "Register with invalid account or password",
			"data": gin.H{
				"token": nil,
			},
		})
		return
	} else {
		output.Printer("Controller", "Register with account="+account+", password="+password)

		// register with account and password
		user_id, user_auth, ok := service.UserRegister(account, password)
		if ok {
			output.Printer("Controller", "Register success.")
			// release token
			token, err := utils.ReleaseToken(user_id, user_auth)
			if err != nil {
				fmt.Println(err)
			}
			c.JSON(http.StatusOK, gin.H{
				"code": consts.SUCCESS,
				"msg":  "Register success",
				"data": gin.H{
					"token": token,
				},
			})
			return
		} else {
			output.Printer("Controller", "Register fail!")
			c.JSON(http.StatusOK, gin.H{
				"code": consts.FAIL,
				"msg":  "Register fail",
				"data": gin.H{
					"token": nil,
				},
			})
			return
		}
	}
}
