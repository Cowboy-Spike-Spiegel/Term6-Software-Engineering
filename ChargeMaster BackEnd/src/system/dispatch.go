package system

import (
	"demo/src/consts"
	"demo/src/dao"
	"demo/src/utils"
	"sort"
	"time"
)

func WaitFastDispatch() {
	if System.WaitArea.FastWaitSize <= 0 {
		//output.Printer("Service", "Nothing for fast dispatch")
		return
	}
	// no space for fast, stay in wait area
	if System.Control.FastRemain <= 0 {
		//output.Printer("Service", "No space for fast dispatch")
		return
	}

	// judge dispatch size
	var dispatch_size int
	if System.WaitArea.FastWaitSize <= System.Control.FastRemain {
		dispatch_size = System.WaitArea.FastWaitSize
	} else {
		dispatch_size = System.Control.FastRemain
	}

	// 1. remove all dispatch_size blocks into tmp memory
	var tmp_orders []Order
	for i := 0; i < dispatch_size; i++ {
		order := System.WaitArea.FastWaits[i]
		order.DispatchTime = time.Now().Format("2006-01-02 15:04:05")
		order.State = consts.OrderDispatched
		tmp_orders = append(tmp_orders, order)
	}
	System.WaitArea.FastWaits = System.WaitArea.FastWaits[dispatch_size:]

	// 2. sort tmp_orders
	sort.Slice(tmp_orders, func(i, j int) bool {
		return tmp_orders[i].ApplyKwh < tmp_orders[j].ApplyKwh
	})

	// 3. find all direct charges
	var rest_ids []int
	for _, id := range System.Control.FastId {
		if System.ChargeArea[id-1].State == consts.ChargeRest {
			rest_ids = append(rest_ids, id)
		}
	}

	// 4. direct dispatch len(rest_ids) blocks
	size := 0
	for ; size < len(rest_ids) && size < dispatch_size; size++ {
		// generate block
		tmp_orders[size].DispatchTime = time.Now().Format("2006-01-02 15:04:05")
		tmp_orders[size].StartTime = time.Now().Format("2006-01-02 15:04:05")
		tmp_orders[size].State = consts.OrderCharging
		tmp_orders[size].ChargeId = System.ChargeArea[Id2Index(rest_ids[size])].Id
		tmp_orders[size].WaitTime = utils.TimeDuration(tmp_orders[size].CreateTime, tmp_orders[size].StartTime)

		// update parameters
		System.WaitArea.FastWaitSize--
		System.ChargeArea[Id2Index(rest_ids[size])].Orders = append(System.ChargeArea[Id2Index(rest_ids[size])].Orders, tmp_orders[size])
		//System.ChargeArea[Id2Index(rest_ids[size])].Count = tmp_orders[size].Order.ApplyKwh
		System.ChargeArea[Id2Index(rest_ids[size])].Size++
		System.Control.FastRemain--

		// update order in db
		dao.OrdersUpdateString(tmp_orders[size].Id, dao.AttributeOrder["State"], tmp_orders[size].State)
		dao.OrdersUpdateString(tmp_orders[size].Id, dao.AttributeOrder["DispatchTime"], tmp_orders[size].DispatchTime)
		dao.OrdersUpdateString(tmp_orders[size].Id, dao.AttributeOrder["StartTime"], tmp_orders[size].StartTime)
		dao.OrdersUpdateInt(tmp_orders[size].Id, dao.AttributeOrder["ChargeId"], tmp_orders[size].ChargeId)
	}
	tmp_orders = tmp_orders[size:]
	dispatch_size -= size
	//fmt.Println("\tAfter direct dispatch", System.WaitArea)

	// 5. if dispatch_size have rest, should dispatch with waiting
	if dispatch_size > 0 {
		type Unit struct {
			id    int     // id of this charge
			count float32 // the rest of this charge count
		}
		var units []Unit
		for _, id := range System.Control.FastId {
			if System.ChargeArea[Id2Index(id)].Size < System.ChargeArea[Id2Index(id)].Capacity {
				var unit Unit
				unit.id = id
				unit.count = System.ChargeArea[Id2Index(id)].Count
				units = append(units, unit)
			}
		}
		// sort to insert the less wait ones
		sort.Slice(units, func(i, j int) bool {
			return units[i].count < units[j].count
		})
		// dispatch the rest into less wait units
		for index := 0; index < dispatch_size; index++ {
			tmp_orders[index].DispatchTime = time.Now().Format("2006-01-02 15:04:05")
			tmp_orders[index].State = consts.OrderDispatched
			tmp_orders[index].ChargeId = units[index].id
			tmp_orders[index].WaitTime = utils.TimeDuration(tmp_orders[index].CreateTime, tmp_orders[index].StartTime)

			System.WaitArea.FastWaitSize--
			System.ChargeArea[Id2Index(units[index].id)].Orders = append(System.ChargeArea[Id2Index(units[index].id)].Orders, tmp_orders[index])
			// System.ChargeArea[Id2Index(rest_ids[index])].Count = tmp_orders[index].Order.SupposeTime
			System.ChargeArea[Id2Index(units[index].id)].Size++
			System.Control.FastRemain--

			// update order in db
			dao.OrdersUpdateString(tmp_orders[index].Id, dao.AttributeOrder["State"], tmp_orders[index].State)
			dao.OrdersUpdateString(tmp_orders[index].Id, dao.AttributeOrder["DispatchTime"], tmp_orders[index].DispatchTime)
			dao.OrdersUpdateInt(tmp_orders[index].Id, dao.AttributeOrder["ChargeId"], tmp_orders[index].ChargeId)
		}
	}
	//fmt.Println("\tAfter not direct dispatch", System.WaitArea)
}

func WaitSlowDispatch() {
	if System.WaitArea.SlowWaitSize <= 0 {
		//output.Printer("Service", "Nothing for slow dispatch")
		return
	}
	// no space for slow, stay in wait area
	if System.Control.SlowRemain <= 0 {
		//output.Printer("Service", "No space for slow dispatch")
		return
	}

	// judge dispatch size
	var dispatch_size int
	if System.WaitArea.SlowWaitSize <= System.Control.SlowRemain {
		dispatch_size = System.WaitArea.SlowWaitSize
	} else {
		dispatch_size = System.Control.SlowRemain
	}

	// 1. remove all dispatch_size blocks into tmp memory
	var tmp_orders []Order
	for i := 0; i < dispatch_size; i++ {
		order := System.WaitArea.SlowWaits[i]
		order.DispatchTime = time.Now().Format("2006-01-02 15:04:05")
		order.State = consts.OrderDispatched
		tmp_orders = append(tmp_orders, order)
	}
	System.WaitArea.SlowWaits = System.WaitArea.SlowWaits[dispatch_size:]

	// 2. sort tmp_orders
	sort.Slice(tmp_orders, func(i, j int) bool {
		return tmp_orders[i].ApplyKwh < tmp_orders[j].ApplyKwh
	})

	// 3. find all direct charges
	var rest_ids []int
	for _, id := range System.Control.SlowId {
		if System.ChargeArea[id-1].State == consts.ChargeRest {
			rest_ids = append(rest_ids, id)
		}
	}

	// 4. direct dispatch len(rest_ids) blocks
	size := 0
	for ; size < len(rest_ids) && size < dispatch_size; size++ {
		// generate block
		tmp_orders[size].DispatchTime = time.Now().Format("2006-01-02 15:04:05")
		tmp_orders[size].StartTime = time.Now().Format("2006-01-02 15:04:05")
		tmp_orders[size].State = consts.OrderCharging
		tmp_orders[size].ChargeId = System.ChargeArea[Id2Index(rest_ids[size])].Id
		tmp_orders[size].WaitTime = utils.TimeDuration(tmp_orders[size].CreateTime, tmp_orders[size].StartTime)

		// update parameters
		System.WaitArea.SlowWaitSize--
		System.ChargeArea[Id2Index(rest_ids[size])].Orders = append(System.ChargeArea[Id2Index(rest_ids[size])].Orders, tmp_orders[size])
		//System.ChargeArea[Id2Index(rest_ids[size])].Count = tmp_orders[size].Order.ApplyKwh
		System.ChargeArea[Id2Index(rest_ids[size])].Size++
		System.Control.SlowRemain--

		// update order in db
		dao.OrdersUpdateString(tmp_orders[size].Id, dao.AttributeOrder["State"], tmp_orders[size].State)
		dao.OrdersUpdateString(tmp_orders[size].Id, dao.AttributeOrder["DispatchTime"], tmp_orders[size].DispatchTime)
		dao.OrdersUpdateString(tmp_orders[size].Id, dao.AttributeOrder["StartTime"], tmp_orders[size].StartTime)
		dao.OrdersUpdateInt(tmp_orders[size].Id, dao.AttributeOrder["ChargeId"], tmp_orders[size].ChargeId)
	}
	tmp_orders = tmp_orders[size:]
	dispatch_size -= size
	//fmt.Println("\tAfter direct dispatch", System.WaitArea)

	// 5. if dispatch_size have rest, should dispatch with waiting
	if dispatch_size > 0 {
		type Unit struct {
			id    int     // id of this charge
			count float32 // the rest of this charge count
		}
		var units []Unit
		for _, id := range System.Control.SlowId {
			if System.ChargeArea[Id2Index(id)].Size < System.ChargeArea[Id2Index(id)].Capacity {
				var unit Unit
				unit.id = id
				unit.count = System.ChargeArea[Id2Index(id)].Count
				units = append(units, unit)
			}
		}
		// sort to insert the less wait ones
		sort.Slice(units, func(i, j int) bool {
			return units[i].count < units[j].count
		})
		// dispatch the rest into less wait units
		for index := 0; index < dispatch_size; index++ {
			tmp_orders[index].DispatchTime = time.Now().Format("2006-01-02 15:04:05")
			tmp_orders[index].State = consts.OrderDispatched
			tmp_orders[index].ChargeId = units[index].id

			System.WaitArea.SlowWaitSize--
			System.ChargeArea[Id2Index(units[index].id)].Orders = append(System.ChargeArea[Id2Index(units[index].id)].Orders, tmp_orders[index])
			// System.ChargeArea[Id2Index(rest_ids[index])].Count = tmp_orders[index].Order.SupposeTime
			System.ChargeArea[Id2Index(units[index].id)].Size++
			System.Control.SlowRemain--

			// update order in db
			dao.OrdersUpdateString(tmp_orders[index].Id, dao.AttributeOrder["State"], tmp_orders[index].State)
			dao.OrdersUpdateString(tmp_orders[index].Id, dao.AttributeOrder["DispatchTime"], tmp_orders[index].DispatchTime)
			dao.OrdersUpdateInt(tmp_orders[index].Id, dao.AttributeOrder["ChargeId"], tmp_orders[index].ChargeId)
		}
	}
	//fmt.Println("\tAfter not direct dispatch", System.WaitArea)
}
