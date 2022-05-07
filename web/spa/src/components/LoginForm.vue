<script lang="ts" setup>
import {reactive, ref} from "vue";
import type {FormInstance, FormItemRule} from "element-plus";
import {useAuthStore} from "@/stores/auth";
import {authLogin} from "@/apis/endpoints/auth";
import {useRouter} from "vue-router";
import type {Arrayable} from "element-plus/es/utils";

const router = useRouter()
const authStore = useAuthStore()
const formLoading = ref(false)
const formRef = ref<FormInstance>()
const form = reactive({
  email: '',
  password: ''
})
const serverErrors = reactive<{ [key: string]: string | undefined }>({
  email: undefined,
  password: undefined
})

const rules: Partial<Record<string, Arrayable<FormItemRule>>> = reactive({
  email: [
    {required: true, message: 'Field is required', trigger: 'blur'},
    {type: 'email', message: 'Field is not correct email address', trigger: 'blur'}
  ],
  password: [
    {required: true, message: 'Field is required', trigger: 'blur'}
  ]
})

const submit = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formLoading.value = true;
  await formEl.validate(async (valid) => {
    if (valid) {
      const {success, data, validationErrors} = await authLogin(form)
      if (success) {
        authStore.set(data)
        await router.replace({name: 'home'})
      } else if (validationErrors) {
        for (const key in validationErrors) {
          if (serverErrors.hasOwnProperty(key)) {
            serverErrors[key] = validationErrors[key].join(' ')
          }
        }
      }
    }
  })
  formLoading.value = false;
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
    <el-form-item label="Email" prop="email" :error="serverErrors.email">
      <el-input v-model="form.email" type="email"/>
    </el-form-item>
    <el-form-item label="Password" prop="password" :error="serverErrors.password">
      <el-input v-model="form.password" type="password"/>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" native-type="submit">Login</el-button>
    </el-form-item>
  </el-form>
</template>
