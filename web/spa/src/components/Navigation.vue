<template>
  <el-menu
    router
    mode="horizontal"
  >
    <el-menu-item :index="menuItem.to" :key="index" v-for="(menuItem, index) in menuItems">
      {{menuItem.title}}
    </el-menu-item>
  </el-menu>
</template>

<script setup lang="ts">
import {reactive} from "vue";

const menuItems = reactive([
  {
    title: 'Login',
    to: 'login'
  },
  {
    title: 'Registration',
    to: 'registration'
  }
])
</script>
