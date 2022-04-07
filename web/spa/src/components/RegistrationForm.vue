<script setup lang="ts">
import {reactive, ref} from "vue";
import authRegistration from '@/apis/endpoints/auth/registration'
import type {FormInstance} from "element-plus";

const ruleFormRef = ref<FormInstance>()
const form = reactive({
  email: '',
  password: '',
  repeatPassword: ''
})

const rules = reactive({
  email: [
    { required: true, message: 'Field is required', trigger: 'blur' },
    { type: 'email', message: 'Field is not correct email address', trigger: 'blur' }
  ],
  password: [
    { required: true, message: 'Field is required', trigger: 'blur' },
  ],
  repeatPassword: [
    { required: true, message: 'Field is required', trigger: 'blur' },
  ]
})

const submit = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  await formEl.validate((valid, fields) => {
    if (valid) {
      const {data, isFinished} = authRegistration(form)
      console.log(data)
    } else {
      console.log('error', fields)
    }
  })
}
</script>

<template>
  <el-form
    ref="ruleFormRef"
    label-width="120px"
    :model="form"
    :rules="rules"
  >
    <el-form-item label="Email" prop="email">
      <el-input type="email" v-model="form.email"/>
    </el-form-item>
    <el-form-item label="Password" prop="password">
      <el-input type="password" v-model="form.password"/>
    </el-form-item>
    <el-form-item label="Repeat password" prop="repeatPassword">
      <el-input type="password" v-model="form.repeatPassword"/>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="submit(ruleFormRef)">Registration</el-button>
    </el-form-item>
  </el-form>
</template>
