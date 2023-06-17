package model

type ResponseClientIndexInformationGet struct {
	Code int64  `json:"code"` // 状态码
	Data User   `json:"data"`
	Msg  string `json:"msg"` // 返回信息，失败的时候返回报错信息，成功的时候返回成功信息
}

type User struct {
	Account  string  `json:"account"`   // 账户名，注册的用户名（邮箱格式）; 参考正则表达式：^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$
	Balance  float64 `json:"balance"`   // 用户余额
	CarArray []Car   `json:"car_array"` // car_id数组，注册时初始化为大小为0
	ID       string  `json:"id"`        // 用户id
	Name     string  `json:"name"`      // 用户名
}

type Car struct {
	CarID         string  `json:"car_id"` // 车辆ID
	Name          string  `json:"name,omitempty"`
	PowerCapacity float64 `json:"power_capacity"` // 总电量
	PowerCurrent  float64 `json:"power_current"`  // 当前电量
	UserID        string  `json:"user_id"`        // 用户ID
}

type ResponseClientIndexOrderGet struct {
	Code int64   `json:"code"` // 状态码
	Data []Order `json:"data"`
	Msg  string  `json:"msg"` // 返回信息，失败的时候返回报错信息，成功的时候返回成功信息
}

type Order struct {
	ApplyKwh     float64 `json:"apply_kwh"`        // 申请度数，可能不是最后充电的度数！
	CarID        string  `json:"car_id"`           // 车辆ID
	ChargeID     string  `json:"charge_id"`        // 调度到的充电桩id，无为空串
	ChargeKwh    float64 `json:"charge_kwh"`       // 实际充电度数，确切的充电度数，由时间算出
	ChargePrice  float64 `json:"charge_price"`     // 充电计价，平均充电计价
	CreateTime   string  `json:"create_time"`      // 创建时间
	DispatchTime string  `json:"dispatch_time"`    // 调度时间，无为xx-xx-xx xx:xx:xx
	Fee          float64 `json:"fee"`              // 订单最终花费
	FinishTime   string  `json:"finish_time"`      // 终止时间，无为xx-xx-xx xx:xx:xx
	FrontCars    int64   `json:"front_cars"`       // 查看前面有多少车，排号号码-当前已进入充电区号码-1
	ID           string  `json:"id"`               // 订单ID
	Mode         string  `json:"mode"`             // 充电模式，F快充，T慢充
	Number       int64   `json:"number,omitempty"` // 排号号码，类似叫餐服务，订单当日递增
	ServicePrice float64 `json:"service_price"`    // 服务计价，服务计价
	StartTime    string  `json:"start_time"`       // 开始时间，无为xx-xx-xx xx:xx:xx
	State        string  `json:"state"`            // 订单状态，WAITING - 在等待区等待; DISPATCHED - 在充电桩内等待; CHARGING - 充电中; FINISHED - 结束
	UserID       string  `json:"user_id"`          // 用户ID
}

type ResponseClientIndexCarGet struct {
	Code int64  `json:"code"` // 状态码
	Data []Car  `json:"data"`
	Msg  string `json:"msg"`
}

type ResponseAdminIndexManageGet struct {
	Code int64   `json:"code"` // 返回状态码，可自定义
	Data Station `json:"data"` // 这里面的数据一定是系统启动后的
	Msg  string  `json:"msg"`
}

type Station struct {
	ChargeArray []Charge   `json:"charge_array"` // 充电桩数组，充电站下的充电桩
	ID          string     `json:"id"`           // id标识
	WaitArray   []CarBlock `json:"wait_array"`   // 等待车位数组，内容是车位
}

// 充电桩
type Charge struct {
	CarBlocks           []CarBlock `json:"car_blocks"`            // 充电车位数组，空格隔开，大小为2，第一个位置和第二个位置
	ID                  string     `json:"id"`                    // 充电桩id，唯一分配
	Mode                string     `json:"mode"`                  // 充电模式，F- 快充 T- 慢充
	Name                string     `json:"name"`                  // 名字
	State               string     `json:"state"`                 // 工作状态，OFF - 关机状态; WORK - 充电状态; REST - 空闲状态; FAULT - 故障状态
	TotalCharge         float64    `json:"total_charge"`          // 总充电总量
	TotalChargeDuration float64    `json:"total_charge_duration"` // 总充电时长，单位为小时
	TotalChargeFee      float64    `json:"total_charge_fee"`      // 总充电费用，单位为RMB
	TotalChargeService  float64    `json:"total_charge_service"`  // 总服务费用，单位为RMB
	TotalChargeTimes    int        `json:"total_charge_times"`    // 总充电次数
	TotalFee            float64    `json:"total_fee"`             // 总费用，单位为RMB
}

// 车位信息
type CarBlock struct {
	ApplyKwh      float64 `json:"apply_kwh"` // 申请的充电量
	CarID         string  `json:"car_id"`    // 车辆ID
	Name          string  `json:"name,omitempty"`
	Number        string  `json:"number"`         // 此车位的车辆分配的number
	PowerCapacity float64 `json:"power_capacity"` // 总电量
	PowerCurrent  float64 `json:"power_current"`  // 当前电量
	State         string  `json:"state"`          // 车位内的车辆状态，WAITING - 在等待区等待DISPATCHED - 在充电桩内等待CHARGING - 充电中
	UserID        string  `json:"user_id"`        // 用户ID
	WaitTime      float64 `json:"wait_time"`      // 等待时间，单位为小时
	Mode          string  `json:"mode"`           // mode类型
}

type ResponseAdminIndexChargeGet struct {
	Code int64      `json:"code"` // 状态码
	Data ChargeMini `json:"data"`
	Msg  string     `json:"msg"`
}

type ChargeMini struct {
	CarBlocks []CarBlock `json:"car_blocks"` // 充电车位数组，空格隔开，大小为2，第一个位置和第二个位置
	ID        string     `json:"id"`         // 充电桩id，唯一分配
	Mode      string     `json:"mode"`       // 充电模式，F- 快充 T- 慢充
	Name      string     `json:"name"`       // 名字
	State     string     `json:"state"`      // 工作状态，OFF - 关机状态; WORK - 充电状态; REST - 空闲状态; FAULT - 故障状态
}

type ResponseAdminIndexReportGet struct {
	Code int64  `json:"code"` // 状态码
	Data Report `json:"data"`
	Msg  string `json:"msg"` // 返回信息，失败的时候返回报错信息，成功的时候返回成功信息
}

type Report struct {
	ID                  string  `json:"id"`                    // 充电桩id，唯一分配
	Mode                string  `json:"mode"`                  // 充电模式，F- 快充 T- 慢充
	Name                string  `json:"name"`                  // 名字
	State               string  `json:"state"`                 // 工作状态，OFF - 关机状态; WORK - 充电状态; REST - 空闲状态; FAULT - 故障状态
	Time                string  `json:"time"`                  // 当前时间
	TotalCharge         float64 `json:"total_charge"`          // 总充电总量
	TotalChargeDuration float64 `json:"total_charge_duration"` // 总充电时长，单位为小时
	TotalChargeFee      float64 `json:"total_charge_fee"`      // 总充电费用，单位为RMB
	TotalChargeService  float64 `json:"total_charge_service"`  // 总服务费用，单位为RMB
	TotalChargeTimes    int     `json:"total_charge_times"`    // 总充电次数
	TotalFee            float64 `json:"total_fee"`             // 总费用，单位为RMB
}
