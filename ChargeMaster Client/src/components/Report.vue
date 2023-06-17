<template>
  <div>
    <el-row :gutter="20">
      <el-col :span="8">
        <div>
          <el-card shadow="hover">
            <div class="userCard">
              <el-avatar :size="80" :icon="UserFilled"/>
              <div class="userInfo">
                <p class="important-font">{{ userMsg.name }}</p>
                <p class="secondary-font">用户</p>
              </div>
            </div>
            <a> 余额：{{ userMsg.balance }}</a>
          </el-card>

        </div>
      </el-col>

      <el-col :span="16">
        <div>
          <el-card shadow="hover">
            <div slot="header">
              <span class="important-font">进行中订单</span>
            </div>
            <div>
              <el-table
                  :data="tableData"
                  stripe
                  border
                  style="padding-top: 10px"
              >
                <el-table-column prop="id" label="订单ID" width="120"></el-table-column>
                <el-table-column prop="car_id" label="车辆ID" width="120"></el-table-column>
                <el-table-column prop="start_time" label="开始时间" width="80"></el-table-column>
                <el-table-column prop="apply_kwh" label="申请度数"></el-table-column>
                <el-table-column prop="state" label="状态"></el-table-column>
              </el-table>
            </div>
          </el-card>
        </div>

      </el-col>
    </el-row>


    <el-row :gutter="20" style="margin-top: 15px">
      <el-col :span="24">
        <el-card style="height:1000px" shadow="hover">
          <div style="height: 600px;" ref="chargeGraph"></div>
        </el-card>

        <el-card style="margin-top: 15px" shadow="hover">
          <Map/>
        </el-card>
      </el-col>


    </el-row>


  </div>
</template>

<script setup lang="ts">
import * as echarts from 'echarts';
import {ref} from "@vue/reactivity";
import {onMounted, onUnmounted} from "vue";
import {UserFilled} from '@element-plus/icons-vue'
import {getCurrentOrder, getUserMsg} from "@/http/index";
import UserMsg from "@/class/UserMsg";
import Map from "@/components/Map.vue";


const chargeGraph = ref()
let currentOrder = []
const userMsg = ref<UserMsg>(new UserMsg())
const tableData = ref<any[]>([])

let cars = []
let carName = []
let carCurrent = []
let carTotal = []
let refreshTimer = null;

const loadData = async () => {

  let data = []

  cars = []
  carName = []
  carCurrent = []
  carTotal = []

  //获取当前订单数据
  currentOrder = (await getCurrentOrder()).data["data"]

  console.log("currentOrder", currentOrder)
  tableData.value=currentOrder

  console.log("data",data)
  //获取用户信息数据


  userMsg.value = (await getUserMsg()).data["data"]

  console.log("userMsg",userMsg.value)

  if (Array.isArray(userMsg.value.car_array)){
    for (let value of userMsg.value.car_array) {

      let carData = {
        car_id: "",
        name: "",
        power_capacity: "",
        power_current: "",
      }
      carData.car_id = value["car_id"]
      carData.name = value["name"]
      carData.power_capacity = value["power_capacity"]
      carData.power_current = value["power_current"]
      cars.push(carData)

    }
  }else{
    let carData = {
      car_id: "",
      name: "",
      power_capacity: "",
      power_current: "",
    }

    carData.car_id = userMsg.value.car_array["car_id"]
    carData.name = userMsg.value.car_array["name"]
    carData.power_capacity = userMsg.value.car_array["power_capacity"]
    carData.power_current = userMsg.value.car_array["power_current"]
    cars.push(carData)


  }



}

onMounted(async () => {

  await loadData()
  initChargeGraph()
  refreshTimer=setInterval(async () => {


    await loadData();
    console.log("shuju",cars,carName,carCurrent,carTotal)
    initChargeGraph();
  }, 5000);

})

onUnmounted(() => {
  // 在组件卸载时清除定时器
  clearInterval(refreshTimer);
});

const initChargeGraph = () => {

  for (let value of cars) {
    carName.push(value.name)
    carTotal.push(value.power_capacity)
    carCurrent.push(value.power_current)
  }


  const myChart = echarts.init(chargeGraph.value);

  let option = {
    title: {
      text: '汽车电量信息'
    },
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'shadow'
      }
    },
    legend: {},
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true
    },
    xAxis: {
      type: 'value',
      boundaryGap: [0, 0.01]
    },
    yAxis: {
      type: 'category',
      data: carName
    },
    series: [
      {
        name: '总电量',
        type: 'bar',
        barWidth: 30,
        data: carTotal,

      },
      {
        name: '当前电量',
        type: 'bar',
        barWidth: 30,
        data: carCurrent
      }
    ]
  };

  myChart.setOption(option);

};


</script>


<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="less" scoped>
.el-card__body {
  padding: 10px;
}

.userCard {
  height: 100%;
  display: flex;
  //border-bottom: 2px solid #DCDFE6;
  border-radius: 2px;
}

.userInfo {
  width: auto;
  padding: 6% 0 0 6%;
}

.important-font {
  font-weight: 900;
  font-size: 25px;
}

.secondary-font {
  color: #909399;
}

.login-info {
  height: 40px;
  text-align: left;
  color: #909399;
}

</style>
