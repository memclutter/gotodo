import {Task} from "@/models/tasks";
import {CustomAxiosResponse} from "@/apis/utils";
import baseAxios from "@/apis/base";
import {TasksListRequest, TasksListResponse} from "@/apis/schemas/tasks";

export async function tasksCreate(data: Task): Promise<CustomAxiosResponse<Task, Task>> {
  return await baseAxios.post('/tasks/', data)
}

export async function tasksDelete(id: Number): Promise<CustomAxiosResponse> {
  return await baseAxios.delete(`/tasks/${id}/`)
}

export async function tasksList(): Promise<CustomAxiosResponse<TasksListResponse, TasksListRequest>> {
  return await baseAxios.get('/tasks/')
}

export async function tasksRetrieve(id: Number): Promise<CustomAxiosResponse<any, Task>> {
  return await baseAxios.get(`/tasks/${id}/`)
}

export async function tasksUpdate(id: number, data: Task): Promise<CustomAxiosResponse<Task, Task>> {
  return await baseAxios.put(`/tasks/${id}/`, data)
}
