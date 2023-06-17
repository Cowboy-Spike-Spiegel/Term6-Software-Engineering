package system

import (
	"demo/src/consts"
	"demo/src/dao"
	"demo/src/utils"
	"fmt"
	"time"
)

func RunSystemWait(WaitArea *WaitList) {
	for true {
		time.Sleep(time.Second * consts.WaitAreaPeriod)
		WaitFastDispatch()
		WaitSlowDispatch()
		SystemPrintState()
	}
}

func RunSystemCharge(charge *Charge) {
	var power float32
	if charge.Mode == consts.ChargeModeFast {
		power = consts.ChargeFastKwh * consts.ChargeAreaPeriod / 3600
	} else if charge.Mode == consts.ChargeModeSlow {
		power = consts.ChargeSlowKwh * consts.ChargeAreaPeriod / 3600
	}

	for true {
		time.Sleep(time.Second * consts.ChargeAreaPeriod)
		if charge.State == consts.ChargeWork {
			// calculate fees
			charge_fee := GetPrice() * power
			service_fee := consts.ChargeServicePrice * power

			// update order parameters
			charge.Orders[0].ChargeKwh += power
			charge.Orders[0].ChargeFee += charge_fee
			charge.Orders[0].ServiceFee += service_fee
			charge.Orders[0].Fee += charge_fee + service_fee
			charge.Orders[0].PowerCurrent += power

			// update car parameters
			str := fmt.Sprintf("%f", charge.Orders[0].PowerCurrent)
			dao.CarsUpdateById(charge.Orders[0].CarId, dao.AttributeCar["PowerCurrent"], str)

			// update charge parameters
			charge.TotalCharge += power
			charge.TotalChargeDuration += consts.ChargeAreaPeriod
			charge.TotalChargeFee += charge_fee
			charge.TotalServiceFee += service_fee
			charge.TotalFee += charge_fee + service_fee
			charge.Count -= power

			if charge.Count <= 0 {
				// stop this order
				order := charge.Orders[0]
				charge.Orders = charge.Orders[1:]

				// update charge
				charge.TotalChargeTimes++
				charge.State = consts.ChargeRest
				charge.Size--
				charge.Count = 0

				// update control
				if charge.Mode == consts.ChargeModeFast {
					System.Control.FastRemain++
				} else if charge.Mode == consts.ChargeModeSlow {
					System.Control.SlowRemain++
				}

				// insert this order
				order.FinishTime = time.Now().Format("2006-01-02 15:04:05")
				order.ChargePrice = order.ChargeFee / order.ChargeKwh
				order.State = consts.OrderFinished
				SystemCompleteOrder(order)
			}
		}
		// rest to start new order
		//fmt.Println("\n\t", charge)
		if charge.State == consts.ChargeRest && charge.Size > 0 {
			charge.State = consts.ChargeWork
			charge.Count = charge.Orders[0].ApplyKwh
			charge.Orders[0].State = consts.CarCharging
			charge.Orders[0].StartTime = time.Now().Format("2006-01-02 15:04:05")
			charge.Orders[0].WaitTime = utils.TimeDuration(charge.Orders[0].CreateTime, charge.Orders[0].StartTime)

			// update order
			dao.OrdersUpdateString(charge.Orders[0].Id, dao.AttributeOrder["State"], charge.Orders[0].State)
			dao.OrdersUpdateString(charge.Orders[0].Id, dao.AttributeOrder["StartTime"], charge.Orders[0].StartTime)
		}
	}
}

func Id2Index(id int) int {
	return id - 1
}

func GetNumber() int {
	res := System.WaitArea.Number
	System.WaitArea.Number++
	return res
}
