<template>

  <el-descriptions title="用户信息" direction="horizontal" column=2 border>
    <el-descriptions-item label="用户ID" width="10">{{ info.id }}</el-descriptions-item>
    <el-descriptions-item label="账户名"  width="80">{{ info.account }}</el-descriptions-item>
    <el-descriptions-item label="用户余额" width="80">{{ info.balance }}</el-descriptions-item>
    <el-descriptions-item label="用户名" width="80">{{ info.name}}</el-descriptions-item>
  </el-descriptions>
  <el-form-item style="margin-top: 20px">
    <el-button type="primary" @click="editName">编辑用户名称</el-button>
    <el-button type="primary" @click="editPwd" style="margin-left: 40px">修改密码</el-button>
  </el-form-item>


  <el-descriptions title="用户车辆信息" direction="horizontal" column=3 border v-for="(item,index) in info.car_array"
  :key="index" :index="index.toString()" style="margin-top: 20px">
    <el-descriptions-item label="车辆ID" width="80">{{ item["car_id"]}}</el-descriptions-item>
    <el-descriptions-item label="用户ID" width="80">{{item["user_id"] }}</el-descriptions-item>
    <el-descriptions-item label="车辆名称" width="80">{{item["name"] }}</el-descriptions-item>
    <el-descriptions-item label="总电量" width="80">{{ item["power_capacity"] }}</el-descriptions-item>
    <el-descriptions-item label="当前电量" width="80">{{ item["power_current"] }}</el-descriptions-item>
  </el-descriptions>
  <EditUserPassword :is-show="pwdShow" @update-pwd="updatePwd" @close-pwd="closePwd"/>
  <EditUserName :is-show="nameShow" @close-name="closeName" @update-name="updateName" />

</template>

<script lang="ts" setup>
import {ref} from "@vue/reactivity";
import UserMsg from "@/class/UserMsg";
import {getUserMsg} from "@/http/index";
import {onMounted} from "vue";
import EditUserPassword from "@/components/EditUserPassword.vue";
import EditUserName from "@/components/EditUserName.vue";


const info = ref<UserMsg>(new UserMsg())

//通过get获取订单信息
const load = async () => {

    info.value = (await getUserMsg()).data["data"]
    console.log("userMsg",info.value)
}



onMounted(async () => {
    await load()
})

const pwdShow = ref(false)
const nameShow=ref(false)
const closePwd = () => {
  pwdShow.value = false
}
const closeName = () => {
  nameShow.value = false
}

const updatePwd=(password:string)=>{
  info.value.password=password
}
const updateName=(name:string)=>{
  info.value.name=name
}

const editPwd=()=>{
  pwdShow.value = true
}
const editName=()=>{
  nameShow.value = true
}

</script>

<style lang="scss" scoped>


</style>