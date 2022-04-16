<template>
  <el-row align="middle" justify="center">
    <el-col :sm="8">
      <el-card shadow="never">
        <LoginForm/>
      </el-card>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
import LoginForm from "@/components/LoginForm.vue"</script>
