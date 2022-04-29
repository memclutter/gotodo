<script setup lang="ts">
import {statuses} from '@/models/tasks'
import {computed} from "vue";
import TaskCard from "@/components/TaskCard.vue";

const props = defineProps({
  items: Array
})

const boardSpan = computed(() => {
  return 24 / statuses.length;
})

const itemsByStatus = {
  'todo': [
    {id: 1, title: 'Task [todo]', status: 'todo'},
    {id: 3, title: 'Task 2 [todo]', status: 'todo'}
  ],
  'inProgress': [

  ],
  'closed': [
    {id: 2, title: 'Task 3', status: 'closed'}
  ]
}
</script>

<template>
  <el-row :gutter="28">
    <el-col :span="boardSpan" :key="status.value" v-for="status in statuses">
      <h4>{{ status.title }}</h4>

      <vue-draggable
        :list="itemsByStatus[status.value]"
        :group="status.value"
        item-key="id"
      >
        <template #item="{element}">
          <task-card
            style="margin-bottom: 10px"
            v-bind="element"
          />
        </template>
      </vue-draggable>
    </el-col>
  </el-row>
</template>
