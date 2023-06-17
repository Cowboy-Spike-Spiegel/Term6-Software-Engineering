package service

import (
	"demo/src/consts"
	"demo/src/dao"
	"demo/src/model"
	"demo/src/output"
	"demo/src/system"
	"fmt"
	"strconv"
	"time"
)

func OrderSearch(user_id string, mode string) ([]model.Order, bool) {
	var orders []model.Order
	if mode == consts.SearchHistory {
		// query state = "FINISHED"
		query_orders, _ := dao.OrdersQueryByUserIdState(user_id, consts.OrderFinished)
		for _, dao_order := range query_orders {
			orders = append(orders, OrderDao2Response(dao_order))
		}
		return orders, true
	} else if mode == consts.SearchCurrent {
		// query state = "CHARGING", "DISPATCHED", "WAITING"
		query_orders, _ := dao.OrdersQueryByUserIdState(user_id, consts.OrderCharging)
		for _, dao_order := range query_orders {
			orders = append(orders, OrderDao2Response(dao_order))
		}
		query_orders, _ = dao.OrdersQueryByUserIdState(user_id, consts.OrderDispatched)
		for _, dao_order := range query_orders {
			orders = append(orders, OrderDao2Response(dao_order))
		}
		query_orders, _ = dao.OrdersQueryByUserIdState(user_id, consts.OrderWaiting)
		for _, dao_order := range query_orders {
			orders = append(orders, OrderDao2Response(dao_order))
		}
		return orders, true
	}
	return orders, false
}

func OrderCreate(user_id string, car_id string, mode string, apply_kwh float32) bool {
	if system.SystemCanLineUp() == false {
		output.Printer("Service", "WaitArea is full")
		return false
	}

	// generate this order
	create_time := time.Now().Format("2006-01-02 15:04:05")
	dao.OrdersInsert(user_id, car_id, mode, apply_kwh, consts.OrderWaiting,
		consts.ChargeServicePrice, create_time)
	output.Printer("Service", "Order create success")
	// dao_order, _ := dao.OrdersQueryLatest()
	dao_order, _ := dao.OrdersQueryByUserIdCreateTime(user_id, create_time)
	order := OrderDao2System(dao_order)

	// generate order suppose_time & put it into WaitArea
	if system.SystemInsertOrder(order) == false {
		output.Printer("Service", "Order insert fail")
	}

	return true
}

func OrderChangeMode(order_id string, mode string) bool {
	return system.SystemChangeOrderMode(order_id, mode)
}

func OrderChangeKwh(order_id string, new_kwh float32) bool {
	return system.SystemChangeOrderApplyKwh(order_id, new_kwh)
}

func OrderComplete(order system.Order) {
	// get charge fee
	//var charge_fee float32

	//var service_fee float32
	//service_fee = consts.ChargeServicePrice * order.Duration
	fmt.Println(order)
}

func OrderDao2System(dao_order dao.SchemaOrder) system.Order {
	var order system.Order

	order.Id = dao_order.Id
	order.UserId = dao_order.UserId
	order.CarId = dao_order.CarId
	order.Mode = dao_order.Mode
	if dao_order.Mode == consts.ChargeModeFast {
		order.ChargeSpeed = consts.ChargeFastKwh
	} else if dao_order.Mode == consts.ChargeModeSlow {
		order.ChargeSpeed = consts.ChargeSlowKwh
	}
	order.ApplyKwh = dao_order.ApplyKwh
	order.ChargeKwh = dao_order.ChargeKwh
	order.State = dao_order.State
	order.ChargePrice = dao_order.ChargePrice
	order.ServicePrice = dao_order.ServicePrice
	order.CreateTime = dao_order.CreateTime
	order.DispatchTime = dao_order.DispatchTime
	charge_id, _ := strconv.Atoi(dao_order.ChargeId)
	order.ChargeId = charge_id
	order.StartTime = dao_order.StartTime
	order.FinishTime = dao_order.FinishTime

	return order
}

func OrderDao2Response(dao_order dao.SchemaOrder) model.Order {
	var order model.Order

	order.ID = dao_order.Id
	order.UserID = dao_order.UserId
	order.CarID = dao_order.CarId
	order.Mode = dao_order.Mode
	order.ApplyKwh = float64(dao_order.ApplyKwh)
	order.ChargeKwh = float64(dao_order.ChargeKwh)
	order.State = dao_order.State
	order.ChargePrice = float64(dao_order.ChargePrice)
	order.ServicePrice = float64(dao_order.ServicePrice)
	order.CreateTime = dao_order.CreateTime
	order.DispatchTime = dao_order.DispatchTime
	charge_id, _ := strconv.Atoi(dao_order.ChargeId)
	order.ChargeID = strconv.Itoa(charge_id)
	order.StartTime = dao_order.StartTime
	order.FinishTime = dao_order.FinishTime
	order.Fee = float64((dao_order.ChargePrice + dao_order.ServicePrice) * dao_order.ChargeKwh)

	return order
}
