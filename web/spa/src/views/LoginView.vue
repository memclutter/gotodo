<template>
  <el-row justify="center" align="center">
    <el-col :span="8">
      <el-card>
        <LoginForm />
      </el-card>
    </el-col>
  </el-row>
</template>

<script setup lang="ts">
import LoginForm from "@/components/LoginForm.vue"
</script>
