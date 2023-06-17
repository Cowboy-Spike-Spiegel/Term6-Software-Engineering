export default class UserMsg{
    id:string=""
    account:string=""
    password:string=""
    name:string=""
    car_array:object[]=[
        {
            car_id:"",
            user_id:"",
            name: "",
            power_capacity: 0,
            power_current: 0,
        }
    ]
    balance:number=0
}
