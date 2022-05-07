import type {ProfilePasswordUpdateRequest, ProfileUpdateRequest} from "@/apis/schemas/profile";
import type {CustomAxiosResponse} from "@/apis/schemas";
import type {User} from "@/models/users";
import {baseAxios} from "@/apis";

export async function profilePasswordUpdate(data: ProfilePasswordUpdateRequest): Promise<CustomAxiosResponse<User, ProfilePasswordUpdateRequest>> {
  return await baseAxios.put('/profile/password/', data)
}

export async function profileRetrieve(): Promise<CustomAxiosResponse<User, any>> {
  return await baseAxios.get('/profile/')
}

export async function profileUpdate(data: ProfileUpdateRequest): Promise<CustomAxiosResponse<User, ProfileUpdateRequest>> {
  return await baseAxios.put('/profile/', data)
}
