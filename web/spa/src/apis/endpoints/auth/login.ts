import baseAxios from "@/apis/base";
import {useAxios} from "@vueuse/integrations/useAxios";

export interface AuthLoginRequest {
    email: String
    password: String
}

export default function (data: AuthLoginRequest) {
    // @todo fix data params for ts
    // @ts-ignore
    return useAxios('/auth/login/', {method: 'POST', data}, baseAxios)
}