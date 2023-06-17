package model

type ObjectAdminLoginPost struct {
	Account  string `json:"account"`  // 账户名，注册的用户名（邮箱格式）; 参考正则表达式：^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$
	Password string `json:"password"` // 密码，限制长度32; 参考正则式：^.{1,32}$
}

type ObjectAdminIndexManageGet struct {
}

type ObjectAdminIndexManagePost struct {
	Mode        string   `json:"mode"`         // 开关模式，SWITCHON - 打开; SWITCHOFF - 关闭
	SwitchArray []string `json:"switch_array"` // 开关的充电站数组
}

type ObjectAdminIndexChargeGet struct {
	ChargeID string `json:"charge_id"` // 充电桩id
}

type ObjectAdminIndexReportGet struct {
	ChargeID string `json:"charge_id"` // 需要查询的充电桩id
}

type ObjectClientLoginPost struct {
	Account  string `json:"account"`  // 账户名，注册的用户名（邮箱格式）; 参考正则表达式：^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$
	Password string `json:"password"` // 密码，限制长度32; 参考正则式：^.{1,32}$
}

type ObjectClientRegisterPost struct {
	Account  string `json:"account"`  // 账户名，注册的用户名（邮箱格式）; 参考正则表达式：^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$
	Name     string `json:"name"`     // 用户名，用户昵称
	Password string `json:"password"` // 密码，^.{1,32}$
}

type ObjectClientIndexInformationGet struct {
}

type ObjectClientIndexInformationPost struct {
	Name        string `json:"name,omitempty"`         // 要更新的名称
	OldPassword string `json:"old_password,omitempty"` // 旧密码，^.{1,32}$
	NewPassword string `json:"new_password,omitempty"` // 新密码，^.{1,32}$
}

type ObjectClientIndexOrderGet struct {
	Mode string `json:"mode"` // 订单类型，CURRENT - 当前订单; HISTORY - 历史订单
}

type ObjectClientIndexOrderPost struct {
	ApplyKwh float64 `json:"apply_kwh"` // 申请度数
	CarID    string  `json:"car_id"`    // 车辆ID
	Mode     string  `json:"mode"`      // 充电模式
}

type ObjectClientIndexOrderPatch struct {
	ApplyKwh float64 `json:"apply_kwh"` // 申请度数
	ID       string  `json:"id"`        // 订单ID
	Mode     string  `json:"mode"`      // 充电模式，F快充，T慢充
}

type ObjectClientIndexOrderDelete struct {
	ID string `json:"id"` // 订单ID
}

type ObjectClientIndexCarGet struct {
}

type ObjectClientIndexCarPost struct {
	Name          string  `json:"name"`           // 车辆名称
	PowerCapacity float64 `json:"power_capacity"` // 总电量
	PowerCurrent  float64 `json:"power_current"`  // 当前电量
}

type ObjectClientIndexCarPatch struct {
	CarID         string  `json:"car_id"`         // 车辆ID
	Name          string  `json:"name"`           // 车辆名称
	PowerCapacity float64 `json:"power_capacity"` // 总电量
	PowerCurrent  float64 `json:"power_current"`  // 当前电量
}

type ObjectClientIndexCarDelete struct {
	CarID string `json:"car_id"` // 车辆ID
}

type ObjectClientIndexRechargePost struct {
	Amount string `json:"amount"` // 充值量，单位为RMB
}
