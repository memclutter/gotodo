<script setup lang="ts">
import {statuses} from '@/models/tasks'
import type {Task} from '@/models/tasks'
import {computed} from "vue";
import TaskCard from "@/components/TaskCard.vue";
import tasksUpdate from "@/apis/endpoints/tasks/update";

const props = defineProps({
  items: Array
})

const boardSpan = computed(() => {
  return 24 / statuses.length;
})

interface DraggableEvent {
  added?: DraggableAddedEvent
  moved?: DraggableMovedEvent
}

interface DraggableAddedEvent {
  element: Task,
  newIndex: number,
}

interface DraggableMovedEvent {
  element: Task,
  oldIndex: number,
  newIndex: number,
}

const changeItems = async (e: DraggableEvent, status) => {
  if (e.added) {
    const {element} = e.added;
    element.status = status.value
    await tasksUpdate(<number>element.id, element)
  } else if (e.moved) {
    const {element, oldIndex, newIndex} = e.moved;
    console.log('moved', status, element, oldIndex, newIndex)
  }
}
</script>

<template>
  <el-row :gutter="28">
    <el-col :span="boardSpan" :key="status.value" v-for="status in statuses">
      <h4>{{ status.title }}</h4>

      <vue-draggable
        :list="items.filter(item => item.status === status.value)"
        group="tasks"
        @change="changeItems($event, status)"
        item-key="id"
        style="width: 100%; height: 100%"
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
