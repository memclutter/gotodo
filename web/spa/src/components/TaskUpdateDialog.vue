<script setup lang="ts">
import {useTasksStore} from "@/stores/tasks";
import {computed} from "vue";

const tasksStore = useTasksStore();

const dialog = computed({
  get() {
    return tasksStore.updateDialog
  },
  set(newValue: boolean) {
    tasksStore.setUpdateDialog(newValue)
  }
})
</script>

<template>
  <el-dialog
    v-model="dialog"
    title="Update task"
    width="30%"
  >
    <task-form
      v-if="dialog"
      v-bind="tasksStore.updateTask"
      @success="tasksStore.update($event); tasksStore.setUpdateDialog(false)"
    />
  </el-dialog>
</template>
