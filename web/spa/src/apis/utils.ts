import type {AxiosResponse} from "axios";

export interface CustomAxiosResponse extends AxiosResponse {
  success: boolean,
  message: string,
  validationErrors?: {[key: string]: string[]}
}
