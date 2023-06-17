package consts

var mode string

const ModeTest = "test"
const ModeRelease = "release"

const SUCCESS = 0
const FAIL = -1

// basic values
const NotExistId = "-1"
const RepeatId = "-2"
const InvalidAuth = -1

// output variables
const LayerLength = 10

// register limits
const AccountMinLength = 10
const PasswordMinLength = 8

// switchers
const SwitchON = "SWITCHON"
const SwitchOFF = "SWITCHOFF"

// orders
const SearchHistory = "HISTORY"
const SearchCurrent = "CURRENT"

const OrderWaiting = "WAITING"
const OrderDispatched = "DISPATCHED"
const OrderCharging = "CHARGING"
const OrderFinished = "FINISHED"

// station information
const ChargeAreaPeriod = 1
const WaitAreaPeriod = 10

const StationId = "0"
const StationName = "BUPT充电站"
const ChargeLimit = 2
const WaitLimit = 6
const NoId = 0
const NoIndex = -1

// charge
const ChargeOFF = "OFF"
const ChargeRest = "REST"
const ChargeWork = "WORK"
const ChargeFault = "FAULT"

const ChargeModeFast = "F"
const ChargeFastKwh float32 = 100
const ChargeModeSlow = "T"
const ChargeSlowKwh float32 = 50

const ChargeServicePrice float32 = 0.8

// car state
const CarWaiting = "WAITING"
const CarDispatched = "DISPATCHED"
const CarCharging = "CHARGING"
const CarUsing = "USING"

// Price
const ChargeHighPrice float32 = 1
const ChargeMiddlePrice float32 = 0.7
const ChargeLowPrice float32 = 0.4

// time
const TimeDefualt = "xxxx-xx-xx xx:xx:xx"
