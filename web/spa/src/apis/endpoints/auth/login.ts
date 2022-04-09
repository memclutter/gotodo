import baseAxios from "@/apis/base";

export interface AuthLoginRequest {
  email: String
  password: String
}

export default async function (data: AuthLoginRequest) {
  return await baseAxios.post('/auth/login/', data)
}
