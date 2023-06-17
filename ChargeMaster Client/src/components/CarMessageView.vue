<template>

  <div v-for="(item,index) in cars"
       :key="index" :index="index.toString()" style="margin-bottom: 30px">
    <el-descriptions title="车辆信息" direction="horizontal" column=2 border style="margin-bottom: 10px">
      <el-descriptions-item label="车辆ID" width="80">{{ item["car_id"] }}</el-descriptions-item>
      <el-descriptions-item label="车辆名称" width="80">{{ item["name"] }}</el-descriptions-item>
      <el-descriptions-item label="总电量" width="80">{{ item["power_capacity"] }}</el-descriptions-item>
      <el-descriptions-item label="当前电量" width="80">{{ item["power_current"] }}</el-descriptions-item>
    </el-descriptions>
    <el-button type="default" @click="addOrder(item)">创建订单</el-button>
    <el-button type="primary" @click="editCarMsg(item)">编辑车辆名称</el-button>
    <el-button type="primary" @click="editCarCurrent(item)">编辑车辆当前电量</el-button>
    <el-button type="primary" @click="editCarPower(item)">编辑车辆总电量</el-button>
    <el-button type="danger" @click="deleteCar()">删除车辆信息</el-button>

  </div>
  <AddOrder :is-show ="addShow" :info="info"  @close-add="closeAdd"  />
  <UpdateCarName :is-show="nameShow" :info="info" @update-data="updateCar" @close-update="closeUpdate"/>
  <update-car-current :is-show="currentShow" :info="info" @update-data="updateCar" @close-update="closeUpdate"/>
  <update-car-power :is-show="powerShow" :info="info" @update-data="updateCar" @close-update="closeUpdate"/>
  <DeleteCarMsg :is-show="deleteShow" @delete="deleteCarMsg" @close="closeDelete"/>



</template>



<script lang="ts" setup>
import {ref} from "@vue/reactivity";
import {getCarMsg} from "@/http/index";
import {onMounted} from "vue";
import UpdateCarName from "@/components/UpdateCarName.vue";
import UpdateCarCurrent from "@/components/updateCarCurrent.vue";
import UpdateCarPower from "@/components/updateCarPower.vue";
import AddOrder from "@/components/AddOrder.vue";

const cars = ref([])
const info = ref({
  car_id: "",
  name: "",
  power_capacity: 0,
  power_current: 0,
})

//通过get获取订单信息
const load = async () => {

  cars.value = (await getCarMsg()).data["data"]
  console.log("carMsg",cars.value,(await getCarMsg()).data)

}

onMounted(async () => {
  await load()
})
const nameShow = ref(false)
const powerShow = ref(false)
const currentShow = ref(false)
const deleteShow = ref(false)
const addShow = ref(false)
const closeUpdate = () => {

  nameShow.value = false
  currentShow.value = false
  powerShow.value = false
  addShow.value=false
}
const closeDelete = () => {
  deleteShow.value = false
}

const closeAdd = () => {
  addShow.value = false
}

const updateCar = () => {

  load()
}
const deleteCarMsg = (name: string, power_capacity: string, power_current: string) => {

}
const editCarMsg = (row: {
  car_id: "",
  name: "",
  power_capacity: 0,
  power_current: 0,
}) => {
  nameShow.value = true
  info.value = row
}
const editCarCurrent = (row: {
  car_id: "",
  name: "",
  power_capacity: 0,
  power_current: 0,
}) => {
  currentShow.value = true
  info.value = row
}
const editCarPower = (row: {
  car_id: "",
  name: "",
  power_capacity: 0,
  power_current: 0,
}) => {
  powerShow.value = true
  info.value = row
}
const deleteCar = () => {
  deleteShow.value = true
}

const addOrder = (row: {
  car_id: "",
  name: "",
  power_capacity: 0,
  power_current: 0,
}) => {
  console.log("addshow",addShow.value)
  addShow.value = true
  console.log("addshow",addShow.value)
  info.value=row
}

</script>

<style lang="scss" scoped>


</style>

