import {defineStore} from "pinia";
import type {Task} from "@/models/tasks";
export type TasksState = {
  items: Task[]
  totalCount: Number
}

export const useTasksStore = defineStore('tasks', {
  state: () => ({
    items: [],
    totalCount: 0
  } as TasksState),
  actions: {
    set(state: TasksState) {
      this.items = state.items
      this.totalCount = state.totalCount
    },
    append(task: Task) {
      this.items.push(task)
      this.totalCount += 1
    }
  }
})
