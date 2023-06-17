import service from "@/http/index/request";


export function register(data:any){
    return service({
        url:'/client/register',
        method:'POST',
        data: data
    })
}

export function login(data:any){
    return service({
        url:'/client/login',
        method:'POST',
        data: data
    })
}

export function getAllMessage(data:any){
    return service({
        url:'/admin/index/manage',
        method:'GET',
        data:data
    })
}