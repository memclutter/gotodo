import baseAxios from "@/apis/base";

export interface TasksUpdateRequest {
  title: String
  status: String
}

export default async function (id: Number, data: TasksUpdateRequest) {
  return await baseAxios.put(`/tasks/${id}/`, data)
}
