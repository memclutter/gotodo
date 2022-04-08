<template>
  <el-row justify="center" align="center">
    <el-col :sm="8">
      <el-card shadow="never">
        <LoginForm />
      </el-card>
    </el-col>
  </el-row>
</template>

<script setup lang="ts">
import LoginForm from "@/components/LoginForm.vue"
</script>
