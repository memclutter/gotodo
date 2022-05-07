import baseAxios from "@/apis/base";
import type {CustomAxiosResponse} from "@/apis/utils";
import type {User} from "@/models/users";
import type {ProfilePasswordUpdateRequest} from "@/apis/schemas/profile";

export default async function (data: ProfilePasswordUpdateRequest): Promise<CustomAxiosResponse<User, ProfilePasswordUpdateRequest>> {
  return await baseAxios.put('/profile/password/', data)
}
