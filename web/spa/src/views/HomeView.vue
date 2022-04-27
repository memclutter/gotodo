<script lang="ts" setup>
import tasksList from '@/apis/endpoints/tasks/list'
import {onMounted} from "vue";
import TaskForm from "@/components/TaskForm.vue";
import TaskBoard from "@/components/TaskBoard.vue";
import {useTasksStore} from "@/stores/tasks";

const tasksStore = useTasksStore()

onMounted(async () => {
  const {success, data} = await tasksList()
  if (success) {
    tasksStore.set(data)
  }
})
</script>
<template>
  <task-form @success="tasksStore.append($event)"/>
  <task-board :items="tasksStore.items"/>
</template>
