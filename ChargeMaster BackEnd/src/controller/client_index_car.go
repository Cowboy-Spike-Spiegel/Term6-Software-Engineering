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

func ClientIndexCarGet(c *gin.Context) {
	// get user id
	id, _ := c.Get("UserId")
	user_id := fmt.Sprintf("%s", id)
	output.Printer("Controller", "Body: user_id="+user_id)

	cars, ok := service.CarInformation(user_id)
	// generate response
	if ok {
		output.Printer("Controller", "Get car success")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.SUCCESS,
			"msg":  "Get car success",
			"data": cars,
		})
	} else {
		output.Printer("Controller", "Get car fail: user not exist")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "Get car fail",
			"data": cars,
		})
	}
}

func ClientIndexCarPost(c *gin.Context) {
	// get body
	var body model.ObjectClientIndexCarPost
	if err := c.BindJSON(&body); err != nil {
		output.Printer("Controller", "Body invalid")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.SUCCESS,
			"msg":  "CarPost body invalid",
			"data": nil,
		})
		return
	}

	// get user id
	id, _ := c.Get("UserId")
	user_id := fmt.Sprintf("%s", id)
	output.Printer("Controller", "Body: user_id="+user_id)

	// get attribute and update
	name := body.Name
	power_capacity := fmt.Sprintf("%f", body.PowerCapacity)
	power_current := fmt.Sprintf("%f", body.PowerCurrent)
	output.Printer("Controller", "Body: name="+name+", capacity="+power_capacity+", current="+power_current)

	// create car
	ok := service.CarCreate(user_id, name, power_capacity, power_current)
	if ok {
		output.Printer("Controller", "Insert car success")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.SUCCESS,
			"msg":  "Insert car success",
			"data": nil,
		})
	} else {
		output.Printer("Controller", "Insert car fail")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "Insert car fail",
			"data": nil,
		})
	}
}

func ClientIndexCarPatch(c *gin.Context) {
	// get body
	var body model.ObjectClientIndexCarPatch
	if err := c.BindJSON(&body); err != nil {
		output.Printer("Controller", "Body invalid")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.SUCCESS,
			"msg":  "CarPatch body invalid",
			"data": nil,
		})
		return
	}

	// get user id
	id, _ := c.Get("UserId")
	user_id := fmt.Sprintf("%s", id)
	output.Printer("Controller", "Body: user_id="+user_id)

	// get attribute and update
	car_id := body.CarID
	name := body.Name
	power_capacity := fmt.Sprintf("%f", body.PowerCapacity)
	power_current := fmt.Sprintf("%f", body.PowerCurrent)
	output.Printer("Controller", "Body: car_id="+car_id+", name="+name+", capacity="+power_capacity+", current="+power_current)

	// judge car_id exist & belong to user
	_, ok := dao.CarsQueryByIdUserId(car_id, user_id)
	if ok == false {
		output.Printer("Controller", "CarId invalid")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "CarId invalid",
			"data": nil,
		})
	}

	// update data from service
	if name != "" {
		ok := service.CarUpdate(car_id, dao.AttributeCar["Name"], name)
		if ok {
			output.Printer("Controller", "Update car name success")
			/*
				c.JSON(http.StatusOK, gin.H{
					"code": consts.SUCCESS,
					"msg":  "Update car name success",
					"data": nil,
				})

			*/
		} else {
			output.Printer("Controller", "Update car name fail")
			/*
				c.JSON(http.StatusOK, gin.H{
					"code": consts.FAIL,
					"msg":  "Update car name fail",
					"data": nil,
				})

			*/
		}
	}
	if power_capacity != "" {
		ok := service.CarUpdate(car_id, dao.AttributeCar["PowerCapacity"], power_capacity)
		if ok {
			output.Printer("Controller", "Update car power capacity success")
			/*
				c.JSON(http.StatusOK, gin.H{
					"code": consts.SUCCESS,
					"msg":  "Update car power capacity success",
					"data": nil,
				})

			*/
		} else {
			output.Printer("Controller", "Update car power capacity fail")
			/*
				c.JSON(http.StatusOK, gin.H{
					"code": consts.FAIL,
					"msg":  "Update car power capacity fail",
					"data": nil,
				})

			*/
		}
	}
	if power_current != "" {
		ok := service.CarUpdate(car_id, dao.AttributeCar["PowerCurrent"], power_current)
		if ok {
			output.Printer("Controller", "Update car power current success")
			c.JSON(http.StatusOK, gin.H{
				"code": consts.SUCCESS,
				"msg":  "Update car power current success",
				"data": nil,
			})
		} else {
			output.Printer("Controller", "Update car power current fail")
			c.JSON(http.StatusOK, gin.H{
				"code": consts.FAIL,
				"msg":  "Update car  power current fail",
				"data": nil,
			})
		}
	}
}

func ClientIndexCarDelete(c *gin.Context) {
	// get body
	var body model.ObjectClientIndexCarDelete
	if err := c.BindJSON(&body); err != nil {

		output.Printer("Controller", "Body invalid")
		c.JSON(http.StatusOK, gin.H{
			"code": consts.FAIL,
			"msg":  "CarDelete body invalid",
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusPaymentRequired, gin.H{
		"code": consts.FAIL,
		"msg":  "CarPatch body invalid",
		"data": nil,
	})
}
