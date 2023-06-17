package dao

// User - 用户
type SchemaUser struct {
	// user id
	Id       string
	Auth     int
	Account  string
	Password string
	Name     string
	CarArray string
	Balance  float32
}

// Car - 车辆信息 - 相比接口设计数据模型多了mode
type SchemaCar struct {
	Id            string
	UserId        string
	Name          string
	PowerCapacity float32
	PowerCurrent  float32
	State         string
}

// Charge - 充电桩
type SchemaCharge struct {
	Id                  string
	Name                string
	Mode                string
	State               string
	TotalCharge         float32
	TotalChargeTimes    float32
	TotalChargeDuration float32
	TotalChargeFee      float32
	TotalChargeService  float32
	TotalFee            float32
}

// Order - 订单
type SchemaOrder struct {
	Id           string
	UserId       string
	CarId        string
	Mode         string
	ApplyKwh     float32
	ChargeKwh    float32
	State        string
	ChargePrice  float32
	ServicePrice float32
	Fee          float32
	CreateTime   string
	DispatchTime string
	ChargeId     string
	StartTime    string
	FinishTime   string
}

func InitAttribute() {
	AttributeUser["Id"] = "id"
	AttributeUser["Auth"] = "auth"
	AttributeUser["Account"] = "account"
	AttributeUser["Password"] = "password"
	AttributeUser["Name"] = "name"
	AttributeUser["CarArray"] = "cars"
	AttributeUser["Balance"] = "balance"

	AttributeCar["Id"] = "id"
	AttributeCar["UserId"] = "user_id"
	AttributeCar["Name"] = "name"
	AttributeCar["PowerCapacity"] = "power_capacity"
	AttributeCar["PowerCurrent"] = "power_current"
	AttributeCar["State"] = "state"

	AttributeCharge["Id"] = "id"
	AttributeCharge["Name"] = "name"
	AttributeCharge["Mode"] = "mode"
	AttributeCharge["State"] = "state"
	AttributeCharge["TotalCharge"] = "total_charge"
	AttributeCharge["TotalChargeTimes"] = "total_charge_times"
	AttributeCharge["TotalChargeDuration"] = "total_charge_duration"
	AttributeCharge["TotalChargeFee"] = "total_charge_fee"
	AttributeCharge["TotalChargeService"] = "total_charge_service"
	AttributeCharge["TotalFee "] = "total_fee"

	AttributeOrder["Id"] = "id"
	AttributeOrder["CarId"] = "car_id"
	AttributeOrder["UserId"] = "user_id"
	AttributeOrder["Mode"] = "mode"
	AttributeOrder["ApplyKwh"] = "apply_kwh"
	AttributeOrder["ChargeKwh"] = "charge_kwh"
	AttributeOrder["State"] = "state"
	AttributeOrder["ChargePrice"] = "charge_price"
	AttributeOrder["ServicePrice"] = "service_price"
	AttributeOrder["Fee"] = "fee"
	AttributeOrder["CreateTime"] = "create_time"
	AttributeOrder["DispatchTime"] = "dispatch_time"
	AttributeOrder["ChargeId"] = "charge_id"
	AttributeOrder["StartTime"] = "start_time"
	AttributeOrder["FinishTime"] = "finish_time"
}
