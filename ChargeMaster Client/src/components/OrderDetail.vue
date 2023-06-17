<template>
  <el-dialog v-model="dialogVisible" width="80%" @close="$emit('closeDetail')" draggable>
    <el-descriptions title="用户订单详情" direction="horizontal" column=4 border>
        <el-descriptions-item label="订单ID">{{ info.id }}</el-descriptions-item>
        <el-descriptions-item label="用户ID">{{ info.user_id }}</el-descriptions-item>
        <el-descriptions-item label="车辆ID">{{ info.car_id }}</el-descriptions-item>
        <el-descriptions-item label="充电模式">{{ info.mode }}</el-descriptions-item>
        <el-descriptions-item label="申请度数">{{ info.apply_kwh }}</el-descriptions-item>
        <el-descriptions-item label="实际充电度数">{{ info.charge_kwh }}</el-descriptions-item>
        <el-descriptions-item label="订单状态">{{ info.state }}</el-descriptions-item>
        <el-descriptions-item label="充电计价">{{ info.charge_price }}</el-descriptions-item>
        <el-descriptions-item label="服务计价">{{ info.service_price }}</el-descriptions-item>
        <el-descriptions-item label="订单最终花费">{{ info.fee }}</el-descriptions-item>
        <el-descriptions-item label="创建时间">{{ info.create_time }}</el-descriptions-item>
        <el-descriptions-item label="调度时间">{{ info.dispatch_time }}</el-descriptions-item>
        <el-descriptions-item label="调度到的充电桩ID">{{ info.charge_id }}</el-descriptions-item>
        <el-descriptions-item label="开始时间">{{ info.start_time }}</el-descriptions-item>
        <el-descriptions-item label="终止时间">{{ info.finish_time }}</el-descriptions-item>
        <el-descriptions-item label="排队号码">{{ info.number }}</el-descriptions-item>
        <el-descriptions-item label="前面车辆">{{ info.front_cars }}</el-descriptions-item>
    </el-descriptions>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="closeDetail">确认</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import {computed, ref} from 'vue'
import UserOrder from "@/class/UserOrder";

//内置函数用于接受父组件参数
const props = defineProps({
  isShow: Boolean,
  info: UserOrder
})
//控制显示绑定,监听,设置界面是否显示
const dialogVisible = computed(() => props.isShow)

//定义向父组件发送的信号关闭信号
const emits = defineEmits(["closeDetail"])


const size = ref('')
const blockMargin = computed(() => {
  const marginMap = {
    large: '32px',
    default: '28px',
    small: '24px',
  }
  return {
    marginTop: marginMap[size.value] || marginMap.default,
  }
})

//关闭界面信号发送
const closeDetail=()=>{
  emits("closeDetail")
}

</script>

<style scoped>

</style>