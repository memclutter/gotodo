import type {User} from "@/models/users";

export interface AuthConfirmationRequest {
  token: string
}

export interface AuthLoginRequest {
  email: string
  password: string
}

export interface AuthLoginResponse {
  accessToken: string,
  refreshToken: string,
  user: User
}

export interface AuthRefreshRequest {
  refreshToken: String
}

export interface AuthRefreshResponse {
  accessToken: string,
  refreshToken: string,
  user: User
}

export interface AuthRegistrationRequest {
  firstName: string,
  lastName: string,
  email: string
  password: string
}
