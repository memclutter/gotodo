import baseAxios from "@/apis/base";
import type {CustomAxiosResponse} from "@/apis/utils";
import type {User} from "@/models/users";

export default async function (): Promise<CustomAxiosResponse<User, any>> {
  return await baseAxios.get('/profile/')
}
