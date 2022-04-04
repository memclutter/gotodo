import {useAxios} from "@vueuse/integrations/useAxios";
import baseAxios from "@/apis/base";

export interface TaskCreateRequest {
    title: String
    status: String
}

export default function (data: TaskCreateRequest) {
    // @todo fix data params for ts
    // @ts-ignore
    return useAxios('/tasks/', {method: 'POST', data}, baseAxios)
}