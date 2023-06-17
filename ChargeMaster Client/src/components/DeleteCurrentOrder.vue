
<template>
    <el-dialog v-model="dialogVisible" width="300px" @close="$emit('closeDelete')" draggable>


        <template #footer>
      <span class="dialog-footer">
        <el-button type="danger" @click="Confirm">确认删除</el-button>
          <el-button type="primary" @click="Cancel">取消</el-button>
      </span>
        </template>
    </el-dialog>
</template>

<script setup lang="ts">
import {computed, ref} from 'vue'

import {ElMessage} from "element-plus";
import CarMsg from "@/class/CarMsg";
import {deleteOrder} from "@/http/index";
import UserOrder from "@/class/UserOrder";

//内置函数用于接受父组件参数
const props = defineProps({
    isShow: Boolean,
    info:UserOrder,
    index:Number
})
//控制显示绑定,监听,设置界面是否显示
const dialogVisible = computed(() => props.isShow)

//定义向父组件发送的信号关闭信号
const emits = defineEmits(["closeDelete","confirmDelete"])

//关闭界面信号发送
const Confirm=()=>{

   deleteOrder({id:props.info?.id}).then(res=>{
     console.log("deleteData",res.data)
     if (res.data["code"]==0){
       ElMessage.success("删除成功")
       console.log("delete",props.index)
       emits("confirmDelete",props.index)
     }
     else{
       ElMessage.error("删除失败")
     }
     emits('closeDelete')
   })

}
const Cancel=()=>{
  ElMessage.info("取消删除")
    emits('closeDelete')
}

</script>

<style scoped>

</style>