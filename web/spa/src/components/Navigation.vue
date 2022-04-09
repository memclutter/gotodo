<template>
  <el-menu
    mode="horizontal"
    router
  >
    <el-menu-item>
      GoTODO
    </el-menu-item>
    <el-menu-item v-for="(menuItem, index) in menuItems" :key="index" :index="menuItem.to">
      {{ menuItem.title }}
    </el-menu-item>
  </el-menu>
</template>

<script lang="ts" setup>
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
