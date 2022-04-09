import baseAxios from "@/apis/base";

export interface AuthRegistrationRequest {
  email: String
  password: String
}

export default async function (data: AuthRegistrationRequest) {
  return await baseAxios.post('/auth/registration/', data)
}
