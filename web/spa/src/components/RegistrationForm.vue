<script lang="ts" setup>
import {reactive, ref} from "vue";
import authRegistration from '@/apis/endpoints/auth/registration'
import type {FormInstance} from "element-plus";
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
    {required: true, message: 'Field is required', trigger: 'blur'},
    {type: 'email', message: 'Field is not correct email address', trigger: 'blur'}
  ],
  password: [
    {required: true, message: 'Field is required', trigger: 'blur'},
  ],
  repeatPassword: [
    {required: true, message: 'Field is required', trigger: 'blur'},
  ]
})

const submit = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formLoading.value = true
  await formEl.validate(async (valid, fields) => {
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
    ref="formRef"
    v-loading="formLoading"
    :model="form"
    :rules="rules"
    label-width="120px"
    @submit.prevent.stop="submit(formRef)"
  >
    <el-form-item label="Email" prop="email">
      <el-input v-model="form.email" type="email"/>
    </el-form-item>
    <el-form-item label="Password" prop="password">
      <el-input v-model="form.password" type="password"/>
    </el-form-item>
    <el-form-item label="Repeat password" prop="repeatPassword">
      <el-input v-model="form.repeatPassword" type="password"/>
    </el-form-item>
    <el-form-item>
      <el-button native-type="submit" type="primary">Registration</el-button>
    </el-form-item>
  </el-form>
</template>
