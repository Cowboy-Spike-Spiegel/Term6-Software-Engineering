package controller

import (
	"demo/src/consts"
	"demo/src/model"
	"demo/src/output"
	"demo/src/service"
	"demo/src/system"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AdminIndexReportGet(c *gin.Context) {
	// get parameter
	charge_id := c.Query("charge_id")

	// get data in system
	id, err := strconv.Atoi(charge_id)
	if err != nil {
		output.Printer("Controller", "Get charge fail: body invalid")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "Get charge fail: body invalid",
			"data": nil,
		})
		return
	}
	if id > len(system.System.ChargeArea) || id <= 0 {
		output.Printer("Controller", "Get charge fail: id invalid")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "Get charge fail: id invalid",
			"data": nil,
		})
		return
	}

	// get user id
	userId, _ := c.Get("UserId")
	user_id := fmt.Sprintf("%s", userId)
	output.Printer("Controller", "Body: user_id="+user_id)

	var response model.ResponseAdminIndexReportGet
	response.Code = consts.SUCCESS
	response.Msg = "Get charge id =" + charge_id + " report"
	response.Data = service.ChargeReport(id)
	c.JSON(http.StatusOK, response)
}
