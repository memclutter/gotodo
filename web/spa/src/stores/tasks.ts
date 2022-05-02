import {defineStore} from "pinia";
import type {Task} from "@/models/tasks";
import type {TasksListResponse} from "@/apis/endpoints/tasks/list";
export type TasksState = {
  items: Task[]
  totalCount: Number,
  createDialog: Boolean,
}

export const useTasksStore = defineStore('tasks', {
  state: () => ({
    items: [],
    totalCount: 0,
    createDialog: false,
  } as TasksState),
  actions: {
    set(state: TasksListResponse) {
      this.items = state.items
      this.totalCount = state.totalCount
    },
    append(task: Task) {
      this.items.push(task)
      this.totalCount += 1
    },
    delete(id: number) {
      this.items = this.items.filter(item => item.id !== id)
      this.totalCount = this.items.length;
    },
    setCreateDialog(createDialog: boolean) {
      this.createDialog = createDialog
    }
  }
})
