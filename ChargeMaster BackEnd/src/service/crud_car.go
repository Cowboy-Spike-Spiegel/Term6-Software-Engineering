package service

import (
	"demo/src/dao"
	"demo/src/model"
	"demo/src/output"
)

func CarInformation(user_id string) ([]model.Car, bool) {
	var cars []model.Car

	// get data from dao, if no exist, return empty
	_, exist := dao.UsersQueryById(user_id)
	if exist == false {
		output.Printer("Service", "User Id invalid")
		return cars, false
	}

	dao_cars := dao.CarsQueryByUserId(user_id)
	for _, value := range dao_cars {
		var car model.Car
		car.CarID = value.Id
		car.UserID = value.UserId
		car.Name = value.Name
		car.PowerCapacity = float64(value.PowerCapacity)
		car.PowerCurrent = float64(value.PowerCurrent)
		cars = append(cars, car)
	}
	return cars, true
}

func CarCreate(user_id string, name string, power_capacity string, power_current string) bool {
	ok := dao.CarsInsert(user_id, name, power_capacity, power_current)
	if ok {
		output.Printer("Service", "Car create success")
		return true
	}
	output.Printer("Service", "User Id invalid")
	return false
}

func CarUpdate(car_id string, attribute string, update string) bool {
	ok := dao.CarsUpdateById(car_id, attribute, update)
	if ok {
		output.Printer("Service", "Car update success")
		return true
	}
	output.Printer("Service", "Car update fail")
	return false
}
