<script setup lang="ts">
import {h, reactive, ref} from "vue";
import authRegistration from '@/apis/endpoints/auth/registration'
import type {FormInstance} from "element-plus";
import 'element-plus/es/components/message/style/css'
import {ElMessage} from "element-plus";

const formLoading = ref(false)
const formRef = ref<FormInstance>()
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
  formLoading.value = true
  await formEl.validate(async(valid, fields) => {
    if (valid) {
      const {success, validationErrors} = await authRegistration(form)
      if (success) {
        ElMessage('Registration success, please check email address')
        formEl.resetFields()
      } else if (validationErrors) {
        // @TODO set validation errors in form
      }
    } else {
      console.log('error', fields)
    }
  })
  formLoading.value = false
}
</script>

<template>
  <el-form
    v-loading="formLoading"
    ref="formRef"
    label-width="120px"
    :model="form"
    :rules="rules"
    @submit.prevent.stop="submit(formRef)"
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
      <el-button type="primary" native-type="submit">Registration</el-button>
    </el-form-item>
  </el-form>
</template>
