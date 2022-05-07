import {ProfilePasswordUpdateRequest, ProfileUpdateRequest} from "@/apis/schemas/profile";
import {CustomAxiosResponse} from "@/apis/utils";
import {User} from "@/models/users";
import baseAxios from "@/apis/base";

export async function profilePasswordUpdate(data: ProfilePasswordUpdateRequest): Promise<CustomAxiosResponse<User, ProfilePasswordUpdateRequest>> {
  return await baseAxios.put('/profile/password/', data)
}

export async function profileRetrieve(): Promise<CustomAxiosResponse<User, any>> {
  return await baseAxios.get('/profile/')
}

export async function profileUpdate(data: ProfileUpdateRequest): Promise<CustomAxiosResponse<User, ProfileUpdateRequest>> {
  return await baseAxios.put('/profile/', data)
}
