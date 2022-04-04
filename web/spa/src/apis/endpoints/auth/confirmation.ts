import {useAxios} from "@vueuse/integrations/useAxios";
import baseAxios from "@/apis/base";

export interface AuthConfirmationRequest {
    token: String
}

export default function (data: AuthConfirmationRequest) {
    // @todo fix data params for ts
    // @ts-ignore
    return useAxios('/auth/confirmation/', {method: 'POST', data}, baseAxios)
}