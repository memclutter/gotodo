import baseAxios from "@/apis/base";

export interface AuthConfirmationRequest {
  token: String
}

export default async function (data: AuthConfirmationRequest) {
  return await baseAxios.post('/auth/confirmation/', data)
}
