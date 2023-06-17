package dao

import (
	"demo/src/consts"
	"demo/src/output"
	"fmt"
)

func CarsQueryById(id string) (SchemaCar, bool) {
	var car SchemaCar
	car.Id = consts.NotExistId

	// query in Cars
	sql := "select * from cars where id="
	output.Printer("Dao", sql+" "+id)
	rows, err := db.Query(sql, id)
	if err != nil {
		output.Printer("Dao", err.Error())
		return car, false
	}

	if rows.Next() {
		err = rows.Scan(&car.Id, &car.UserId, &car.Name, &car.PowerCapacity, &car.PowerCurrent, &car.State)
		if err != nil {
			return car, false
		}
		return car, true
	}
	return car, false
}

func CarsQueryByUserId(user_id string) []SchemaCar {
	var cars []SchemaCar

	// query in Cars
	sql := "select * from cars where user_id=?"
	output.Printer("Dao", sql+" "+user_id)
	rows, err := db.Query(sql, user_id)
	if err != nil {
		fmt.Println(err)
		return cars
	}

	for rows.Next() {
		var car SchemaCar
		err = rows.Scan(&car.Id, &car.UserId, &car.Name, &car.PowerCapacity, &car.PowerCurrent, &car.State)
		if err != nil {
			fmt.Println(err)
			output.Printer("Dao", err.Error())
		}
		cars = append(cars, car)
	}
	return cars
}

func CarsQueryByIdUserId(id string, user_id string) (SchemaCar, bool) {
	var car SchemaCar
	car.Id = consts.NotExistId

	// query in Cars
	sql := "select * from cars where id=? and user_id=?"
	output.Printer("Dao", sql+" "+id+" "+user_id)
	rows, err := db.Query(sql, id, user_id)
	if err != nil {
		fmt.Println(err)
		return car, false
	}

	if rows.Next() {
		err = rows.Scan(&car.Id, &car.UserId, &car.Name, &car.PowerCapacity, &car.PowerCurrent, &car.State)
		if err != nil {
			fmt.Println(err)
			output.Printer("Dao", err.Error())
		}
		return car, true
	}
	return car, false
}

func CarsInsert(user_id string, name string, power_capacity string, power_current string) bool {

	// insert
	sql := "insert into cars (" +
		AttributeCar["UserId"] + ", " +
		AttributeCar["Name"] + ", " +
		AttributeCar["PowerCapacity"] + ", " +
		AttributeCar["PowerCurrent"] + ", " +
		AttributeCar["State"] +
		") values (?,?,?,?,?)"
	output.Printer("Dao", sql+" "+user_id+" "+name+" "+power_capacity+" "+power_current+" "+consts.CarUsing)

	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer stmt.Close()
	if _, err = stmt.Exec(user_id, name, power_capacity, power_current, consts.CarUsing); err != nil {
		fmt.Println(err)
		output.Printer("Dao", err.Error())
		return false
	}
	return true
}

func CarsUpdateById(id string, attribute string, update string) bool {
	sql := "update cars set " + attribute + "=? where id=?"
	output.Printer("Dao", sql+" "+" "+update+" "+id)
	_, err := db.Exec(sql, update, id)
	if err != nil {
		fmt.Println(err)
		output.Printer("Dao", err.Error())
		return false
	}
	return true
}

func CarsAddCurrentById(id string, add_kwh float32) bool {
	sql := "update cars set power_current=power_current+? where id=?"
	output.Printer("Dao", sql+" "+" "+fmt.Sprintf("%f", add_kwh)+" "+id)
	_, err := db.Exec(sql, add_kwh, id)
	if err != nil {
		fmt.Println(err)
		output.Printer("Dao", err.Error())
		return false
	}
	return true
}
