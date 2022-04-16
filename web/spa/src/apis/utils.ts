import type {AxiosResponse} from "axios";

export interface CustomAxiosResponse<T = any, D = any> extends AxiosResponse<T, D> {
  success: boolean,
  message: string,
  validationErrors?: {[key: string]: string[]}
}
