import baseAxios from "@/apis/base";
import type {CustomAxiosResponse} from "@/apis/utils";
import type {User} from "@/models/users";

export interface ProfilePasswordUpdateRequest {
  oldPassword: String,
  newPassword: String
}

export default async function (data: ProfilePasswordUpdateRequest): Promise<CustomAxiosResponse<User, ProfilePasswordUpdateRequest>> {
  return await baseAxios.put('/profile/password/', data)
}
