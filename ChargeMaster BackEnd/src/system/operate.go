package system

import (
	"demo/src/consts"
	"demo/src/dao"
	"demo/src/output"
	"sort"
	"strconv"
	"time"
)

func SystemCanLineUp() bool {
	// wait array full, no block to line up
	return System.WaitArea.FastWaitSize+System.WaitArea.SlowWaitSize < System.WaitArea.Capacity
}

func SystemInsertOrder(order Order) bool {
	// search car, and generate name, powers
	car, ok := dao.CarsQueryByIdUserId(order.CarId, order.UserId)
	if ok == false {
		return false
	}

	// Number
	order.Number = GetNumber()
	// car power
	order.Name = car.Name
	order.PowerCapacity = car.PowerCapacity
	order.PowerCurrent = car.PowerCurrent

	// line up in wait array
	//System.WaitArea.Mutex.Lock()
	if order.Mode == consts.ChargeModeFast {
		// Fast Mode
		order.ChargeSpeed = consts.ChargeFastKwh
		System.WaitArea.FastWaits = append(System.WaitArea.FastWaits, order)
		System.WaitArea.FastWaitSize++
	} else if order.Mode == consts.ChargeModeSlow {
		// Slow Mode
		order.ChargeSpeed = consts.ChargeSlowKwh
		System.WaitArea.SlowWaits = append(System.WaitArea.SlowWaits, order)
		System.WaitArea.SlowWaitSize++
	}
	//fmt.Println(System.WaitArea)
	//System.WaitArea.Mutex.Unlock()

	return true
}

func SystemCompleteOrder(order Order) {
	dao.OrdersUpdateString(order.Id, dao.AttributeOrder["ChargeKwh"],
		strconv.FormatFloat(float64(order.ChargeKwh), 'f', 6, 64))
	dao.OrdersUpdateString(order.Id, dao.AttributeOrder["State"], order.State)
	dao.OrdersUpdateString(order.Id, dao.AttributeOrder["ChargePrice"],
		strconv.FormatFloat(float64(order.ChargePrice), 'f', 6, 64))
	dao.OrdersUpdateString(order.Id, dao.AttributeOrder["ServicePrice"],
		strconv.FormatFloat(float64(order.ServicePrice), 'f', 6, 64))
	dao.OrdersUpdateString(order.Id, dao.AttributeOrder["Fee"],
		strconv.FormatFloat(float64(order.Fee), 'f', 6, 64))
	dao.OrdersUpdateString(order.Id, dao.AttributeOrder["DispatchTime"], order.DispatchTime)
	//dao.OrdersUpdateInt(order.Id, dao.AttributeOrder["ChargeId"], order.ChargeId)
	dao.OrdersUpdateString(order.Id, dao.AttributeOrder["StartTime"], order.StartTime)
	dao.OrdersUpdateString(order.Id, dao.AttributeOrder["FinishTime"], order.FinishTime)
	//dao.CarsAddCurrentById(order.CarId, order.ChargeKwh)
}

func SystemStopOrder(order_id string) bool {
	order, ok := dao.OrdersQueryById(order_id)
	if ok == false {
		output.Printer("System", "Close order id invalid")
		return false
	}
	//fmt.Println("\nBefore delete")
	//SystemPrintState()
	// finished, forbidden
	if order.State == consts.OrderFinished {
		output.Printer("System", "Order finished")
		return false
	} else if order.State == consts.CarCharging {
		// stop charging
		if order.Mode == consts.ChargeModeFast {
			for _, id := range System.Control.FastId {
				// find it
				if System.ChargeArea[Id2Index(id)].Orders[0].Id == order_id {
					// stop this order
					order := System.ChargeArea[Id2Index(id)].Orders[0]
					System.ChargeArea[Id2Index(id)].Orders = System.ChargeArea[Id2Index(id)].Orders[1:]

					// update charge
					System.ChargeArea[Id2Index(id)].TotalChargeTimes++
					System.ChargeArea[Id2Index(id)].State = consts.ChargeRest
					System.ChargeArea[Id2Index(id)].Size--
					System.ChargeArea[Id2Index(id)].Count = 0

					// update control
					System.Control.FastRemain++

					// insert this order
					order.FinishTime = time.Now().Format("2006-01-02 15:04:05")
					if order.ChargeKwh > 0 {
						order.ChargePrice = order.ChargeFee / order.ChargeKwh
					} else {
						order.ChargePrice = GetPrice()
					}
					order.State = consts.OrderFinished
					SystemCompleteOrder(order)
					break
				}
			}
		} else if order.Mode == consts.ChargeModeSlow {
			for _, id := range System.Control.SlowId {
				// find it
				if System.ChargeArea[Id2Index(id)].Orders[0].Id == order_id {
					// stop this order
					order := System.ChargeArea[Id2Index(id)].Orders[0]
					System.ChargeArea[Id2Index(id)].Orders = System.ChargeArea[Id2Index(id)].Orders[1:]

					// update charge
					System.ChargeArea[Id2Index(id)].TotalChargeTimes++
					System.ChargeArea[Id2Index(id)].State = consts.ChargeRest
					System.ChargeArea[Id2Index(id)].Size--
					System.ChargeArea[Id2Index(id)].Count = 0

					// update control
					System.Control.SlowRemain++

					// insert this order
					order.FinishTime = time.Now().Format("2006-01-02 15:04:05")
					if order.ChargeKwh > 0 {
						order.ChargePrice = order.ChargeFee / order.ChargeKwh
					} else {
						order.ChargePrice = GetPrice()
					}

					order.State = consts.OrderFinished
					SystemCompleteOrder(order)
					break
				}
			}
		}
		//fmt.Println("\nCharging delete")
	} else if order.State == consts.OrderDispatched {
		// stop charging
		if order.Mode == consts.ChargeModeFast {

			for _, id := range System.Control.FastId {
				// find it
				if System.ChargeArea[Id2Index(id)].Orders[1].Id == order_id {
					// stop this order
					order := System.ChargeArea[Id2Index(id)].Orders[1]
					System.ChargeArea[Id2Index(id)].Orders =
						System.ChargeArea[Id2Index(id)].Orders[:len(System.ChargeArea[Id2Index(id)].Orders)-1]

					// update charge
					System.ChargeArea[Id2Index(id)].Size--

					// update control
					System.Control.FastRemain++

					// insert this order
					order.FinishTime = time.Now().Format("2006-01-02 15:04:05")
					order.State = consts.OrderFinished
					SystemCompleteOrder(order)
				}
			}
		} else if order.Mode == consts.ChargeModeSlow {
			for _, id := range System.Control.SlowId {
				// find it
				if System.ChargeArea[Id2Index(id)].Orders[1].Id == order_id {
					// stop this order
					order := System.ChargeArea[Id2Index(id)].Orders[1]
					System.ChargeArea[Id2Index(id)].Orders =
						System.ChargeArea[Id2Index(id)].Orders[:len(System.ChargeArea[Id2Index(id)].Orders)-1]

					// update charge
					System.ChargeArea[Id2Index(id)].Size--

					// update control
					System.Control.SlowRemain++

					// insert this order
					order.FinishTime = time.Now().Format("2006-01-02 15:04:05")
					order.State = consts.OrderFinished
					SystemCompleteOrder(order)
					break
				}
			}
		}
		//fmt.Println("\nDispatch delete")
	} else if order.State == consts.OrderWaiting {
		orders := []Order{}
		if order.Mode == consts.ChargeModeFast {
			for _, obj := range System.WaitArea.FastWaits {
				// find it
				if obj.Id == order_id {
					// stop this order
					order := obj
					// update wait area
					System.WaitArea.FastWaitSize--
					// insert this order
					order.FinishTime = time.Now().Format("2006-01-02 15:04:05")
					order.State = consts.OrderFinished
					SystemCompleteOrder(order)
				} else {
					orders = append(orders, obj)
				}
			}
			System.WaitArea.FastWaits = orders
		} else if order.Mode == consts.ChargeModeSlow {
			for _, obj := range System.WaitArea.SlowWaits {
				// find it
				if obj.Id == order_id {
					// stop this order
					order := obj
					// update wait area
					System.WaitArea.SlowWaitSize--
					// insert this order
					order.FinishTime = time.Now().Format("2006-01-02 15:04:05")
					order.State = consts.OrderFinished
					SystemCompleteOrder(order)
				} else {
					orders = append(orders, obj)
				}
			}
			System.WaitArea.FastWaits = orders
		}
		//fmt.Println("\nWaiting delete")
	}
	SystemPrintState()
	return true
}

func SystemChangeOrderMode(order_id string, mode string) bool {
	order, ok := dao.OrdersQueryById(order_id)
	if ok == false {
		output.Printer("System", "Order id="+order_id+"(not exist) change mode fail")
		return false
	}

	// not waiting, can't convert
	if order.State != consts.OrderWaiting {
		output.Printer("System", "Order id="+order_id+" not in waiting")
		return false
	}

	// convert and re line up
	var orders []Order
	var trans_order Order
	trans_order.Id = strconv.Itoa(consts.NoId)
	if order.Mode == consts.ChargeModeFast {
		for _, obj := range System.WaitArea.FastWaits {
			if obj.Id == order_id {
				trans_order = obj
			} else {
				orders = append(orders, obj)
			}
		}
		// exactly find it
		if order.Id != strconv.Itoa(consts.NoId) {
			// transfer to slow queue
			System.WaitArea.FastWaits = orders
			System.WaitArea.FastWaitSize--
			trans_order.Number = GetNumber()
			trans_order.ChargeSpeed = consts.ChargeSlowKwh
			System.WaitArea.SlowWaits = append(System.WaitArea.SlowWaits, trans_order)
			System.WaitArea.SlowWaitSize++
		}
	} else if order.Mode == consts.ChargeModeSlow {
		for _, obj := range System.WaitArea.SlowWaits {
			if obj.Id == order_id {
				trans_order = obj
			} else {
				orders = append(orders, obj)
			}
		}
		// exactly find it
		if order.Id != strconv.Itoa(consts.NoId) {
			// transfer to fast queue
			System.WaitArea.SlowWaits = orders
			System.WaitArea.SlowWaitSize--
			trans_order.Number = GetNumber()
			trans_order.ChargeSpeed = consts.ChargeFastKwh
			System.WaitArea.FastWaits = append(System.WaitArea.FastWaits, trans_order)
			System.WaitArea.FastWaitSize++
		}
	}
	return true
}

func SystemChangeOrderApplyKwh(order_id string, new_kwh float32) bool {
	order, ok := dao.OrdersQueryById(order_id)
	if ok == false {
		output.Printer("System", "Order id="+order_id+"(not exist) change apply_kwh fail")
		return false
	}

	// not waiting, can't convert
	if order.State != consts.OrderWaiting {
		output.Printer("System", "Order id="+order_id+" not in waiting")
		return false
	}

	// convert kwh
	if order.Mode == consts.ChargeModeFast {
		for index, _ := range System.WaitArea.FastWaits {
			if System.WaitArea.FastWaits[index].Id == order_id {
				System.WaitArea.FastWaits[index].ApplyKwh = new_kwh
			}
		}
	} else if order.Mode == consts.ChargeModeSlow {
		for index, _ := range System.WaitArea.SlowWaits {
			if System.WaitArea.SlowWaits[index].Id == order_id {
				System.WaitArea.SlowWaits[index].ApplyKwh = new_kwh
			}
		}
	}
	return true
}

func ChargeOpenById(id string) bool {
	charge_id, err := strconv.Atoi(id)
	if err != nil {
		output.Printer("System", "Invalid charge_id")
		return false
	}

	// change state if state == OFF
	if System.ChargeArea[Id2Index(charge_id)].State == consts.ChargeOFF {
		// remove dispatched cars into wait area
		ok := dao.ChargesUpdateById(id, dao.AttributeCharge["State"], consts.ChargeRest)
		if ok == false {
			output.Printer("Service", "Charge not exist, id="+id)
			return false
		}

		// change charge state
		System.ChargeArea[Id2Index(charge_id)].State = consts.ChargeRest
		output.Printer("Service", "Charge open, id="+id)

		// change system control parameters
		if System.ChargeArea[Id2Index(charge_id)].Mode == consts.ChargeModeFast {
			System.Control.FastId = append(System.Control.FastId, charge_id)
			System.Control.FastCapacity += consts.ChargeLimit
			System.Control.FastRemain += consts.ChargeLimit
		} else if System.ChargeArea[Id2Index(charge_id)].Mode == consts.ChargeModeSlow {
			System.Control.SlowId = append(System.Control.SlowId, charge_id)
			System.Control.SlowCapacity += consts.ChargeLimit
			System.Control.SlowRemain += consts.ChargeLimit
		}
		// change system charge parameters
		System.ChargeArea[Id2Index(charge_id)].Size = 0
		System.ChargeArea[Id2Index(charge_id)].TotalCharge = 0
		System.ChargeArea[Id2Index(charge_id)].TotalChargeTimes = 0
		System.ChargeArea[Id2Index(charge_id)].TotalChargeDuration = 0
		System.ChargeArea[Id2Index(charge_id)].TotalChargeFee = 0
		System.ChargeArea[Id2Index(charge_id)].TotalServiceFee = 0
		System.ChargeArea[Id2Index(charge_id)].TotalFee = 0
		System.ChargeArea[Id2Index(charge_id)].Count = 0

		// remove same state dispatched to wait
		if System.ChargeArea[Id2Index(charge_id)].Mode == consts.ChargeModeFast {
			for _, id := range System.Control.FastId {
				if System.ChargeArea[Id2Index(id)].Size == consts.ChargeLimit {
					// remove
					order := System.ChargeArea[Id2Index(id)].Orders[1]
					System.ChargeArea[Id2Index(id)].Orders = System.ChargeArea[Id2Index(id)].Orders[:1]
					System.ChargeArea[Id2Index(id)].Size--

					// into wait area
					System.WaitArea.FastWaits = append(System.WaitArea.FastWaits, order)
					System.WaitArea.FastWaitSize++
				}
			}
			// sort, for next dispatch
			sort.Slice(System.WaitArea.FastWaits, func(i, j int) bool {
				return System.WaitArea.FastWaits[i].Number < System.WaitArea.FastWaits[j].Number
			})
		} else if System.ChargeArea[Id2Index(charge_id)].Mode == consts.ChargeModeSlow {
			for _, id := range System.Control.SlowId {
				if System.ChargeArea[Id2Index(id)].Size == consts.ChargeLimit {
					// remove
					order := System.ChargeArea[Id2Index(id)].Orders[1]
					System.ChargeArea[Id2Index(id)].Orders = System.ChargeArea[Id2Index(id)].Orders[:1]
					System.ChargeArea[Id2Index(id)].Size--

					// into wait area
					System.WaitArea.SlowWaits = append(System.WaitArea.SlowWaits, order)
					System.WaitArea.SlowWaitSize++
				}
			}
			// sort, for next dispatch
			sort.Slice(System.WaitArea.SlowWaits, func(i, j int) bool {
				return System.WaitArea.SlowWaits[i].Number < System.WaitArea.SlowWaits[j].Number
			})
		}
		SystemPrintState()
	} else if System.ChargeArea[Id2Index(charge_id)].State == consts.ChargeFault {
		// state = fault, can't change
		output.Printer("System", "charge id="+id+" is fault")
		return false

	} else if System.ChargeArea[Id2Index(charge_id)].State == consts.ChargeWork ||
		System.ChargeArea[Id2Index(charge_id)].State == consts.ChargeRest {
		// is open, can't open
		output.Printer("System", "charge id="+id+" has been opened")
		return false
	}
	return true
}

func ChargeCloseById(id string) bool {
	charge_id, err := strconv.Atoi(id)
	if err != nil {
		output.Printer("System", "Invalid charge_id")
		return false
	}

	// change state if state == OFF
	if System.ChargeArea[Id2Index(charge_id)].State == consts.ChargeWork ||
		System.ChargeArea[Id2Index(charge_id)].State == consts.ChargeRest {
		// end charge orders, and generate new order with same number

		// stop the 2ed order(with parameters change)
		if System.ChargeArea[Id2Index(charge_id)].Size >= consts.ChargeLimit {
			order := System.ChargeArea[Id2Index(charge_id)].Orders[consts.ChargeLimit-1]
			System.ChargeArea[Id2Index(charge_id)].Orders = System.ChargeArea[Id2Index(charge_id)].Orders[:1]

			// update charge
			System.ChargeArea[Id2Index(charge_id)].Size--
			// update control
			if order.Mode == consts.ChargeModeFast {
				System.Control.FastRemain++
			} else if order.Mode == consts.ChargeModeSlow {
				System.Control.SlowRemain++
			}
			// insert this order
			order.FinishTime = time.Now().Format("2006-01-02 15:04:05")
			if order.ChargeKwh > 0 {
				order.ChargePrice = order.ChargeFee / order.ChargeKwh
			} else {
				order.ChargePrice = GetPrice()
			}
			order.State = consts.OrderFinished
			SystemCompleteOrder(order)

			// reset this new order
			order.ApplyKwh = order.ApplyKwh - order.ChargeKwh
			order.ChargeKwh = 0
			order.State = consts.OrderWaiting
			order.ChargePrice = 0
			order.ServicePrice = consts.ChargeServicePrice
			order.ChargeFee = 0
			order.ServiceFee = 0
			order.Fee = 0
			order.CreateTime = time.Now().Format("2006-01-02 15:04:05")
			order.DispatchTime = consts.TimeDefualt
			order.ChargeId = 0
			order.StartTime = consts.TimeDefualt
			order.FinishTime = consts.TimeDefualt

			// create and insert it
			dao.OrdersInsert(order.UserId, order.CarId, order.Mode, order.ApplyKwh,
				order.State, order.ServicePrice, order.CreateTime)

			dao_order, _ := dao.OrdersQueryLatest()
			order.Id = dao_order.Id
			if order.Mode == consts.ChargeModeFast {
				System.WaitArea.FastWaits = append(System.WaitArea.FastWaits, order)
				System.WaitArea.FastWaitSize++
			} else if order.Mode == consts.ChargeModeSlow {
				System.WaitArea.SlowWaits = append(System.WaitArea.SlowWaits, order)
				System.WaitArea.SlowWaitSize++
			}
		}

		// stop the 1st order(with parameters change)
		if System.ChargeArea[Id2Index(charge_id)].Size > 0 {
			order := System.ChargeArea[Id2Index(charge_id)].Orders[0]
			System.ChargeArea[Id2Index(charge_id)].Orders = System.ChargeArea[Id2Index(charge_id)].Orders[1:]
			// update charge
			System.ChargeArea[Id2Index(charge_id)].TotalChargeTimes++
			System.ChargeArea[Id2Index(charge_id)].State = consts.ChargeRest
			System.ChargeArea[Id2Index(charge_id)].Size--
			System.ChargeArea[Id2Index(charge_id)].Count = 0
			// update control
			if order.Mode == consts.ChargeModeFast {
				System.Control.FastRemain++
			} else if order.Mode == consts.ChargeModeSlow {
				System.Control.SlowRemain++
			}
			// insert this order
			order.FinishTime = time.Now().Format("2006-01-02 15:04:05")
			if order.ChargeKwh > 0 {
				order.ChargePrice = order.ChargeFee / order.ChargeKwh
			} else {
				order.ChargePrice = GetPrice()
			}
			order.State = consts.OrderFinished
			SystemCompleteOrder(order)

			// reset this new order
			order.ApplyKwh = order.ApplyKwh - order.ChargeKwh
			order.ChargeKwh = 0
			order.State = consts.OrderWaiting
			order.ChargePrice = 0
			order.ServicePrice = consts.ChargeServicePrice
			order.ChargeFee = 0
			order.ServiceFee = 0
			order.Fee = 0
			order.CreateTime = time.Now().Format("2006-01-02 15:04:05")
			order.DispatchTime = consts.TimeDefualt
			order.ChargeId = 0
			order.StartTime = consts.TimeDefualt
			order.FinishTime = consts.TimeDefualt

			// create and insert it
			dao.OrdersInsert(order.UserId, order.CarId, order.Mode, order.ApplyKwh,
				order.State, order.ServicePrice, order.CreateTime)
			dao_order, _ := dao.OrdersQueryLatest()
			order.Id = dao_order.Id
			if order.Mode == consts.ChargeModeFast {
				System.WaitArea.FastWaits = append(System.WaitArea.FastWaits, order)
				System.WaitArea.FastWaitSize++
			} else if order.Mode == consts.ChargeModeSlow {
				System.WaitArea.SlowWaits = append(System.WaitArea.SlowWaits, order)
				System.WaitArea.SlowWaitSize++
			}
		}

		// change charge state
		System.ChargeArea[Id2Index(charge_id)].State = consts.ChargeOFF
		output.Printer("Service", "Charge close, id="+id)

		// change system control parameters
		if System.ChargeArea[Id2Index(charge_id)].Mode == consts.ChargeModeFast {
			// delete this charge_id
			var ids []int
			for _, value := range System.Control.FastId {
				if value != charge_id {
					ids = append(ids, value)
				}
			}
			System.Control.FastId = ids
			System.Control.FastCapacity -= consts.ChargeLimit
			System.Control.FastRemain -= consts.ChargeLimit
		} else if System.ChargeArea[Id2Index(charge_id)].Mode == consts.ChargeModeSlow {
			// delete this charge_id
			var ids []int
			for _, value := range System.Control.SlowId {
				if value != charge_id {
					ids = append(ids, value)
				}
			}
			System.Control.SlowId = ids
			System.Control.SlowCapacity -= consts.ChargeLimit
			System.Control.SlowRemain -= consts.ChargeLimit
		}
		// change system charge parameters
		/*
			System.ChargeArea[Id2Index(charge_id)].Size = 0
			System.ChargeArea[Id2Index(charge_id)].TotalCharge = 0
			System.ChargeArea[Id2Index(charge_id)].TotalChargeTimes = 0
			System.ChargeArea[Id2Index(charge_id)].TotalChargeDuration = 0
			System.ChargeArea[Id2Index(charge_id)].TotalChargeFee = 0
			System.ChargeArea[Id2Index(charge_id)].TotalServiceFee = 0
			System.ChargeArea[Id2Index(charge_id)].TotalFee = 0

		*/

		// remove same state dispatched to wait
		if System.ChargeArea[Id2Index(charge_id)].Mode == consts.ChargeModeFast {
			for _, id := range System.Control.FastId {
				if System.ChargeArea[Id2Index(id)].Size == consts.ChargeLimit {
					// remove
					order := System.ChargeArea[Id2Index(id)].Orders[1]
					System.ChargeArea[Id2Index(id)].Orders = System.ChargeArea[Id2Index(id)].Orders[:1]
					System.ChargeArea[Id2Index(id)].Size--

					// into wait area
					System.WaitArea.FastWaits = append(System.WaitArea.FastWaits, order)
					System.WaitArea.FastWaitSize++
				}
			}
			// sort, for next dispatch
			sort.Slice(System.WaitArea.FastWaits, func(i, j int) bool {
				return System.WaitArea.FastWaits[i].Number < System.WaitArea.FastWaits[j].Number
			})
		} else if System.ChargeArea[Id2Index(charge_id)].Mode == consts.ChargeModeSlow {
			for _, id := range System.Control.SlowId {
				if System.ChargeArea[Id2Index(id)].Size == consts.ChargeLimit {
					// remove
					order := System.ChargeArea[Id2Index(id)].Orders[1]
					System.ChargeArea[Id2Index(id)].Orders = System.ChargeArea[Id2Index(id)].Orders[:1]
					System.ChargeArea[Id2Index(id)].Size--

					// into wait area
					System.WaitArea.SlowWaits = append(System.WaitArea.SlowWaits, order)
					System.WaitArea.SlowWaitSize++
				}
			}
			// sort, for next dispatch
			sort.Slice(System.WaitArea.SlowWaits, func(i, j int) bool {
				return System.WaitArea.SlowWaits[i].Number < System.WaitArea.SlowWaits[j].Number
			})
		}
		SystemPrintState()
	} else if System.ChargeArea[Id2Index(charge_id)].State == consts.ChargeFault ||
		System.ChargeArea[Id2Index(charge_id)].State == consts.ChargeOFF {
		output.Printer("System", "charge id="+id+" has been closed")
		return false
	}
	return true
}

func WaitListFindNumberFrontCars(mode string, car_id string) (int64, int64) {
	var front_cars, number int64
	front_cars = 0
	number = 0
	if mode == consts.ChargeModeFast {
		for _, value := range System.WaitArea.FastWaits {
			if car_id == value.CarId {
				number = int64(value.Number)
			} else {
				front_cars++
			}
		}
	} else if mode == consts.ChargeModeSlow {
		for _, value := range System.WaitArea.SlowWaits {
			if car_id == value.CarId {
				number = int64(value.Number)
			} else {
				front_cars++
			}
		}
	}
	return number, front_cars
}

func ChargeFindNumber(state string, mode string, car_id string) int64 {
	if mode == consts.ChargeModeFast {
		for _, id := range System.Control.FastId {
			if state == consts.OrderDispatched && len(System.ChargeArea[Id2Index(id)].Orders) > 1 &&
				System.ChargeArea[Id2Index(id)].Orders[1].CarId == car_id {
				return int64(System.ChargeArea[Id2Index(id)].Orders[1].Number)
			} else if state == consts.OrderCharging && len(System.ChargeArea[Id2Index(id)].Orders) > 0 &&
				System.ChargeArea[Id2Index(id)].Orders[0].CarId == car_id {
				return int64(System.ChargeArea[Id2Index(id)].Orders[0].Number)
			}
		}
	} else if mode == consts.ChargeModeSlow {
		for _, id := range System.Control.SlowId {
			if state == consts.OrderDispatched && len(System.ChargeArea[Id2Index(id)].Orders) > 1 &&
				System.ChargeArea[Id2Index(id)].Orders[1].CarId == car_id {
				return int64(System.ChargeArea[Id2Index(id)].Orders[1].Number)
			} else if state == consts.OrderCharging && len(System.ChargeArea[Id2Index(id)].Orders) > 0 &&
				System.ChargeArea[Id2Index(id)].Orders[0].CarId == car_id {
				return int64(System.ChargeArea[Id2Index(id)].Orders[0].Number)
			}
		}
	}
	return 0
}
