<script setup lang="ts">
import {reactive} from "vue";
import authLogin from '@/apis/endpoints/auth/login'

const form = reactive({
  email: '',
  password: ''
})

function submit() {
  const {data, isFinished} = authLogin(form)
  console.log(data)
}
</script>

<template>
  <el-form :model="form" label-width="120px">
    <el-form-item label="Email">
      <el-input type="email" v-model="form.email"/>
    </el-form-item>
    <el-form-item label="Password">
      <el-input type="password" v-model="form.password"/>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="submit">Login</el-button>
    </el-form-item>
  </el-form>
</template>