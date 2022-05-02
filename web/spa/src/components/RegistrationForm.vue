<script lang="ts" setup>
import {reactive, ref} from "vue";
import authRegistration from '@/apis/endpoints/auth/registration'
import type {FormInstance, FormItemRule} from "element-plus";
import {ElMessage} from "element-plus";
import type {Arrayable} from "element-plus/es/utils";

const formLoading = ref(false)
const formRef = ref<FormInstance>()
const form = reactive({
  firstName: '',
  lastName: '',
  email: '',
  password: '',
  repeatPassword: ''
})
const serverErrors = reactive<{ [key: string]: string | undefined }>({
  firstName: undefined,
  lastName: undefined,
  email: undefined,
  password: undefined
})

const rules: Partial<Record<string, Arrayable<FormItemRule>>> = reactive({
  firstName: [
    {required: true, message: 'Field is required', trigger: 'blur'},
  ],
  lastName: [
    {required: true, message: 'Field is required', trigger: 'blur'},
  ],
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
  await formEl.validate(async (valid) => {
    if (!valid) return
    const {success, validationErrors} = await authRegistration(form)
    if (success) {
      ElMessage('Registration success, please check email address')
      formEl.resetFields()
    } else if (validationErrors) {
      for (const key in validationErrors) {
        if (serverErrors.hasOwnProperty(key)) {
          serverErrors[key] = validationErrors[key].join(' ')
        }
      }
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
    <el-form-item label="Firstname" prop="firstName" :error="serverErrors.firstName">
      <el-input v-model="form.firstName"/>
    </el-form-item>
    <el-form-item label="Lastname" prop="lastName" :error="serverErrors.lastName">
      <el-input v-model="form.lastName"/>
    </el-form-item>
    <el-form-item label="Email" prop="email" :error="serverErrors.email">
      <el-input v-model="form.email" type="email"/>
    </el-form-item>
    <el-form-item label="Password" prop="password" :error="serverErrors.password">
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
