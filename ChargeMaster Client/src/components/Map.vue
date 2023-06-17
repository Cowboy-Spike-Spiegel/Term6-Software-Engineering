<template>
  <div id="myMap">
  </div>
</template>

<script setup lang="ts">
import {onMounted} from "vue";
onMounted(()=>{
  const BMapGL = window.BMapGL
  var map = new BMapGL.Map("myMap");
  map.centerAndZoom(new BMapGL.Point(116.358103,39.961554), 15);
  map.enableScrollWheelZoom(true);


  var marker1 = new BMapGL.Marker(new BMapGL.Point(116.363429,39.969149));
  map.addOverlay(marker1);
  var opts = {
	    width : 200,     // 信息窗口宽度
	    height: 100,     // 信息窗口高度
	    title : "波普特大学充电站" , // 信息窗口标题
	    message:"这里是波普特大学充电站"
	}

  var infoWindow = new BMapGL.InfoWindow("地址：北京市海淀区西土城路10号", opts);

  marker1.addEventListener("click", function(){
		map.openInfoWindow(infoWindow, new BMapGL.Point(116.363429,39.969149)); //开启信息窗口
	});
})

</script>

<style scoped>
#myMap{
  height: 400px;
}
</style>