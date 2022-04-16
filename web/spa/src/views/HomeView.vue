<script lang="ts" setup>
import tasksList from '@/apis/endpoints/tasks/list'
import {onMounted, ref} from "vue";
import type {Task} from '@/models/tasks'
import TaskForm from "@/components/TaskForm.vue";

const tasks = ref<Task[]>([])
onMounted(async () => {
  const {success, data} = await tasksList()
  if (success) {
    tasks.value = data.items;
  }
})
</script>
<template>
  <task-form />
</template>
