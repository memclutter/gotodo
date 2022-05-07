import baseAxios from "@/apis/base";
import type {CustomAxiosResponse} from "@/apis/utils";
import type {TasksListRequest, TasksListResponse} from "@/apis/schemas/tasks";

export default async function (): Promise<CustomAxiosResponse<TasksListResponse, TasksListRequest>> {
  return await baseAxios.get('/tasks/')
}
