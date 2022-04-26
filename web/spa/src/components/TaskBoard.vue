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
</script>

<template>
  <el-row>
    <el-col :span="boardSpan" :key="status.value" v-for="status in statuses">
      <h4>{{status.title}}</h4>
      <task-card
        v-for="task in items.filter(item => item.status === status.value)"
        :key="task.id"
        v-bind="task"
      />
    </el-col>
  </el-row>
</template>
