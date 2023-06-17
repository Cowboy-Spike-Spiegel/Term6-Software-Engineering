
<template>
    <el-dialog v-model="dialogVisible" width="80%" @close="$emit('closePwd')" draggable>
      <el-row >
        <el-col :span="5">
          请输入原密码
        </el-col>
        <el-col :span="10">
          <el-input v-model="inputOldPwd"/>
        </el-col>
      </el-row>

        <el-row >
            <el-col :span="5">
                修改用户密码
            </el-col>
            <el-col :span="10">
                <el-input v-model="inputPwd"/>
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
import UserMsg from "@/class/UserMsg";
import {postUserMsg} from "@/http/index";


//内置函数用于接受父组件参数
const props = defineProps({
    isShow: Boolean,
})
//控制显示绑定,监听,设置界面是否显示
const dialogVisible = computed(() => props.isShow)
const inputPwd=ref()
const inputOldPwd=ref()

//定义向父组件发送的信号关闭信号
const emits = defineEmits(["closePwd","updatePwd"])

let data={
  old_password:"",
  new_password:"",
  name:""
}

//关闭界面信号发送
const Confirm=()=>{

  let response={}
  postUserMsg({ old_password:inputOldPwd.value,
    new_password:inputPwd.value,
    name:""}).then(res=>{
      response=res.data
      console.log("res",response)
    if(response["code"]==0){

      emits("updatePwd",inputPwd.value)
      ElMessage.success("编辑成功")
    }
    else {
      ElMessage.error("编辑失败")
    }

    inputPwd.value=""
    emits("closePwd")
  })


}
const Cancel=()=>{
  ElMessage.info("取消编辑")
  inputPwd.value=""
  emits("closePwd")
}

</script>

<style scoped>

</style>