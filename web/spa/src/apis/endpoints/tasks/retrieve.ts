import baseAxios from "@/apis/base";

export default async function (id: Number) {
  return await baseAxios.get(`/tasks/${id}/`)
}
