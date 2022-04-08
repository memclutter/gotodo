<script setup lang="ts">
import {reactive, ref} from "vue";
import authLogin from '@/apis/endpoints/auth/login'
import type {FormInstance} from "element-plus";
import authRegistration from "@/apis/endpoints/auth/registration";

const formLoading = ref(false)
const formRef = ref<FormInstance>()
const form = reactive({
  email: '',
  password: ''
})

const rules = reactive({
  email: [
    { required: true, message: 'Field is required', trigger: 'blur' },
    { type: 'email', message: 'Field is not correct email address', trigger: 'blur' }
  ],
  password: [
    { required: true, message: 'Field is required', trigger: 'blur' }
  ]
})

const submit = async(formEl: ref<FormInstance> | undefined) => {
  if (!formEl) return
  formLoading.value = true;
  await formEl.validate((valid, fields) => {
    if (valid) {
      const {data, isFinished} = authRegistration(form)
      console.log(data)
    } else {
      console.log('error', fields)
    }
  })
  formLoading.value = false;
}
</script>

<template>
  <el-form
    v-loading="formLoading"
    ref="formRef"
    :model="form"
    :rules="rules"
    label-width="120px"
  >
    <el-form-item label="Email" prop="email">
      <el-input type="email" v-model="form.email"/>
    </el-form-item>
    <el-form-item label="Password" prop="password">
      <el-input type="password" v-model="form.password"/>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="submit(formRef)">Login</el-button>
    </el-form-item>
  </el-form>
</template>
