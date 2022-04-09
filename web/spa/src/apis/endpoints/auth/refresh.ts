import baseAxios from "@/apis/base";

export interface AuthRefreshRequest {
  refreshToken: String
}

export default async function (data: AuthRefreshRequest) {
  return await baseAxios.post('/auth/refresh/', data)
}
