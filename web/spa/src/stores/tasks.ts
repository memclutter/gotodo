import {defineStore} from "pinia";
import type {Task} from "@/models/tasks";
import type {TasksListResponse} from "@/apis/endpoints/tasks/list";
export type TasksState = {
  items: Task[]
  totalCount: Number,
  createDialog: Boolean,
  createStatus: String
  updateDialog: Boolean,
  updateTask: Task
}

export const useTasksStore = defineStore('tasks', {
  state: () => ({
    items: [],
    totalCount: 0,
    createDialog: false,
    createStatus: '',
    updateDialog: false,
    updateTask: <Task>{}
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
    update(task: Task) {
      this.items = this.items.map(item => item.id === task.id ? task : item)
    },
    delete(id: number) {
      this.items = this.items.filter(item => item.id !== id)
      this.totalCount = this.items.length;
    },
    setCreateDialog(createDialog: boolean) {
      this.createDialog = createDialog
    },
    setCreateStatus(createStatus: string) {
      this.createStatus = createStatus
    },
    setUpdateDialog(updateDialog: boolean) {
      this.updateDialog = updateDialog
    },
    setUpdateTask(taskId: number) {
      this.updateTask = this.items.find(item => item.id === taskId)
    },
    unsetUpdateTask() {
      this.updateTask = <Task>{}
    }
  }
})
