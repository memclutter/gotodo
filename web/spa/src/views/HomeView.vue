<script lang="ts" setup>
import tasksList from '@/apis/endpoints/tasks/list'
import {onMounted} from "vue";
import TaskBoard from "@/components/TaskBoard.vue";
import {useTasksStore} from "@/stores/tasks";
import TaskUpdateDialog from "@/components/TaskUpdateDialog.vue";

const tasksStore = useTasksStore()

onMounted(async () => {
  const {success, data} = await tasksList()
  if (success) {
    tasksStore.set(data)
  }
})
</script>
<template>
  <task-board :items="tasksStore.items"/>
  <task-create-dialog />
  <task-update-dialog />
</template>
