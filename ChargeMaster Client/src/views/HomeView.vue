<template>
  <el-container style="position: absolute;left: 0;right: 0;top: 0;bottom: 0;overflow: hidden;">
    <el-header class="d-flex align-items-center" style="background: #3C7699">
      <a class="h5 text-light mb-0 mr-auto">客户端管理系统</a>

      <el-menu
          default-active="0"
          mode="horizontal"
          background-color="#3C7699"
          text-color="#fff"
          active-text-color="#D8E699"
          :ellipsis="false"
          @select="headerSelect">

        <el-menu-item v-for="(item,index) in navBar.list" :key="index" :index="index.toString()">{{item.name}}</el-menu-item>
      </el-menu>
    </el-header>
    <el-container style="height: 100%;padding-bottom: 60px" >
      <el-aside width="200px" style="background: #3C7699">

        <el-menu
            default-active="0"
            background-color="#3C7699"
            text-color="#fff"
            active-text-color="#D8E699"
            @select="slideSelect"
            :ellipsis="false"
            style="height: 100%">
          <el-menu-item v-for="(item,index) in slideMenu" :key="index"
                        :index="index.toString()">
            <el-icon>
              <component :is="item.icon"/>
            </el-icon>
            <span>{{item.name}}</span>
          </el-menu-item>
        </el-menu>


      </el-aside>
      <el-main>

        <component :is="slideMenuComponent"/>

      </el-main>
    </el-container>
  </el-container>
</template>

<script setup lang="ts">
import {ref} from "@vue/reactivity";
import OrderMessageView from "@/components/OrderMessageView.vue";
import firstView from "@/components/Report.vue";
import UserMessage from "@/components/UserMessage.vue";
import CarMessageView from "@/components/CarMessageView.vue";
import AddCar from "@/components/AddCar.vue";
import AddOrder from "@/components/AddOrder.vue";
import EditOrder from "@/components/EditOrder.vue";
import Report from "@/components/Report.vue";




const navBar={

  list:[
    {
      name:"首页",
      subActive:'0',
      subMenu:[
        {
          icon:"House",
          name:"总体数据概览",
          component:Report
        },
      ]
    },
    {name:"订单",
      subActive:'0',
      subMenu:[
        {
          icon:"Odometer",
          name: "当前订单信息",
          component:EditOrder

        },

        {
          icon:"ChatRound",
          name:"历史订单信息",
          component:OrderMessageView
        } ,
      ]},
    {name:"管理车辆",
      subActive:'0',
      subMenu:[
        {
          icon:"Van",
          name:"车辆信息",
          component:CarMessageView
        },
        {
          icon:"CirclePlus",
          name:"添加车辆信息",
          component:AddCar,
        },
      ]},
    {name:"用户信息",
      subActive:'0',
      subMenu:[
        {
          icon:"User",
          name:"用户信息详情",
          component:UserMessage
        },

      ]},
  ]
}

const slideMenu=ref<any>(navBar.list[0].subMenu)

const slideMenuActive=ref<any>(navBar.list[0].subActive)

const slideMenuComponent=ref<any>(slideMenu.value[slideMenuActive.value].component)


//顶栏选择
const headerSelect = (key: string, keyPath: string[]) => {

  //更新左侧导航栏
  slideMenu.value=navBar.list[key].subMenu

  //更新组件页面
  slideMenuComponent.value=slideMenu.value[slideMenuActive.value].component

}

const slideSelect=(key: string, keyPath: string[]) => {
  //更新组件页面
  slideMenuComponent.value=slideMenu.value[key].component
}
</script>

<style lang="scss" scoped>


::-webkit-scrollbar {
  -webkit-appearance: none;
  width: 10px;
  height: 10px;
}

::-webkit-scrollbar-track {
  background-color: rgba(0, 0, 0, .1);
  border-radius: 0;
}

::-webkit-scrollbar-thumb {
  background-color: rgba(0, 0, 0, .25);
  transition: color .2s ease;
  border-radius: 5px;
  cursor: pointer;
}
</style>