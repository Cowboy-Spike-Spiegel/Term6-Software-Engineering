package dao

import (
	"demo/src/consts"
	"demo/src/output"
	"fmt"
	"strconv"
	"time"
)

func OrdersQueryById(id string) (SchemaOrder, bool) {
	var order SchemaOrder
	order.Id = consts.NotExistId

	// query in orders
	sql := "select * from orders where id=?"
	output.Printer("Dao", sql+" "+id)
	rows, err := db.Query(sql, id)
	if err != nil {
		fmt.Println(err)
		return order, false
	}

	if rows.Next() {
		err = rows.Scan(&order.Id, &order.UserId, &order.CarId, &order.Mode, &order.ApplyKwh, &order.ChargeKwh,
			&order.State, &order.ChargePrice, &order.ServicePrice, &order.Fee, &order.CreateTime, &order.DispatchTime,
			&order.ChargeId, &order.StartTime, &order.FinishTime)
		if err != nil {
			fmt.Println(err)
			output.Printer("Dao", err.Error())
		}
		return order, true
	}
	return order, false
}

func OrdersQueryLatest() (SchemaOrder, bool) {
	var order SchemaOrder
	order.Id = consts.NotExistId

	// query in orders
	sql := "SELECT * from orders where id = (SELECT max(id) FROM orders)"
	output.Printer("Dao", sql)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
		return order, false
	}

	if rows.Next() {
		err = rows.Scan(&order.Id, &order.UserId, &order.CarId, &order.Mode, &order.ApplyKwh, &order.ChargeKwh,
			&order.State, &order.ChargePrice, &order.ServicePrice, &order.Fee, &order.CreateTime, &order.DispatchTime,
			&order.ChargeId, &order.StartTime, &order.FinishTime)
		if err != nil {
			fmt.Println(err)
			output.Printer("Dao", err.Error())
		}
		return order, true
	}
	return order, false
}

func OrdersQueryByUserId(user_id string, car_id string) (SchemaOrder, bool) {
	var order SchemaOrder
	order.Id = consts.NotExistId

	// query in orders
	sql := "select * from orders where user_id=? and car_id=?"
	output.Printer("Dao", sql+" "+user_id+" "+car_id)
	rows, err := db.Query(sql, user_id, car_id)
	if err != nil {
		fmt.Println(err)
		return order, false
	}

	if rows.Next() {
		err = rows.Scan(&order.Id, &order.UserId, &order.CarId, &order.Mode, &order.ApplyKwh, &order.ChargeKwh,
			&order.State, &order.ChargePrice, &order.ServicePrice, &order.Fee, &order.CreateTime, &order.DispatchTime,
			&order.ChargeId, &order.StartTime, &order.FinishTime)
		if err != nil {
			fmt.Println(err)
			output.Printer("Dao", err.Error())
		}
		return order, true
	}
	return order, false
}

func OrdersQueryByUserIdCarId(user_id string, car_id string) (SchemaOrder, bool) {
	var order SchemaOrder
	order.Id = consts.NotExistId

	// query in orders
	sql := "select * from orders where user_id=? and car_id=?"
	output.Printer("Dao", sql+" "+user_id+" "+car_id)
	rows, err := db.Query(sql, user_id, car_id)
	if err != nil {
		fmt.Println(err)
		return order, false
	}

	if rows.Next() {
		err = rows.Scan(&order.Id, &order.UserId, &order.CarId, &order.Mode, &order.ApplyKwh, &order.ChargeKwh,
			&order.State, &order.ChargePrice, &order.ServicePrice, &order.Fee, &order.CreateTime, &order.DispatchTime,
			&order.ChargeId, &order.StartTime, &order.FinishTime)
		if err != nil {
			fmt.Println(err)
			output.Printer("Dao", err.Error())
		}
		return order, true
	}
	return order, false
}

func OrdersQueryByUserIdState(user_id string, state string) ([]SchemaOrder, bool) {
	var orders []SchemaOrder

	// query in orders
	sql := "select * from orders where user_id=? and state=?"
	output.Printer("Dao", sql+" "+user_id+" "+state)
	rows, err := db.Query(sql, user_id, state)
	if err != nil {
		fmt.Println(err)
		return orders, false
	}

	for rows.Next() {
		var order SchemaOrder
		err = rows.Scan(&order.Id, &order.UserId, &order.CarId, &order.Mode, &order.ApplyKwh, &order.ChargeKwh,
			&order.State, &order.ChargePrice, &order.ServicePrice, &order.Fee, &order.CreateTime, &order.DispatchTime,
			&order.ChargeId, &order.StartTime, &order.FinishTime)
		if err != nil {
			fmt.Println(err)
			output.Printer("Dao", err.Error())
		}
		orders = append(orders, order)
	}
	return orders, true
}

func OrdersQueryByUserIdCreateTime(user_id string, create_time string) (SchemaOrder, bool) {
	var order SchemaOrder
	order.Id = consts.NotExistId

	// query in orders
	sql := "select * from orders where user_id=? and create_time=?"
	output.Printer("Dao", sql+" "+user_id+" "+create_time)
	rows, err := db.Query(sql, user_id, create_time)
	if err != nil {
		fmt.Println(err)
		return order, false
	}

	if rows.Next() {
		var order SchemaOrder
		err = rows.Scan(&order.Id, &order.UserId, &order.CarId, &order.Mode, &order.ApplyKwh, &order.ChargeKwh,
			&order.State, &order.ChargePrice, &order.ServicePrice, &order.Fee, &order.CreateTime, &order.DispatchTime,
			&order.ChargeId, &order.StartTime, &order.FinishTime)
		if err != nil {
			fmt.Println(err)
			output.Printer("Dao", err.Error())
			return order, false
		}
		return order, true
	}
	return order, false
}

func OrdersInsert(user_id string, car_id string, mode string, apply_kwh float32,
	state string, service_price float32, create_time string) bool {
	// insert
	sql := "insert into orders (" +
		AttributeOrder["UserId"] + ", " +
		AttributeOrder["CarId"] + ", " +
		AttributeOrder["Mode"] + ", " +
		AttributeOrder["ApplyKwh"] + ", " +
		AttributeOrder["ChargeKwh"] + ", " +
		AttributeOrder["State"] + ", " +
		AttributeOrder["ChargePrice"] + ", " +
		AttributeOrder["ServicePrice"] + ", " +
		AttributeOrder["Fee"] + ", " +
		AttributeOrder["CreateTime"] + ", " +
		AttributeOrder["DispatchTime"] + ", " +
		AttributeOrder["ChargeId"] + ", " +
		AttributeOrder["StartTime"] + ", " +
		AttributeOrder["FinishTime"] +
		") values (?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	output.Printer("Dao", sql)

	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	if _, err = stmt.Exec(user_id, car_id, mode, apply_kwh, 0, state, GetPrice(), service_price, 0, create_time,
		"xxxx-xx-xx xx:xx:xx", consts.NoId, "xxxx-xx-xx xx:xx:xx", "xxxx-xx-xx xx:xx:xx"); err != nil {
		fmt.Println(err)
		output.Printer("Dao", err.Error())
		return false
	}
	return true
}

func OrdersUpdateString(id string, attribute string, update string) bool {
	sql := "update orders set " + attribute + "=? where id=?"
	output.Printer("Dao", sql+" "+" "+update+" "+id)
	_, err := db.Exec(sql, update, id)
	if err != nil {
		fmt.Println(err)
		output.Printer("Dao", err.Error())
		return false
	}
	return true
}

func OrdersUpdateInt(id string, attribute string, update int) bool {
	sql := "update orders set " + attribute + "=? where id=?"
	obj := strconv.Itoa(update)
	output.Printer("Dao", sql+" "+" "+obj+" "+id)
	_, err := db.Exec(sql, update, id)
	if err != nil {
		fmt.Println(err)
		output.Printer("Dao", err.Error())
		return false
	}
	return true
}

func OrdersUpdateDifference(id string, attribute string, difference string) bool {
	return true
}

func GetPrice() float32 {
	hour := time.Now().Hour()
	if hour < 7 {
		return consts.ChargeLowPrice
	} else if hour < 10 {
		return consts.ChargeMiddlePrice
	} else if hour < 15 {
		return consts.ChargeHighPrice
	} else if hour < 18 {
		return consts.ChargeMiddlePrice
	} else if hour < 21 {
		return consts.ChargeHighPrice
	} else if hour < 23 {
		return consts.ChargeMiddlePrice
	} else {
		return consts.ChargeLowPrice
	}
}
