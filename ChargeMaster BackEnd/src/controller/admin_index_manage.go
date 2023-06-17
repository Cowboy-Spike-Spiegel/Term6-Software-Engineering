package controller

import (
	"demo/src/consts"
	"demo/src/model"
	"demo/src/output"
	"demo/src/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminIndexManagerGet(c *gin.Context) {
	// get user id
	id, _ := c.Get("UserId")
	user_id := fmt.Sprintf("%s", id)
	output.Printer("Controller", "Body: user_id="+user_id)

	var response model.ResponseAdminIndexManageGet
	response.Code = consts.SUCCESS
	response.Msg = "Get station success"
	response.Data = service.StationInformation()
	c.JSON(http.StatusOK, response)
}

func AdminIndexManagerPost(c *gin.Context) {
	// get body
	var body model.ObjectAdminIndexManagePost
	if err := c.BindJSON(&body); err != nil {

		output.Printer("Service", "Manage body invalid")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "Manage body invalid",
			"data": nil,
		})
		return
	}

	// get user id
	id, _ := c.Get("UserId")
	user_id := fmt.Sprintf("%s", id)
	output.Printer("Controller", "Body: user_id="+user_id)

	// perform switch
	var rt bool
	if body.Mode == consts.SwitchON {
		rt = service.ChargeSwitchON(body.SwitchArray)
	} else if body.Mode == consts.SwitchOFF {
		rt = service.ChargeSwitchOFF(body.SwitchArray)
	} else {
		output.Printer("Controller", "Mode invalid")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "Mode invalid",
			"data": nil,
		})
	}
	if rt == false {
		output.Printer("Controller", "Switch fail")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "Switch fail",
			"data": nil,
		})
	} else {
		output.Printer("Controller", "Switch success")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.SUCCESS,
			"msg":  "Switch success",
			"data": nil,
		})
	}
}
