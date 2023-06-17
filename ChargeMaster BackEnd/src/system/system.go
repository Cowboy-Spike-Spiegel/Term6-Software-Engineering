package system

import (
	"demo/src/consts"
	"demo/src/dao"
	"demo/src/output"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var System ChargeMaster

func Init() {
	SystemInit()
	output.Printer("Service", "ChargeMaster init success")

	// run wait area
	go RunSystemWait(&System.WaitArea)
	for index := 0; index < len(System.ChargeArea); index++ {
		go RunSystemCharge(&System.ChargeArea[index])
	}
}

func SystemInit() {
	// generate WaitList
	System.WaitArea.Capacity = consts.WaitLimit
	System.WaitArea.FastWaitSize = 0
	System.WaitArea.SlowWaitSize = 0
	System.WaitArea.Number = 0

	System.Control.FastCapacity = 0
	System.Control.FastRemain = 0
	System.Control.SlowCapacity = 0
	System.Control.SlowRemain = 0

	// generate []Charge
	charges := dao.ChargesQueryAll()
	for _, value := range charges {

		var charge Charge

		// ChargeArea
		id, _ := strconv.Atoi(value.Id)
		charge.Id = id
		charge.Name = value.Name
		charge.Mode = value.Mode
		charge.State = value.State
		charge.TotalCharge = 0
		charge.TotalChargeTimes = 0
		charge.TotalChargeDuration = 0
		charge.TotalChargeFee = 0
		charge.TotalServiceFee = 0
		charge.TotalFee = 0
		// chargeBlocks don't init, remain empty list
		charge.Capacity = consts.ChargeLimit
		charge.Size = 0
		charge.Count = 0
		System.ChargeArea = append(System.ChargeArea, charge)

		// not fault, add control information
		if value.State != consts.ChargeFault {
			// WaitArea and Mode
			if value.Mode == consts.ChargeModeFast {
				System.Control.FastId = append(System.Control.FastId, charge.Id)
				System.Control.FastCapacity += consts.ChargeLimit
				System.Control.FastRemain += consts.ChargeLimit
			} else if value.Mode == consts.ChargeModeSlow {
				System.Control.SlowId = append(System.Control.SlowId, charge.Id)
				System.Control.SlowCapacity += consts.ChargeLimit
				System.Control.SlowRemain += consts.ChargeLimit
			}

			// update charge state, because not fault convert to rest
			dao.ChargesUpdateById(strconv.Itoa(charge.Id), dao.AttributeCharge["State"], consts.ChargeRest)
		}
	}
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

type ChargeMaster struct {
	ChargeArea []Charge
	WaitArea   WaitList
	Control    ControlObject
}

type Charge struct {
	Mutex sync.Mutex // only Item.State..., Orders can be written outside besides init

	// charge information
	Id                  int
	Name                string
	Mode                string
	State               string
	TotalCharge         float32
	TotalChargeTimes    int
	TotalChargeDuration float32 // seconds
	TotalChargeFee      float32
	TotalServiceFee     float32
	TotalFee            float32

	// charge run parameters
	Capacity int
	Size     int
	Orders   []Order
	Count    float32 // count for remain seconds to end the first order
}

type WaitList struct {
	Mutex sync.Mutex

	FastWaits    []Order
	SlowWaits    []Order
	Capacity     int // wait capacity
	FastWaitSize int // all fast wait size
	SlowWaitSize int // all slow wait size

	Number int // order's number
}

type ControlObject struct {
	Mutex sync.Mutex

	FastId       []int // id for FastCharge
	FastCapacity int   // fast charge capacity sum
	FastRemain   int   // all fast charge remain blocks

	SlowId       []int // id for SlowCharge
	SlowCapacity int   // slow charge capacity sum
	SlowRemain   int   // all slow charge remain blocks
}

type Order struct {
	// 订单ID
	Id           string
	UserId       string
	CarId        string
	Mode         string
	ChargeSpeed  float32
	ApplyKwh     float32
	ChargeKwh    float32
	State        string // 订单状态，WAITING - 在等待区等待 DISPATCHED - 在充电桩内等待 CHARGING - 充电中 FINISHED - 结束
	ChargePrice  float32
	ServicePrice float32
	ChargeFee    float32
	ServiceFee   float32
	Fee          float32
	CreateTime   string
	DispatchTime string
	ChargeId     int // 调度到的充电桩id, 无为0
	StartTime    string
	FinishTime   string

	Number   int
	WaitTime float32

	// 车辆附加信息
	Name          string
	PowerCapacity float32
	PowerCurrent  float32
}

func SystemPrintState() {
	if os.Args[1] == "test" {
		fmt.Println()
		fmt.Println("Control", System.Control)
		fmt.Println("Charge Area")
		for _, charge := range System.ChargeArea {
			fmt.Println("--id =", charge.Id, "size =", charge.Size, "state =", charge.State, "orders =")
			for index, order := range charge.Orders {
				fmt.Println("  ", index, order)
			}
		}
		fmt.Println("Wait Area, FastWaitSize =", System.WaitArea.FastWaitSize, "SlowWaitSize =", System.WaitArea.SlowWaitSize)
		for _, obj := range System.WaitArea.FastWaits {
			fmt.Println("--order: id =", obj.Id, "state =", obj.State, "wait time =", obj.WaitTime)
		}
		for _, obj := range System.WaitArea.SlowWaits {
			fmt.Println("--order: id =", obj.Id, "state =", obj.State, "wait time =", obj.WaitTime)
		}
		fmt.Println()
	}
}
