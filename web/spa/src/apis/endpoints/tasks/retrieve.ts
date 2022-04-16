import baseAxios from "@/apis/base";
import type {CustomAxiosResponse} from "@/apis/utils";
import type {Task} from "@/models/tasks";

export default async function (id: Number): Promise<CustomAxiosResponse<any, Task>> {
  return await baseAxios.get(`/tasks/${id}/`)
}
