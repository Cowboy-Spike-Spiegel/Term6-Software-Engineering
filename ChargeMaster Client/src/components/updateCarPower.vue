
<template>
  <el-dialog v-model="dialogVisible" width="80%" @close="$emit('closeUpdate')" draggable>

    <el-row >
      <el-col :span="5">
        总电量
      </el-col>
      <el-col :span="10">
        <el-input v-model="power_capacity"/>
      </el-col>
    </el-row>

    <template #footer>
      <span class="dialog-footer">
        <el-button @click="Confirm">确认</el-button>
          <el-button @click="Cancel">取消</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import {computed, ref} from 'vue'

import {ElMessage} from "element-plus";
import CarMsg from "@/class/CarMsg";
import {patchCarMsg} from "@/http/index";
import {json} from "stream/consumers";

//内置函数用于接受父组件参数
const props = defineProps({
  isShow: Boolean,
  info:{
    car_id:"",
    name:"",
    power_capacity: 0,
    power_current:0,
  }
})
//控制显示绑定,监听,设置界面是否显示
const dialogVisible = computed(() => props.isShow)
// const car_id=ref()
const name=ref()
const power_capacity=ref()
const power_current=ref()

//定义向父组件发送的信号关闭信号
// const emits = defineEmits(["closeUpdate"])
const emits = defineEmits(["closeUpdate","updateData"])

//关闭界面信号发送
const Confirm=()=>{
  console.log("info",props.info.car_id)


  patchCarMsg({
    car_id:props.info.car_id,
    name:"",
    power_capacity:parseFloat(power_capacity.value),
    power_current:""

  }).then(res=>{

      if ( res.data["code"] ==0){
        ElMessage.success("编辑成功")
        emits("updateData")
      }
      else {
        ElMessage.error("编辑失败")
    }


    // car_id.value=""
    name.value=""
    power_capacity.value=""
    power_current.value=""
    emits('closeUpdate')
  })

}
const Cancel=()=>{

  ElMessage.info("取消编辑")
  name.value=""
  power_capacity.value=""
  power_current.value=""
  emits('closeUpdate')
}

</script>

<style scoped>

</style>