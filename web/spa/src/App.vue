<script lang="ts" setup>
import {RouterView} from 'vue-router'
import Navigation from "@/components/Navigation.vue"</script>

<template>
  <el-container>
    <el-header>
      <navigation/>
    </el-header>
    <el-main>
      <RouterView/>
    </el-main>
  </el-container>
</template>
