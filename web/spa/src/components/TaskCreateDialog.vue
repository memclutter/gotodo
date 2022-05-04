<script setup lang="ts">
import {useTasksStore} from "@/stores/tasks";
import {computed} from "vue";

const tasksStore = useTasksStore();

const dialog = computed({
  get() {
    return tasksStore.createDialog
  },
  set(newValue: boolean) {
    tasksStore.setCreateDialog(newValue)
  }
})
</script>

<template>
  <el-dialog
    v-model="dialog"
    title="Create a new task"
    width="30%"
  >
    <task-form
      v-if="dialog"
      :status="tasksStore.createStatus"
      @success="tasksStore.append($event); tasksStore.setCreateDialog(false)"
    />
  </el-dialog>
</template>
