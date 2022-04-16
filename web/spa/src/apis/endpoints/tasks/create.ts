import baseAxios from "@/apis/base";
import type {CustomAxiosResponse} from "@/apis/utils";

export interface TaskCreateRequest {
  title: String
  status: String
}

export default async function (data: TaskCreateRequest): Promise<CustomAxiosResponse> {
  return await baseAxios.post('/tasks/', data)
}
