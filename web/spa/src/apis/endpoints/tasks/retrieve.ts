import baseAxios from "@/apis/base";
import type {CustomAxiosResponse} from "@/apis/utils";

export default async function (id: Number): Promise<CustomAxiosResponse> {
  return await baseAxios.get(`/tasks/${id}/`)
}
