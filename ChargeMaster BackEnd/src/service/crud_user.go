package service

import (
	"demo/src/consts"
	"demo/src/dao"
	"demo/src/model"
	"demo/src/output"
)

func UserLogin(account string, password string) (string, int, bool) {
	user, exist := dao.UsersQueryByAP(account, password)
	if exist == false {
		output.Printer("Service", "User not exist")
		return consts.NotExistId, consts.InvalidAuth, false
	}
	output.Printer("Service", "User exist")
	return user.Id, user.Auth, true
}

func AdminLogin(account string, password string) (string, int, bool) {
	user, exist := dao.UsersQueryAdminByAP(account, password)
	if exist == false {
		output.Printer("Service", "User not exist")
		return consts.NotExistId, consts.InvalidAuth, false
	}
	output.Printer("Service", "User exist")
	return user.Id, user.Auth, true
}

func UserSearch(account string) bool {
	_, exist := dao.UsersQueryByAccount(account)
	if exist == true {
		output.Printer("Service", "User found account="+account)
		return exist
	}
	output.Printer("Service", "User couldn't find account="+account)
	return exist
}

func UserRegister(account string, password string) (string, int, bool) {
	user, ok := dao.UsersInsert(account, password)
	if ok == false {
		output.Printer("Service", "User register fail")
		return consts.RepeatId, consts.InvalidAuth, false
	}
	output.Printer("Service", "User register success")
	return user.Id, user.Auth, true
}

func UserInformation(user_id string) (model.User, bool) {
	var information model.User

	// get data from dao, if no exist, return empty
	user, exist := dao.UsersQueryById(user_id)
	if exist == false {
		output.Printer("Service", "User not exist id="+user_id)
		return information, false
	}
	cars := dao.CarsQueryByUserId(user_id)

	// generate ModelUser
	information.ID = user.Id
	information.Account = user.Account
	information.Name = user.Name
	for _, value := range cars {
		var car model.Car
		car.CarID = value.Id
		car.UserID = value.UserId
		car.Name = value.Name
		car.PowerCapacity = float64(value.PowerCapacity)
		car.PowerCurrent = float64(value.PowerCurrent)
		information.CarArray = append(information.CarArray, car)
	}
	information.Balance = float64(user.Balance)
	output.Printer("Service", "Generate information for id="+user_id)
	return information, true
}

func UserChangePsdAuth(id string, password string) bool {
	_, exist := dao.UsersQueryByIdP(id, password)
	if exist {
		return true
	}
	return false
}

func UserUpdate(user_id string, attribute string, update string) bool {

	ok := dao.UsersUpdate(user_id, attribute, update)
	if ok == false {
		output.Printer("Service", "Update fail for id="+user_id)
		return false
	}
	output.Printer("Service", "Update success for id="+user_id)
	return true
}
