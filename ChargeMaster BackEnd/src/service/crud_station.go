package service

import (
	"demo/src/consts"
	"demo/src/model"
	"demo/src/system"
	"demo/src/utils"
	"sort"
	"strconv"
	"time"
)

func ChargeSwitchON(charge_ids []string) bool {
	rt := true
	for _, id := range charge_ids {
		if system.ChargeOpenById(id) == false {
			rt = false
		}
	}
	return rt
}

func ChargeSwitchOFF(charge_ids []string) bool {
	rt := true
	for _, id := range charge_ids {
		if system.ChargeCloseById(id) == false {
			rt = false
		}
	}
	return rt
}

func StationInformation() model.Station {
	var station model.Station
	station.ID = consts.StationId

	// generate charge area
	for index := 0; index < len(system.System.ChargeArea); index++ {
		obj := system.System.ChargeArea[index]
		// 需要查询充电桩现状
		var charge model.Charge
		charge.ID = strconv.Itoa(obj.Id)
		charge.Name = obj.Name
		charge.Mode = obj.Mode
		charge.State = obj.State
		charge.TotalCharge = float64(obj.TotalCharge)
		charge.TotalChargeTimes = obj.TotalChargeTimes
		charge.TotalChargeDuration = float64(obj.TotalChargeDuration)
		charge.TotalChargeFee = float64(obj.TotalChargeFee)
		charge.TotalChargeService = float64(obj.TotalServiceFee)
		charge.TotalFee = float64(obj.TotalFee)

		// charge.CarBlocks
		for _, order := range obj.Orders {
			var block model.CarBlock
			block.Number = strconv.Itoa(order.Number)
			block.CarID = order.CarId
			block.UserID = order.UserId
			block.Name = order.Name
			block.Mode = order.Mode
			block.PowerCapacity = float64(order.PowerCapacity)
			block.PowerCurrent = float64(order.PowerCurrent)
			block.ApplyKwh = float64(order.ApplyKwh)
			if order.StartTime == consts.TimeDefualt {
				block.WaitTime = float64(utils.TimeDuration(order.CreateTime, time.Now().Format("2006-01-02 15:04:05")))
			} else {
				block.WaitTime = float64(utils.TimeDuration(order.CreateTime, order.StartTime))
			}
			block.State = order.State
			charge.CarBlocks = append(charge.CarBlocks, block)
		}
		station.ChargeArray = append(station.ChargeArray, charge)
	}

	// generate wait
	for _, order := range system.System.WaitArea.FastWaits {
		var block model.CarBlock
		block.Number = strconv.Itoa(order.Number)
		block.CarID = order.CarId
		block.UserID = order.UserId
		block.Name = order.Name
		block.Mode = order.Mode
		block.PowerCapacity = float64(order.PowerCapacity)
		block.PowerCurrent = float64(order.PowerCurrent)
		block.ApplyKwh = float64(order.ApplyKwh)
		block.WaitTime = float64(utils.TimeDuration(order.CreateTime, time.Now().Format("2006-01-02 15:04:05")))
		block.State = order.State
		station.WaitArray = append(station.WaitArray, block)
	}
	for _, order := range system.System.WaitArea.SlowWaits {
		var block model.CarBlock
		block.Number = strconv.Itoa(order.Number)
		block.CarID = order.CarId
		block.UserID = order.UserId
		block.Name = order.Name
		block.PowerCapacity = float64(order.PowerCapacity)
		block.PowerCurrent = float64(order.PowerCurrent)
		block.ApplyKwh = float64(order.ApplyKwh)
		block.WaitTime = float64(utils.TimeDuration(order.CreateTime, time.Now().Format("2006-01-02 15:04:05")))
		block.State = order.State
		station.WaitArray = append(station.WaitArray, block)
	}
	sort.Slice(station.WaitArray, func(i, j int) bool {
		return station.WaitArray[i].Number < station.WaitArray[j].Number
	})
	return station
}

func ChargeInformation(charge_id int) model.ChargeMini {
	var charge model.ChargeMini

	// get charge in system
	obj := system.System.ChargeArea[system.Id2Index(charge_id)]
	charge.ID = strconv.Itoa(obj.Id)
	charge.Name = obj.Name
	charge.Mode = obj.Mode
	charge.State = obj.State

	// charge.CarBlocks
	for _, order := range obj.Orders {
		var block model.CarBlock
		block.Number = strconv.Itoa(order.Number)
		block.CarID = order.CarId
		block.UserID = order.UserId
		block.Name = order.Name
		block.PowerCapacity = float64(order.PowerCapacity)
		block.PowerCurrent = float64(order.PowerCurrent)
		block.ApplyKwh = float64(order.ApplyKwh)
		if order.StartTime == consts.TimeDefualt {
			block.WaitTime = float64(utils.TimeDuration(order.CreateTime, time.Now().Format("2006-01-02 15:04:05")))
		} else {
			block.WaitTime = float64(utils.TimeDuration(order.CreateTime, order.StartTime))
		}
		block.State = order.State
		charge.CarBlocks = append(charge.CarBlocks, block)
	}
	return charge
}

func ChargeReport(charge_id int) model.Report {
	var report model.Report

	obj := system.System.ChargeArea[system.Id2Index(charge_id)]
	report.ID = strconv.Itoa(obj.Id)
	report.Mode = obj.Mode
	report.Name = obj.Name
	report.State = obj.State
	report.Time = time.Now().Format("2006-01-02 15:04:05")
	report.TotalCharge = float64(obj.TotalCharge)
	report.TotalChargeDuration = float64(obj.TotalChargeDuration)
	report.TotalChargeFee = float64(obj.TotalChargeFee)
	report.TotalChargeService = float64(obj.TotalServiceFee)
	report.TotalChargeTimes = obj.TotalChargeTimes
	report.TotalFee = float64(obj.TotalFee)

	return report
}
