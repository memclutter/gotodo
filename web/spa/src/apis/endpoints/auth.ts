import {baseAxios} from "@/apis";
import type {CustomAxiosResponse} from "@/apis/schemas";
import type {AuthConfirmationRequest} from "@/apis/schemas/auth";
import type {
  AuthLoginRequest,
  AuthLoginResponse,
  AuthRefreshRequest,
  AuthRefreshResponse,
  AuthRegistrationRequest
} from "@/apis/schemas/auth";

export async function authConfirmation(data: AuthConfirmationRequest): Promise<CustomAxiosResponse> {
  return await baseAxios.post('/auth/confirmation/', data)
}

export async function authLogin(data: AuthLoginRequest): Promise<CustomAxiosResponse<AuthLoginResponse, AuthLoginRequest>> {
  return await baseAxios.post('/auth/login/', data)
}

export async function authRefresh(data: AuthRefreshRequest): Promise<CustomAxiosResponse<AuthRefreshResponse, AuthRefreshRequest>> {
  return await baseAxios.post('/auth/refresh/', data)
}

export async function authRegistration(data: AuthRegistrationRequest): Promise<CustomAxiosResponse> {
  return await baseAxios.post('/auth/registration/', data)
}
