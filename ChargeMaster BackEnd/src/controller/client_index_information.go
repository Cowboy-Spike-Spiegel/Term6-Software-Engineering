package controller

import (
	"demo/src/consts"
	"demo/src/dao"
	"demo/src/model"
	"demo/src/output"
	"demo/src/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ClientIndexInformationGet(c *gin.Context) {
	// get user id
	id, _ := c.Get("UserId")
	user_id := fmt.Sprintf("%s", id)
	output.Printer("Controller", "Body: user_id="+user_id)

	// get information from service
	user, ok := service.UserInformation(user_id)
	// generate response
	if ok {
		output.Printer("Controller", "Get information success")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.SUCCESS,
			"msg":  "Get information success",
			"data": user,
		})
	} else {
		output.Printer("Controller", "Get information fail")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "Get information fail",
			"data": nil,
		})
	}
}

func ClientIndexInformationPost(c *gin.Context) {
	// get body
	var body model.ObjectClientIndexInformationPost
	if err := c.BindJSON(&body); err != nil {

		output.Printer("Service", "Information body invalid")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.SUCCESS,
			"msg":  "Information body invalid",
			"data": nil,
		})
		return
	}

	// get user id
	id, _ := c.Get("UserId")
	user_id := fmt.Sprintf("%s", id)
	output.Printer("Controller", "Body: user_id="+user_id)

	// get attributes and updates
	old_password := body.OldPassword
	new_password := body.NewPassword
	name := body.Name
	output.Printer("Controller", "Body: old_psd="+old_password+", new_psd="+new_password+", name="+name)

	// update name
	if name != "" {
		ok := service.UserUpdate(user_id, dao.AttributeUser["Name"], name)
		if ok {
			output.Printer("Controller", "Update name success")
			c.JSON(http.StatusOK, gin.H{
				"code": consts.SUCCESS,
				"msg":  "Update name success",
				"data": nil,
			})
		} else {
			output.Printer("Controller", "Update name fail")
			c.JSON(http.StatusOK, gin.H{
				"code": consts.FAIL,
				"msg":  "Update name fail",
				"data": nil,
			})
		}
	}

	// update password
	if old_password != "" {
		// find user by user_id and old password
		ok := service.UserChangePsdAuth(user_id, old_password)
		if ok == false {
			output.Printer("Controller", "Update password with error old password")
			c.JSON(http.StatusOK, gin.H{
				"code": consts.FAIL,
				"msg":  "Update password with error old password",
				"data": nil,
			})
			return
		}

		ok = service.UserUpdate(user_id, dao.AttributeUser["Password"], new_password)
		if ok {
			output.Printer("Controller", "Update password success")
			c.JSON(http.StatusOK, gin.H{
				"code": consts.SUCCESS,
				"msg":  "Update password success",
				"data": nil,
			})
		} else {
			output.Printer("Controller", "Update information fail")
			c.JSON(http.StatusOK, gin.H{
				"code": consts.FAIL,
				"msg":  "Update password fail",
				"data": nil,
			})
		}
	}
}
