<template>
  <div>
    <el-form :model="input" label-width="120px">

      <el-form-item label="车辆名称">
        <el-input v-model="input.name"/>
      </el-form-item>
      <el-form-item label="总电量">
        <el-input v-model="input.power_capacity"/>
      </el-form-item>
      <el-form-item label="当前电量">
        <el-input v-model="input.power_current"/>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="addClick">确认新增</el-button>
        <el-button type="primary" @click="cancelClick">取消</el-button>
      </el-form-item>
    </el-form>
  </div>
  <!--    <el-button type="primary" @click="click"  plain>Primary</el-button>-->
</template>

<script lang="ts" setup>
import {ref} from 'vue'
import {ElMessage} from 'element-plus'
import {postCarMsg} from "@/http/index";

let data = {
  name: "",
  power_capacity: "",
  power_current: "",
}
const input = ref(data)
const addClick = () => {

  postCarMsg({
    power_current: parseFloat(input.value.power_current),
    power_capacity: parseFloat(input.value.power_capacity),
    name: input.value.name,
  }).then(res => {
    console.log("res", res.data)
    if (res.data["code"] == 0) {
      ElMessage.success("添加成功")
    } else {
      ElMessage.success("添加失败")
    }
    input.value.name = ""
    input.value.power_capacity = ""
    input.value.power_current = ""
  })


}
const cancelClick = () => {
  ElMessage.success("取消新增")
  // // 执行后端api
  //  ...
  // // 清空操作,用户可能连续添加多个数据


  input.value.name = ""
  input.value.power_capacity = ""
  input.value.power_current = ""
}

</script>