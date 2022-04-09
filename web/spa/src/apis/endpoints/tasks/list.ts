import baseAxios from "@/apis/base";

export default async function () {
  return await baseAxios.get('/tasks/')
}
