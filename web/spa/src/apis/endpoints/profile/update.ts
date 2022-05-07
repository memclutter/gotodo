import baseAxios from "@/apis/base";
import type {CustomAxiosResponse} from "@/apis/utils";
import type {User} from "@/models/users";
import type {ProfileUpdateRequest} from "@/apis/schemas/profile";

export default async function (data: ProfileUpdateRequest): Promise<CustomAxiosResponse<User, ProfileUpdateRequest>> {
  return await baseAxios.put('/profile/', data)
}
