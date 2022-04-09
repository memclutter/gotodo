import baseAxios from "@/apis/base";

export interface TaskCreateRequest {
  title: String
  status: String
}

export default async function (data: TaskCreateRequest) {
  return await baseAxios.post('/tasks/', data)
}
