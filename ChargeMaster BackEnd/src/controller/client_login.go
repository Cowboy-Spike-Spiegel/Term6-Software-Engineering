package controller

import (
	"demo/src/consts"
	"demo/src/model"
	"demo/src/output"
	"demo/src/service"
	"demo/src/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func ClientLoginGet(c *gin.Context) {

}

func ClientLoginPost(c *gin.Context) {
	// get body
	var body model.ObjectClientLoginPost
	if err := c.BindJSON(&body); err != nil {
		data, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println("ERROR", data)
		}
		fmt.Println(data)
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

	// query for account and password
	output.Printer("Controller", "Login with account="+body.Account+", password="+body.Password)
	user_id, user_auth, ok := service.UserLogin(body.Account, body.Password)
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
