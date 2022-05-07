import baseAxios from "@/apis/base";
import type {CustomAxiosResponse} from "@/apis/utils";
import type {User} from "@/models/users";

export interface ProfileUpdateRequest {
  firstName: String,
  lastName: String
}

export default async function (data: ProfileUpdateRequest): Promise<CustomAxiosResponse<User, ProfileUpdateRequest>> {
  return await baseAxios.put('/profile/', data)
}
