<template>
  <el-row justify="center" align="center">
    <el-col :span="8">
      <el-card shadow="never">
        <RegistrationForm/>
      </el-card>
    </el-col>
  </el-row>
</template>
<script setup lang="ts">
import RegistrationForm from "@/components/RegistrationForm.vue";
</script>
