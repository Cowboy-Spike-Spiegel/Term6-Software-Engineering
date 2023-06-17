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
)

func ClientIndexOrderGet(c *gin.Context) {
	mode := c.Query("mode")

	// get user id
	id, _ := c.Get("UserId")
	user_id := fmt.Sprintf("%s", id)
	output.Printer("Controller", "Body: user_id="+user_id)

	// search orders
	orders, ok := service.OrderSearch(user_id, mode)
	if ok == false {
		output.Printer("Service", "Search order fail")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "Search order fail",
			"data": nil,
		})
		return
	}

	// generate Number and FrontCars
	if mode == consts.SearchCurrent {
		for _, value := range orders {
			if value.State == consts.OrderWaiting {
				value.Number, value.FrontCars = system.WaitListFindNumberFrontCars(value.Mode, value.ID)
			} else {
				value.Number = system.ChargeFindNumber(value.State, value.Mode, value.CarID)
			}
		}
	}

	output.Printer("Service", "Search order success")
	var response model.ResponseClientIndexOrderGet
	response.Code = consts.SUCCESS
	response.Msg = "Search orders in mode=" + mode
	response.Data = orders
	c.JSON(http.StatusOK, response)
}

func ClientIndexOrderPost(c *gin.Context) {
	// get body
	var body model.ObjectClientIndexOrderPost
	if err := c.BindJSON(&body); err != nil {

		output.Printer("Service", "Order body invalid")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "Order body invalid",
			"data": nil,
		})
		return
	}

	// get user id
	id, _ := c.Get("UserId")
	user_id := fmt.Sprintf("%s", id)
	output.Printer("Controller", "Body: user_id="+user_id)

	// get parameters
	car_id := body.CarID
	mode := body.Mode
	apply_kwh := body.ApplyKwh

	ok := service.OrderCreate(user_id, car_id, mode, float32(apply_kwh))
	if ok {
		output.Printer("Controller", "Order create success")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.SUCCESS,
			"msg":  "Order create success",
			"data": nil,
		})
		return
	} else {
		output.Printer("Controller", "Order create fail")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "Order create fail",
			"data": nil,
		})
		return
	}
}

func ClientIndexOrderPatch(c *gin.Context) {
	// get body
	var body model.ObjectClientIndexOrderPatch
	if err := c.BindJSON(&body); err != nil {

		output.Printer("Service", "Order body invalid")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "Order body invalid",
			"data": nil,
		})
		return
	}

	order_id := body.ID
	mode := body.Mode
	new_kwh := body.ApplyKwh

	if new_kwh > 0 {
		ok := service.OrderChangeKwh(order_id, float32(new_kwh))
		if ok {
			output.Printer("Controller", "Order change apply_kwh success")
			c.JSON(http.StatusOK, gin.H{
				"code": consts.SUCCESS,
				"msg":  "Order change apply_kwh success",
				"data": nil,
			})
			return
		} else {
			output.Printer("Controller", "Order change apply_kwh fail")
			c.JSON(http.StatusOK, gin.H{
				"code": consts.FAIL,
				"msg":  "Order change apply_kwh fail",
				"data": nil,
			})
			return
		}
	} else {
		output.Printer("Controller", "Apply_kwh invalid")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "Apply_kwh invalid",
			"data": nil,
		})
		return
	}
	if mode != "" {
		if mode == consts.ChargeModeFast || mode == consts.ChargeModeSlow {
			ok := service.OrderChangeMode(order_id, mode)
			if ok {
				output.Printer("Controller", "Order switch mode success")
				c.JSON(http.StatusOK, gin.H{
					"code": consts.SUCCESS,
					"msg":  "Order switch mode success",
					"data": nil,
				})
				return
			} else {
				output.Printer("Controller", "Order switch mode fail")
				c.JSON(http.StatusOK, gin.H{
					"code": consts.FAIL,
					"msg":  "Order switch mode fail",
					"data": nil,
				})
				return
			}
		} else {
			output.Printer("Controller", "Mode invalid")
			c.JSON(http.StatusOK, gin.H{
				"code": consts.FAIL,
				"msg":  "Mode invalid",
				"data": nil,
			})
			return
		}
	}
}

func ClientIndexOrderDelete(c *gin.Context) {
	// get body
	var body model.ObjectClientIndexOrderDelete
	if err := c.BindJSON(&body); err != nil {

		output.Printer("Service", "Order body invalid")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "Order body invalid",
			"data": nil,
		})
		return
	}

	order_id := body.ID
	ok := system.SystemStopOrder(order_id)
	if ok {
		output.Printer("Controller", "Order close success")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.SUCCESS,
			"msg":  "Order close success",
			"data": nil,
		})
		return
	} else {
		output.Printer("Controller", "Order close fail")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "Order close fail",
			"data": nil,
		})
		return
	}
}
