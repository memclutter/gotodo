import baseAxios from "@/apis/base";
import type {CustomAxiosResponse} from "@/apis/utils";
import type {Task} from "@/models/tasks";

export interface TasksListRequest {}

export interface TasksListResponse {
  totalCount: number,
  items: Task[]
}

export default async function (): Promise<CustomAxiosResponse<TasksListResponse, TasksListRequest>> {
  return await baseAxios.get('/tasks/')
}
