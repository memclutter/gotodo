import {useAxios} from "@vueuse/integrations/useAxios";
import baseAxios from "@/apis/base";

export interface TasksUpdateRequest {
    title: String
    status: String
}

export default function (id: Number, data: TasksUpdateRequest) {
    // @todo fix data params for ts
    // @ts-ignore
    return useAxios(`/tasks/${id}/`, {method: 'PUT', data}, baseAxios)
}