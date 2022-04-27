<script lang="ts" setup>
import {reactive, ref} from "vue";
import type {FormInstance, FormItemRule} from "element-plus";
import tasksCreate from "@/apis/endpoints/tasks/create";
import {useRouter} from "vue-router";
import type {Arrayable} from "element-plus/es/utils";
import TaskStatusSelect from "@/components/TaskStatusSelect.vue";

const router = useRouter()
const formLoading = ref(false)
const formRef = ref<FormInstance>()
const form = reactive({
  title: '',
  description: '',
  status: ''
})
const serverErrors = reactive<{ [key: string]: string | undefined }>({
  title: undefined,
  description: undefined,
  status: undefined
})

const rules: Partial<Record<string, Arrayable<FormItemRule>>> = reactive({
  title: [
    {required: true, message: 'Field is required', trigger: 'blur'},
  ],
  status: [
    {required: true, message: 'Field is required', trigger: 'blur'}
  ]
})

const submit = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formLoading.value = true;
  await formEl.validate(async (valid) => {
    if (valid) {
      const {success, validationErrors} = await tasksCreate(form)
      if (success) {
        // @TODO handle success
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
    <el-form-item label="Title" prop="title" :error="serverErrors.title">
      <el-input v-model="form.title" type="text"/>
    </el-form-item>
    <el-form-item label="Status" prop="status" :error="serverErrors.status">
      <task-status-select v-model="form.status" />
    </el-form-item>
    <el-form-item>
      <el-button type="primary" native-type="submit">Submit</el-button>
    </el-form-item>
  </el-form>
</template>
