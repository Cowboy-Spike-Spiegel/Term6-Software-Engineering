<template>
  <div class="container" style="width: auto">
    <el-row>
      <el-col :span="12">
        <el-input v-model="SearchVal" placeholder="Please input OrderId" class="input-with-select"
                  @keyup.enter="enterSearch">
          <template #append>
            <el-button :icon="Search" @click="enterSearch"></el-button>
          </template>
        </el-input>
      </el-col>
    </el-row>

    <el-table :data="tableData" style="width: 100%">
      <el-table-column prop="id" label="Id" width="auto"/>
      <el-table-column prop="car_id" label="CarId" width="auto"/>
      <el-table-column prop="apply_kwh" label="Apply_kwh" width="auto"/>
      <el-table-column prop="mode" label="Mode" width="auto"/>
      <el-table-column label="Operations" width="auto">
        <template #default="scope">
          <el-button size="small"  type="default" @click="showDetail(scope.row)">Detail</el-button>
          <el-button size="small" type="primary" @click="editOrder(scope.row)">Edit</el-button>
          <el-button size="small"  type="danger" @click="deleteOrder(scope.$index,scope.row)">Delete</el-button>

        </template>
      </el-table-column>
    </el-table>
  </div>
  <UpdateOrder :isupdate="updateDetailShow"   :info="info" @updateData="updateData" @updateCloseDetail="updateCloseDetail"/>
  <DeleteCarMsg :is-show="deleteShow" :info="info"  :index="deleteIndex" @close-delete="closeDelete" @confirm-delete="confirmDelete"/>
  <Detail :is-show="detailShow" :info="info" @close-detail="closeDetail"></Detail>
</template>

<script lang="ts" setup>
import {Search} from "@element-plus/icons-vue";
import {ref} from "@vue/reactivity";
import UserOrder from "@/class/UserOrder";
import {getCurrentOrder} from "@/http/index";
import {onMounted} from "vue";
import UpdateOrder from "@/components/UpdateOrder.vue";
import DeleteCarMsg from "@/components/DeleteCurrentOrder.vue";
import Detail from "@/components/OrderDetail.vue";


let dataAll = []
const tableData = ref<any[]>([])


//通过get获取订单信息
const load = async () => {

  dataAll = (await getCurrentOrder()).data["data"]

  console.log("dataAll",dataAll)
  tableData.value = dataAll
  console.log("tableData",tableData.value)

}

const SearchVal = ref()
const updateDetailShow = ref(false)
const deleteShow=ref(false)
const detailShow=ref(false)
const deleteIndex=ref(0)

onMounted(async () => {
  await load()
})
const enterSearch = () => {
  let data = []

  if (!SearchVal.value) {
    tableData.value = dataAll

    return
  }
  for (let value of dataAll) {
    if (value["car_id"].includes(SearchVal.value)) {
      data.push(value)
    }
  }
  tableData.value = data
  console.log("page", tableData.value.length)

}
//子组件向父组件发送信号的执行相应函数
const updateCloseDetail = () => {
  updateDetailShow.value = false
}

const info = ref<UserOrder>(new UserOrder())
//


const showDetail = (row: UserOrder) => {
  detailShow.value = true
  info.value = row
}
const editOrder = (row: UserOrder) => {
  updateDetailShow.value = true
  info.value = row
}


const deleteOrder = (index:number, row: UserOrder) => {
  deleteShow.value = true
  deleteIndex.value=index
  info.value=row
}

const confirmDelete=(index:number)=>{
  console.log("deleteIndex",index)
  tableData.value = tableData.value.slice(0, index).concat(tableData.value.slice(index + 1));

}
const closeDelete = () => {
  deleteShow.value = false

}

const closeDetail = () => {
  detailShow.value = false
}


const updateData=(mode:string,apply_kwh:number)=>{
  info.value.mode=mode
  info.value.apply_kwh=apply_kwh
}
</script>

<style lang="scss" scoped>
.container {
  margin: 100px auto;
  width: 50%;
}

.table {
  margin: 10px 0;
}

.input-with-select {
  width: 100%;
  margin-right: 5px;
}


</style>