<template>
  <el-row align="center" justify="center">
    <el-col :span="8">
      <el-card shadow="never">
        <RegistrationForm/>
      </el-card>
    </el-col>
  </el-row>
</template>
<script lang="ts" setup>
import RegistrationForm from "@/components/RegistrationForm.vue";</script>
