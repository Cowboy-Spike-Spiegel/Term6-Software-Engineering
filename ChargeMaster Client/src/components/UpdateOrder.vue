
<template>
    <el-dialog v-model="dialogVisible" width="80%" @close="$emit('updateCloseDetail')" draggable>
        <el-row >
            <el-col :span="5">
                申请充电度数
            </el-col>
            <el-col :span="10">
                <el-input v-model="apply_kwh"/>
            </el-col>
        </el-row>

        <el-form-item label="充电模式">
            <el-radio-group v-model="mode">
                <el-radio label="F">快充</el-radio>
                <el-radio label="T">慢充</el-radio>
            </el-radio-group>
        </el-form-item>

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
import UserOrder from "@/class/UserOrder";
import {patchOrder} from "@/http/index";


//内置函数用于接受父组件参数
const props = defineProps({
    isupdate: Boolean,
    info:UserOrder
})
//控制显示绑定,监听,设置界面是否显示
const dialogVisible = computed(() => props.isupdate)
const id=ref()
const apply_kwh=ref()
const mode=ref()


//定义向父组件发送的信号关闭信号
const emits = defineEmits(["updateCloseDetail","updateData"])

//关闭界面信号发送
const Confirm=()=>{

    if (isNaN(apply_kwh.value)||!apply_kwh.value){
      apply_kwh.value=props.info?.apply_kwh
    }
    if (mode.value!="T"||mode.value!="F"){
      mode.value=props.info?.mode
    }
    console.log("mode",mode.value)


  patchOrder({
    id:props.info?.id,
    mode:mode.value,
    apply_kwh:parseFloat(apply_kwh.value)
  }).then(res=>{
      if (res.data["code"]==0){
        ElMessage.success("编辑成功")
        emits("updateData",mode.value,apply_kwh.value)
      }
      else {
        ElMessage.error("编辑失败")
        apply_kwh.value=""
        mode.value=""
      }
    emits("updateCloseDetail")
  })


}
const Cancel=()=>{
    id.value=""
    apply_kwh.value=""
    mode.value=""
    emits("updateCloseDetail")
}

</script>

<style scoped>

</style>