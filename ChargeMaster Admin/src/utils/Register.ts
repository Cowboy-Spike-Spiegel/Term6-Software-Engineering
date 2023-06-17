import {ref} from "vue";

export interface RegisterUser {
    username: string;
    email: string;
    password: string;
    password2: string;
}

interface RegisterRules {
    username: ({
        required: boolean;
        message: string;
        trigger: string;
    } | {
        required: boolean;
        message: string;
        min: number;
        trigger: string;
    })[];
    email: {
        required: boolean;
        message: string;
        type: string;
        trigger: string;
    }[];
    password: ({
        required: boolean;
        message: string;
        trigger: string;
        min?: undefined;
    } | {
        required: boolean;
        message: string;
        min: number;
        trigger: string;
    })[];
    password2: ({
        required: boolean;
        message: string;
        trigger: string;
    } | {
        required: boolean;
        message: string;
        min: number;
        trigger: string;
    } | {
        validator: (rule: RegisterRules, value: string, callback: any) => void;
        trigger: string;
    })[];
}

export const registerUser = ref<RegisterUser>({
    username: "",
    email: "",
    password: "",
    password2: "",
});


const validatePass2 = (rule: any, value: any, callback: any) => {
    if (value === '') {
        callback(new Error('请再次输入密码'))
    } else if (value !== registerUser.value.password) {
        callback(new Error("两次输入密码不一致"))
    } else {
        callback()
    }
}

export const registerRules = ref<RegisterRules>({
    username: [
        {required: true, message: "请输入用户名", trigger: 'blur'},
        {required: true, message: "用户名长度应大于五位", min: 5, trigger: 'blur'},
    ],
    email: [
        {required: true, type: "email", message: "请输入正确的邮箱", trigger: 'blur'},
    ],
    password: [
        {required: true, message: "请输入密码", trigger: 'blur'},
        {required: true, message: "密码至少为6位", min: 6, trigger: 'blur'},
    ],
    password2: [
        {required: true, message: "请输入密码", trigger: 'blur'},
        {required: true, message: "密码至少为6位", min: 6, trigger: 'blur'},
        {validator: validatePass2, trigger: 'blur'},
    ],
})