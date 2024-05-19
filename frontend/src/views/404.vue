<template>
  <div class="_404">
      <el-text tag="b" @click="$router.replace('home')">抱歉，页面未找到，<span>{{countDown}}</span>s后自动跳转到首页</el-text>
    <img src="@/assets/404.jpg" alt="页面未找到">
  </div>


</template>
<script setup>

import {ref, onMounted, onBeforeUnmount} from 'vue';
import {useRouter} from "vue-router";
const router = useRouter()
//5s设置返回
const countDown = ref(5);

onMounted(() => {
  const timer = setInterval(() => {
    countDown.value--;
    if (countDown.value === 0) {
      clearInterval(timer);
      router.replace('home');
    }
  }, 1000);

  //单页面背景设置与清除
  document
      .querySelector("body")
      .setAttribute("style", "background-color: #dceebc");
  onBeforeUnmount(() => {
    document.querySelector("body").removeAttribute("style")
  })
});
</script>
<style scoped>
body {
  background-color: #dceebc;
}

._404 {
  width: 30%;
  margin: 5rem auto;
}

._404 img {
  width: 30rem;
}

._404 a {
  color: #010101;
}

._404 a:hover {
  color: skyblue;
}
body {
  background-color: initial;
}
</style>