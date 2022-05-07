import baseAxios from "@/apis/base";
import type {CustomAxiosResponse} from "@/apis/utils";
import type {AuthConfirmationRequest} from "@/apis/schemas/auth";

export default async function (data: AuthConfirmationRequest): Promise<CustomAxiosResponse> {
  return await baseAxios.post('/auth/confirmation/', data)
}
