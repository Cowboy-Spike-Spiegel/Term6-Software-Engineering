export default class UserOrder {
    id:string=""
    user_id:string=""
    car_id:string=""
    mode:string=""
    apply_kwh:number=0
    charge_kwh:number=0
    state:string="FINISHED"
    charge_price:number=0
    service_price:number=0
    fee:number=0
    create_time:string=""
    dispatch_time:string="xx-xx-xx xx:xx:xx"
    charge_id:string=""
    start_time:string="xx-xx-xx xx:xx:xx"
    finish_time:string="xx-xx-xx xx:xx:xx"
    number:number=0
    front_cars :number=0
}
//导出这个类，定义一中输出接口