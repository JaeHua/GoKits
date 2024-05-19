
<template>
  <el-menu
      :default-active="activeIndex"
      class="el-menu-demo"
      mode="horizontal"
      :ellipsis="false"
      @select="handleSelect"
  >
    <el-menu-item index="0" @click="$router.replace('home')">
      <img
          style="width: 100px"
          src="@/assets/logo.png"
          alt="Element logo"
      />
    </el-menu-item>
    <div class="flex-grow" />

    <el-menu-item index="1" v-if="$route.name!='Register'&&!userInfo" @click="$router.push({name:'Register'})"><el-text tag="b"><el-icon><Promotion /></el-icon>注册</el-text></el-menu-item>
    <el-menu-item index="1" v-if="$route.name!='Login'&&!userInfo" @click="$router.push({name:'Login'})"><el-text tag="b"><el-icon><Key /></el-icon>登陆</el-text></el-menu-item>
    &nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp

    <el-sub-menu v-if="userInfo" index="2"  >
      <template #title >      <el-avatar
          src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
      /><el-text tag="b">{{userInfo.name}}</el-text></template>
      <el-menu-item index="2-1" @click="$router.replace('profile')"><el-text tag="b"><el-icon><User /></el-icon>个人中心</el-text></el-menu-item>
      <el-menu-item index="2-3" @click="$router.replace('home')" ><el-text tag="b"><el-icon><House /></el-icon>主页</el-text></el-menu-item>

      <el-menu-item index="2-2" @click="logout" ><el-text tag="b"><el-icon><SwitchButton /></el-icon>登出</el-text></el-menu-item>
<!--      <el-menu-item index="2-3">item three</el-menu-item>-->

    </el-sub-menu>
  </el-menu>
</template>

<script lang="ts" setup>
import { useStore } from 'vuex'
import {computed, ref} from 'vue'

const activeIndex = ref('1')

const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
const store = useStore()
const userInfo = computed(()=>{
  // return JSON.parse(storageService.get(storageService.USER_INFO))
  return store.state.userModule.userInfo
})
const logout =()=>{
  store.dispatch('userModule/logout')
}
import { User,SwitchButton,Promotion,Key,House,Menu } from '@element-plus/icons-vue'
</script>

<style scoped>
.flex-grow {
  flex-grow: 1;
}
</style>
