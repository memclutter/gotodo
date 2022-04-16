import baseAxios from "@/apis/base";
import type {CustomAxiosResponse} from "@/apis/utils";

export interface TasksUpdateRequest {
  title: String
  status: String
}

export default async function (id: Number, data: TasksUpdateRequest): Promise<CustomAxiosResponse> {
  return await baseAxios.put(`/tasks/${id}/`, data)
}
