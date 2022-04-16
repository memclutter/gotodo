import baseAxios from "@/apis/base";
import type {CustomAxiosResponse} from "@/apis/utils";
import type {Task} from "@/models/tasks";

export default async function (data: Task): Promise<CustomAxiosResponse<Task, Task>> {
  return await baseAxios.post('/tasks/', data)
}
