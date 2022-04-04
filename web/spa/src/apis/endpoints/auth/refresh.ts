import {useAxios} from "@vueuse/integrations/useAxios";
import baseAxios from "@/apis/base";

export interface AuthRefreshRequest {
    refreshToken: String
}

export default function (data: AuthRefreshRequest) {
    // @todo fix data params for ts
    // @ts-ignore
    return useAxios('/auth/refresh/', {method: 'POST', data}, baseAxios)
}