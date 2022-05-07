<script setup lang="ts">
import {reactive, ref} from "vue";
import type {FormInstance, FormItemRule} from "element-plus";
import type {Arrayable} from "element-plus/es/utils";
import {profileUpdate} from "@/apis/endpoints/profile";
import {useAuthStore} from "@/stores/auth";
import {authRefresh} from "@/apis/endpoints/auth";

const authStore = useAuthStore()
const formLoading = ref<boolean>(false);
const formRef = ref<FormInstance>()
const form = reactive({
  firstName: '' || authStore.user.firstName,
  lastName: '' || authStore.user.lastName
})
const serverErrors = reactive<{ [key: string]: string | undefined }>({
  firstName: undefined,
  lastName: undefined
})

const rules: Partial<Record<string, Arrayable<FormItemRule>>> = reactive({
  firstName: [
    {required: true, message: 'Field is required', trigger: 'blur'},
  ],
  lastName: [
    {required: true, message: 'Field is required', trigger: 'blur'}
  ]
})

const submit = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formLoading.value = true;
  await formEl.validate(async (valid) => {
    if (valid) {
      const {success, validationErrors} = await profileUpdate(form)
      if (success) {
        const {success, data} = await authRefresh({refreshToken: authStore.refreshToken});
        if (success) {
          authStore.set(data)
        }
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
    <el-form-item label="Firstname" prop="firstName" :error="serverErrors.firstName">
      <el-input v-model="form.firstName"/>
    </el-form-item>
    <el-form-item label="Lastname" prop="lastName" :error="serverErrors.lastName">
      <el-input v-model="form.lastName"/>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" native-type="submit">Save</el-button>
    </el-form-item>
  </el-form>
</template>
