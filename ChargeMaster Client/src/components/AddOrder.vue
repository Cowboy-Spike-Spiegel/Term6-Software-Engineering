<template>
  <el-dialog v-model="dialogVisible" width="80%" @close="$emit('closeAdd')" draggable>
    <el-form :model="input" label-width="120px">

<!--      <el-form-item label="车牌号">-->
<!--        <el-input v-model="input.car_id"/>-->
<!--      </el-form-item>-->
      <el-form-item label="申请度数">
        <el-input v-model="input.apply_kwh"/>
      </el-form-item>
      <el-form-item label="充电模式">
        <el-radio-group v-model="input.mode">
          <el-radio label="F">快充</el-radio>
          <el-radio label="T">慢充</el-radio>
        </el-radio-group>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="check">确认</el-button>
        <el-button type="primary" @click="cancle">取消</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>>

</template>

<script lang="ts" setup>
import {computed, ref} from 'vue'
import {ElMessage} from 'element-plus'
import {postOrder} from "@/http/index";

let data = {
  car_id: "",
  mode: "",
  apply_kwh: "",
}

const props = defineProps({
  isShow: Boolean,
  info:{
    car_id:"",
    name:"",
    power_capacity: 0,
    power_current:0,
  }
})

const input = ref(data)
const dialogVisible = computed(() => props.isShow)
const emits = defineEmits(["closeAdd"])

const check = () => {


  console.log("data", input.value)
  // if(parseFloat(input.value.apply_kwh) <= (props.info.power_capacity-props.info.power_current))
  if(true)
  {
    postOrder({
      car_id: props.info.car_id,
      mode: input.value.mode,
      apply_kwh: parseFloat(input.value.apply_kwh)
    }).then(res => {
      console.log("res",res.data)
      if (res.data["code"] == 0) {
        ElMessage.success("创建订单成功")

      } else {
        ElMessage.success("创建订单失败")
      }
      data = {
        car_id: "",
        mode: "",
        apply_kwh: "",
      }
      input.value = data


    })
  }
  else {
    ElMessage.error("创建订单失败,申请电量不合理")

  }


  emits("closeAdd")

}
const cancle = () => {

  ElMessage.info("取消编辑")

  data = {
    car_id: "",
    mode: "",
    apply_kwh: "",
  }
  input.value = data
  emits("closeAdd")
}

</script>