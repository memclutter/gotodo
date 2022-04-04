import {useAxios} from "@vueuse/integrations/useAxios";
import baseAxios from "@/apis/base";

export interface AuthRegistrationRequest {
    email: String
    password: String
}

export default function (data: AuthRegistrationRequest) {
    // @todo fix data params for ts
    // @ts-ignore
    return useAxios('/auth/registration/', {method: 'POST', data}, baseAxios)
}